#epoll工作原理
epoll同样只告知那些就绪的文件描述符，而且当我们调用epoll_wait()获得就绪文件描述符时，返回的不是实际的描述符，而是一个代表就绪描述符数量的值，你只需要去epoll指定的一个数组中依次取得相应数量的文件描述符即可，这里也使用了内存映射（mmap）技术，这样便彻底省掉了这些文件描述符在系统调用时复制的开销。
 
另一个本质的改进在于epoll采用基于事件的就绪通知方式。在select/poll中，进程只有在调用一定的方法后，内核才对所有监视的文件描述符进行扫描，而epoll事先通过epoll_ctl()来注册一个文件描述符，一旦基于某个文件描述符就绪时，内核会采用类似callback的回调机制，迅速激活这个文件描述符，当进程调用epoll_wait()时便得到通知。
##Epoll的2种工作方式-水平触发（LT）和边缘触发（ET）
假如有这样一个例子：

1. 我们已经把一个用来从管道中读取数据的文件句柄(RFD)添加到epoll描述符
2. 这个时候从管道的另一端被写入了2KB的数据
3. 调用epoll_wait(2)，并且它会返回RFD，说明它已经准备好读取操作
4. 然后我们读取了1KB的数据
5. 调用epoll_wait(2)......
###Edge Triggered 工作模式
如果我们在第1步将RFD添加到epoll描述符的时候使用了EPOLLET标志，那么在第5步调用epoll_wait(2)之后将有可能会挂起，因为剩余的数据还存在于文件的输入缓冲区内，而且数据发出端还在等待一个针对已经发出数据的反馈信息。只有在监视的文件句柄上发生了某个事件的时候 ET 工作模式才会汇报事件。因此在第5步的时候，调用者可能会放弃等待仍在存在于文件输入缓冲区内的剩余数据。在上面的例子中，会有一个事件产生在RFD句柄上，因为在第2步执行了一个写操作，然后，事件将会在第3步被销毁。因为第4步的读取操作没有读空文件输入缓冲区内的数据，因此我们在第5步调用 epoll_wait(2)完成后，是否挂起是不确定的。epoll工作在ET模式的时候，必须使用非阻塞套接口，以避免由于一个文件句柄的阻塞读/阻塞写操作把处理多个文件描述符的任务饿死。最好以下面的方式调用ET模式的epoll接口，在后面会介绍避免可能的缺陷。

- 基于非阻塞文件句柄
- 只有当read(2)或者write(2)返回EAGAIN时才需要挂起，等待。但这并不是说每次read()时都需要循环读，直到读到产生一个EAGAIN才认为此次事件处理完成，当read()返回的读到的数据长度小于请求的数据长度时，就可以确定此时缓冲中已没有数据了，也就可以认为此事读事件已处理完成。

###Level Triggered 工作模式
相反的，以LT方式调用epoll接口的时候，它就相当于一个速度比较快的poll(2)，并且无论后面的数据是否被使用，因此他们具有同样的职能。因为即使使用ET模式的epoll，在收到多个chunk的数据的时候仍然会产生多个事件。调用者可以设定EPOLLONESHOT标志，在 epoll_wait(2)收到事件后epoll会与事件关联的文件句柄从epoll描述符中禁止掉。因此当EPOLLONESHOT设定后，使用带有 EPOLL_CTL_MOD标志的epoll_ctl(2)处理文件句柄就成为调用者必须作的事情。

LT(level triggered)是epoll缺省的工作方式，并且同时支持block和no-block socket.在这种做法中，内核告诉你一个文件描述符是否就绪了，然后你可以对这个就绪的fd进行IO操作。如果你不作任何操作，内核还是会继续通知你 的，所以，这种模式编程出错误可能性要小一点。传统的select/poll都是这种模型的代表．
 
ET (edge-triggered)是高速工作方式，只支持no-block socket，它效率要比LT更高。ET与LT的区别在于，当一个新的事件到来时，ET模式下当然可以从epoll_wait调用中获取到这个事件，可是如果这次没有把这个事件对应的套接字缓冲区处理完，在这个套接字中没有新的事件再次到来时，在ET模式下是无法再次从epoll_wait调用中获取这个事件的。而LT模式正好相反，只要一个事件对应的套接字缓冲区还有数据，就总能从epoll_wait中获取这个事件。

