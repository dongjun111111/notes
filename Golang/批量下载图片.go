package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	_ "time"
)

const (
	bufferSize = 128 * 1024 //写图片文件的缓冲区大小
)

var (
	numPoller     = flag.Int("p", 2, "page loader num")
	numDownloader = flag.Int("d", 5, "image downloader num")
	savePath      = flag.String("s", "./downloads/", "save path")
	imgExp        = regexp.MustCompile(`<a\s+class="img"\s+href="[a-zA-Z0-9_\-/:\.%?=]+">[\r\n\t\s]*<img\s+src="([^"'<>]*)"\s*/?>`)
	img2Exp       = regexp.MustCompile(`<a href="(.*)" target="_blank" class="view_img_link">`)
)

type image struct {
	url      string
	filename string
}

type sexyContext struct {
	pollerDone   chan struct{}
	images       map[string]int
	imagesLock   *sync.Mutex
	imageChan    chan *image
	pageIndex    int32
	rootURL      string
	done         bool
	imageCounter int32
	okCounter    int32
}

func main() {
	flag.Parse()
	ctx := &sexyContext{
		pollerDone: make(chan struct{}),
		images:     make(map[string]int),
		imagesLock: &sync.Mutex{},
		imageChan:  make(chan *image, 100),
		pageIndex:  1,
		rootURL:    "http://jandan.net/ooxx",
	}
	os.MkdirAll(*savePath, 0777)
	ctx.start()

}

func (ctx *sexyContext) start() {
	for i := 0; i < *numPoller; i++ {
		go ctx.downloadPage()
	}
	waits := sync.WaitGroup{}
	for i := 0; i < *numDownloader; i++ {
		go func() {
			waits.Add(1)
			ctx.downloadImage()
			waits.Done()
		}()
	}

	<-ctx.pollerDone
	ctx.done = true
	//close(ctx.pollerDone)
	waits.Wait()
	fmt.Printf("fetch done get %d ok %d\n", ctx.imageCounter, ctx.okCounter)
}

func (ctx *sexyContext) downloadPage() {
	isDone := false
	for !isDone {
		select {
		case <-ctx.pollerDone:
			isDone = true
		default:
			url := fmt.Sprintf("%s/page-%d", ctx.rootURL, atomic.AddInt32(&ctx.pageIndex, 1))
			fmt.Printf("download page %s\n", url)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("failed to load url %s with error %v", url, err)
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("failed to load url %s with error %v", url, err)
				} else {

					ctx.parsePage(body)
				}
			}
		}
	}

}
func GetUrl(url string) []byte {
	ret, err := http.Get(url)
	if err != nil {
		status := map[string]string{}
		status["status"] = "400"
		status["url"] = url
		panic(status)
	}
	body := ret.Body
	data, _ := ioutil.ReadAll(body)
	return data
}

func (ctx *sexyContext) parsePage(body []byte) {
	//fmt.Printf("%s\n", string(body))
	body2 := string(body)
	idx := img2Exp.FindAllStringSubmatch(body2, -1)
	fmt.Println("idx", idx)
	if idx == nil {
		ctx.pollerDone <- struct{}{}
	} else {
		for _, n := range idx {
			data := GetUrl(n[1])
			if len(data) > 10 {
				body := string(data)
				part := regexp.MustCompile(`<img src="(.*)"`)
				match := part.FindAllStringSubmatch(body, -1)
				fmt.Println("match")

				for _, v := range match {
					str := strings.Split(v[1], "/")
					length := len(str)
					imgeUrl := v[1]
					filename := str[length-1]
					image := &image{url: imgeUrl, filename: filename}
					//atomic.AddInt32(&ctx.imageCounter, 1)
					//ctx.imageChan <- image
					fmt.Printf("start download %s\n", image.url)
					atomic.AddInt32(&ctx.okCounter, 1)
					resp, err := http.Get(image.url)
					if err != nil {
						fmt.Printf("failed to load url %s with error %v\n", image.url, err)
					} else {
						defer resp.Body.Close()
						saveFile := *savePath + image.filename //path.Base(imgUrl)

						img, err := os.Create(saveFile)
						if err != nil {
							fmt.Print(err)

						} else {
							defer img.Close()

							imgWriter := bufio.NewWriterSize(img, bufferSize)

							_, err = io.Copy(imgWriter, resp.Body)
							if err != nil {
								fmt.Print(err)

							}
							imgWriter.Flush()
						}
					}

				}

			}
		}
	}
}
func (ctx *sexyContext) downloadImage() {
	isDone := false
	for !isDone {
		select {
		case <-ctx.pollerDone:
			if len(ctx.imageChan) == 0 {
				isDone = true
			}
		case image := <-ctx.imageChan:
			fmt.Printf("start download %s\n", image.url)
			atomic.AddInt32(&ctx.okCounter, 1)
			resp, err := http.Get(image.url)
			if err != nil {
				fmt.Printf("failed to load url %s with error %v\n", image.url, err)
			} else {
				defer resp.Body.Close()
				saveFile := *savePath + image.filename //path.Base(imgUrl)

				img, err := os.Create(saveFile)
				if err != nil {
					fmt.Print(err)

				} else {
					defer img.Close()

					imgWriter := bufio.NewWriterSize(img, bufferSize)

					_, err = io.Copy(imgWriter, resp.Body)
					if err != nil {
						fmt.Print(err)

					}
					imgWriter.Flush()
				}

			}
		}
	}

}