因此，LT模式下开发基于epoll的应用要简单些，不太容易出错。而在ET模式下事件发生时，如果没有彻底地将缓冲区数据处理完，则会导致缓冲区中的用户请求得不到响应。

Nginx默认采用ET模式来使用epoll。
 
##epoll的优点：

1. 支持一个进程打开大数目的socket描述符(FD)

select 最不能忍受的是一个进程所打开的FD是有一定限制的，由FD_SETSIZE设置，默认值是2048。对于那些需要支持的上万连接数目的IM服务器来说显然太少了。这时候你一是可以选择修改这个宏然后重新编译内核，不过资料也同时指出这样会带来网络效率的下降，二是可以选择多进程的解决方案(传统的 Apache方案)，不过虽然linux上面创建进程的代价比较小，但仍旧是不可忽视的，加上进程间数据同步远比不上线程间同步的高效，所以也不是一种完美的方案。不过 epoll则没有这个限制，它所支持的FD上限是最大可以打开文件的数目，这个数字一般远大于2048,举个例子,在1GB内存的机器上大约是10万左右，具体数目可以cat /proc/sys/fs/file-max察看,一般来说这个数目和系统内存关系很大。
 
2. IO效率不随FD数目增加而线性下降

传统的select/poll另一个致命弱点就是当你拥有一个很大的socket集合，不过由于网络延时，任一时间只有部分的socket是"活跃"的，但是select/poll每次调用都会线性扫描全部的集合，导致效率呈现线性下降。但是epoll不存在这个问题，它只会对"活跃"的socket进行操作---这是因为在内核实现中epoll是根据每个fd上面的callback函数实现的。那么，只有"活跃"的socket才会主动的去调用 callback函数，其他idle状态socket则不会，在这点上，epoll实现了一个"伪"AIO，因为这时候推动力在os内核。在一些 benchmark中，如果所有的socket基本上都是活跃的---比如一个高速LAN环境，epoll并不比select/poll有什么效率，相反，如果过多使用epoll_ctl,效率相比还有稍微的下降。但是一旦使用idle connections模拟WAN环境,epoll的效率就远在select/poll之上了。
 
3. 使用mmap加速内核与用户空间的消息传递

这点实际上涉及到epoll的具体实现了。无论是select,poll还是epoll都需要内核把FD消息通知给用户空间，如何避免不必要的内存拷贝就很重要，在这点上，epoll是通过内核于用户空间mmap同一块内存实现的。而如果你想我一样从2.5内核就关注epoll的话，一定不会忘记手工 mmap这一步的。
 
4. 内核微调

这一点其实不算epoll的优点了，而是整个linux平台的优点。也许你可以怀疑linux平台，但是你无法回避linux平台赋予你微调内核的能力。比如，内核TCP/IP协议栈使用内存池管理sk_buff结构，那么可以在运行时期动态调整这个内存pool(skb_head_pool)的大小--- 通过echo XXXX>/proc/sys/net/core/hot_list_length完成。再比如listen函数的第2个参数(TCP完成3次握手的数据包队列长度)，也可以根据你平台内存大小动态调整。更甚至在一个数据包面数目巨大但同时每个数据包本身大小却很小的特殊系统上尝试最新的NAPI网卡驱动架构。
 
##linux下epoll实现高效处理百万句柄
开发高性能网络程序时，windows开发者们言必称iocp，linux开发者们则言必称epoll。大家都明白epoll是一种IO多路复用技术，可以非常高效的处理数以百万计的socket句柄，比起以前的select和poll效率高大发了。我们用起epoll来都感觉挺爽，确实快，那么，它到底为什么可以高速处理这么多并发连接呢？
 
使用起来很清晰，首先要调用epoll_create建立一个epoll对象。参数size是内核保证能够正确处理的最大句柄数，多于这个最大数时内核可不保证效果。
 
epoll_ctl可以操作上面建立的epoll，例如，将刚建立的socket加入到epoll中让其监控，或者把 epoll正在监控的某个socket句柄移出epoll，不再监控它等等。
 
epoll_wait在调用时，在给定的timeout时间内，当在监控的所有句柄中有事件发生时，就返回用户态的进程。
 
从上面的调用方式就可以看到epoll比select/poll的优越之处：因为后者每次调用时都要传递你所要监控的所有socket给select/poll系统调用，这意味着需要将用户态的socket列表copy到内核态，如果以万计的句柄会导致每次都要copy几十几百KB的内存到内核态，非常低效。而我们调用epoll_wait时就相当于以往调用select/poll，但是这时却不用传递socket句柄给内核，因为内核已经在epoll_ctl中拿到了要监控的句柄列表。

所以，实际上在你调用epoll_create后，内核就已经在内核态开始准备帮你存储要监控的句柄了，每次调用epoll_ctl只是在往内核的数据结构里塞入新的socket句柄。

当一个进程调用epoll_creaqte方法时，Linux内核会创建一个eventpoll结构体，这个结构体中有两个成员与epoll的使用方式密切相关。

每一个epoll对象都有一个独立的eventpoll结构体，这个结构体会在内核空间中创造独立的内存，用于存储使用epoll_ctl方法向epoll对象中添加进来的事件。这样，重复的事件就可以通过红黑树而高效的识别出来。

此外，epoll还维护了一个双链表，用户存储发生的事件。当epoll_wait调用时，仅仅观察这个list链表里有没有数据即eptime项即可。有数据就返回，没有数据就sleep，等到timeout时间到后即使链表没数据也返回。所以，epoll_wait非常高效。
 
而且，通常情况下即使我们要监控百万计的句柄，大多一次也只返回很少量的准备就绪句柄而已，所以，epoll_wait仅需要从内核态copy少量的句柄到用户态而已，如何能不高效？！
 
那么，这个准备就绪list链表是怎么维护的呢？当我们执行epoll_ctl时，除了把socket放到epoll文件系统里file对象对应的红黑树上之外，还会给内核中断处理程序注册一个回调函数，告诉内核，如果这个句柄的中断到了，就把它放到准备就绪list链表里。所以，当一个socket上有数据到了，内核在把网卡上的数据copy到内核中后就来把socket插入到准备就绪链表里了。
 
如此，一颗红黑树，一张准备就绪句柄链表，少量的内核cache，就帮我们解决了大并发下的socket处理问题。执行epoll_create时，创建了红黑树和就绪链表，执行epoll_ctl时，如果增加socket句柄，则检查在红黑树中是否存在，存在立即返回，不存在则添加到树干上，然后向内核注册回调函数，用于当中断事件来临时向准备就绪链表中插入数据。执行epoll_wait时立刻返回准备就绪链表里的数据即可。

##epoll的使用方法
那么究竟如何来使用epoll呢？其实非常简单。
 
通过在包含一个头文件#include <sys/epoll.h> 以及几个简单的API将可以大大的提高你的网络服务器的支持人数。
 
首先通过create_epoll(int maxfds)来创建一个epoll的句柄。这个函数会返回一个新的epoll句柄，之后的所有操作将通过这个句柄来进行操作。在用完之后，记得用close()来关闭这个创建出来的epoll句柄。
 
之后在你的网络主循环里面，每一帧的调用epoll_wait(int epfd, epoll_event events, int max events, int timeout)来查询所有的网络接口，看哪一个可以读，哪一个可以写了。基本的语法为：
<pre>
nfds = epoll_wait(kdpfd, events, maxevents, -1);
</pre>
其中kdpfd为用epoll_create创建之后的句柄，events是一个epoll_event*的指针，当epoll_wait这个函数操作成功之后，epoll_events里面将储存所有的读写事件。max_events是当前需要监听的所有socket句柄数。最后一个timeout是 epoll_wait的超时，为0的时候表示马上返回，为-1的时候表示一直等下去，直到有事件返回，为任意正整数的时候表示等这么长的时间，如果一直没有事件，则返回。一般如果网络主循环是单独的线程的话，可以用-1来等，这样可以保证一些效率，如果是和主逻辑在同一个线程的话，则可以用0来保证主循环的效率。
 
epoll_wait返回之后应该是一个循环，遍历所有的事件。

几乎所有的epoll程序都使用下面的框架：
<pre>
for( ; ; )  
   {  
       nfds = epoll_wait(epfd,events,20,500);  
       for(i=0;i<nfds;++i)  
       {  
           if(events[i].data.fd==listenfd) //有新的连接  
           {  
               connfd = accept(listenfd,(sockaddr *)&clientaddr, &clilen); //accept这个连接  
               ev.data.fd=connfd;  
               ev.events=EPOLLIN|EPOLLET;  
               epoll_ctl(epfd,EPOLL_CTL_ADD,connfd,&ev); //将新的fd添加到epoll的监听队列中  
           }  
  
           else if( events[i].events&EPOLLIN ) //接收到数据，读socket  
           {  
               n = read(sockfd, line, MAXLINE)) < 0    //读  
               ev.data.ptr = md;     //md为自定义类型，添加数据  
               ev.events=EPOLLOUT|EPOLLET;  
               epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);//修改标识符，等待下一个循环时发送数据，异步处理的精髓  
           }  
           else if(events[i].events&EPOLLOUT) //有数据待发送，写socket  
           {  
               struct myepoll_data* md = (myepoll_data*)events[i].data.ptr;    //取数据  
               sockfd = md->fd;  
               send( sockfd, md->ptr, strlen((char*)md->ptr), 0 );        //发送数据  
               ev.data.fd=sockfd;  
               ev.events=EPOLLIN|EPOLLET;  
               epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev); //修改标识符，等待下一个循环时接收数据  
           }  
           else  
           {  
               //其他的处理  
           }  
       }  
   }  
</pre>

##epoll的程序实例(C++)
<pre>
#include <stdio.h>  
#include <stdlib.h>  
#include <unistd.h>  
#include <errno.h>  
#include <sys/socket.h>  
#include <netdb.h>  
#include <fcntl.h>  
#include <sys/epoll.h>  
#include <string.h>  
  
#define MAXEVENTS 64  
  
//函数:  
//功能:创建和绑定一个TCP socket  
//参数:端口  
//返回值:创建的socket  
static int  
create_and_bind (char *port)  
{  
  struct addrinfo hints;  
  struct addrinfo *result, *rp;  
  int s, sfd;  
  
  memset (&hints, 0, sizeof (struct addrinfo));  
  hints.ai_family = AF_UNSPEC;     /* Return IPv4 and IPv6 choices */  
  hints.ai_socktype = SOCK_STREAM; /* We want a TCP socket */  
  hints.ai_flags = AI_PASSIVE;     /* All interfaces */  
  
  s = getaddrinfo (NULL, port, &hints, &result);  
  if (s != 0)  
    {  
      fprintf (stderr, "getaddrinfo: %s\n", gai_strerror (s));  
      return -1;  
    }  
  
  for (rp = result; rp != NULL; rp = rp->ai_next)  
    {  
      sfd = socket (rp->ai_family, rp->ai_socktype, rp->ai_protocol);  
      if (sfd == -1)  
        continue;  
  
      s = bind (sfd, rp->ai_addr, rp->ai_addrlen);  
      if (s == 0)  
        {  
          /* We managed to bind successfully! */  
          break;  
        }  
  
      close (sfd);  
    }  
  
  if (rp == NULL)  
    {  
      fprintf (stderr, "Could not bind\n");  
      return -1;  
    }  
  
  freeaddrinfo (result);  
  
  return sfd;  
}  
  
  
//函数  
//功能:设置socket为非阻塞的  
static int  
make_socket_non_blocking (int sfd)  
{  
  int flags, s;  
  
  //得到文件状态标志  
  flags = fcntl (sfd, F_GETFL, 0);  
  if (flags == -1)  
    {  
      perror ("fcntl");  
      return -1;  
    }  
  
  //设置文件状态标志  
  flags |= O_NONBLOCK;  
  s = fcntl (sfd, F_SETFL, flags);  
  if (s == -1)  
    {  
      perror ("fcntl");  
      return -1;  
    }  
  
  return 0;  
}  
  
//端口由参数argv[1]指定  
int  
main (int argc, char *argv[])  
{  
  int sfd, s;  
  int efd;  
  struct epoll_event event;  
  struct epoll_event *events;  
  
  if (argc != 2)  
    {  
      fprintf (stderr, "Usage: %s [port]\n", argv[0]);  
      exit (EXIT_FAILURE);  
    }  
  
  sfd = create_and_bind (argv[1]);  
  if (sfd == -1)  
    abort ();  
  
  s = make_socket_non_blocking (sfd);  
  if (s == -1)  
    abort ();  
  
  s = listen (sfd, SOMAXCONN);  
  if (s == -1)  
    {  
      perror ("listen");  
      abort ();  
    }  
  
  //除了参数size被忽略外,此函数和epoll_create完全相同  
  efd = epoll_create1 (0);  
  if (efd == -1)  
    {  
      perror ("epoll_create");  
      abort ();  
    }  
  
  event.data.fd = sfd;  
  event.events = EPOLLIN | EPOLLET;//读入,边缘触发方式  
  s = epoll_ctl (efd, EPOLL_CTL_ADD, sfd, &event);  
  if (s == -1)  
    {  
      perror ("epoll_ctl");  
      abort ();  
    }  
  
  /* Buffer where events are returned */  
  events = calloc (MAXEVENTS, sizeof event);  
  
  /* The event loop */  
  while (1)  
    {  
      int n, i;  
  
      n = epoll_wait (efd, events, MAXEVENTS, -1);  
      for (i = 0; i < n; i++)  
        {  
          if ((events[i].events & EPOLLERR) ||  
              (events[i].events & EPOLLHUP) ||  
              (!(events[i].events & EPOLLIN)))  
            {  
              /* An error has occured on this fd, or the socket is not 
                 ready for reading (why were we notified then?) */  
              fprintf (stderr, "epoll error\n");  
              close (events[i].data.fd);  
              continue;  
            }  
  
          else if (sfd == events[i].data.fd)  
            {  
              /* We have a notification on the listening socket, which 
                 means one or more incoming connections. */  
              while (1)  
                {  
                  struct sockaddr in_addr;  
                  socklen_t in_len;  
                  int infd;  
                  char hbuf[NI_MAXHOST], sbuf[NI_MAXSERV];  
  
                  in_len = sizeof in_addr;  
                  infd = accept (sfd, &in_addr, &in_len);  
                  if (infd == -1)  
                    {  
                      if ((errno == EAGAIN) ||  
                          (errno == EWOULDBLOCK))  
                        {  
                          /* We have processed all incoming 
                             connections. */  
                          break;  
                        }  
                      else  
                        {  
                          perror ("accept");  
                          break;  
                        }  
                    }  
  
                                  //将地址转化为主机名或者服务名  
                  s = getnameinfo (&in_addr, in_len,  
                                   hbuf, sizeof hbuf,  
                                   sbuf, sizeof sbuf,  
                                   NI_NUMERICHOST | NI_NUMERICSERV);//flag参数:以数字名返回  
                                  //主机地址和服务地址  
  
                  if (s == 0)  
                    {  
                      printf("Accepted connection on descriptor %d "  
                             "(host=%s, port=%s)\n", infd, hbuf, sbuf);  
                    }  
  
                  /* Make the incoming socket non-blocking and add it to the 
                     list of fds to monitor. */  
                  s = make_socket_non_blocking (infd);  
                  if (s == -1)  
                    abort ();  
  
                  event.data.fd = infd;  
                  event.events = EPOLLIN | EPOLLET;  
                  s = epoll_ctl (efd, EPOLL_CTL_ADD, infd, &event);  
                  if (s == -1)  
                    {  
                      perror ("epoll_ctl");  
                      abort ();  
                    }  
                }  
              continue;  
            }  
          else  
            {  
              /* We have data on the fd waiting to be read. Read and 
                 display it. We must read whatever data is available 
                 completely, as we are running in edge-triggered mode 
                 and won't get a notification again for the same 
                 data. */  
              int done = 0;  
  
              while (1)  
                {  
                  ssize_t count;  
                  char buf[512];  
  
                  count = read (events[i].data.fd, buf, sizeof(buf));  
                  if (count == -1)  
                    {  
                      /* If errno == EAGAIN, that means we have read all 
                         data. So go back to the main loop. */  
                      if (errno != EAGAIN)  
                        {  
                          perror ("read");  
                          done = 1;  
                        }  
                      break;  
                    }  
                  else if (count == 0)  
                    {  
                      /* End of file. The remote has closed the 
                         connection. */  
                      done = 1;  
                      break;  
                    }  
  
                  /* Write the buffer to standard output */  
                  s = write (1, buf, count);  
                  if (s == -1)  
                    {  
                      perror ("write");  
                      abort ();  
                    }  
                }  
  
              if (done)  
                {  
                  printf ("Closed connection on descriptor %d\n",  
                          events[i].data.fd);  
  
                  /* Closing the descriptor will make epoll remove it 
                     from the set of descriptors which are monitored. */  
                  close (events[i].data.fd);  
                }  
            }  
        }  
    }  
  free (events);  
  
  close (sfd);  
  
  return EXIT_SUCCESS;  
}  
</pre>
运行方式：

- 在一个终端运行此程序：epoll.out PORT
- 另一个终端：telnet  127.0.0.1 PORT