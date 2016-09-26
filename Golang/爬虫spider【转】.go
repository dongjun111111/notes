package main

import (
	"fmt"
	. "github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"labix.org/v2/mgo"
	"strings"
	"sync"
	"time"
)

var CACHEMAP map[string]bool = make(map[string]bool)

type Joke struct {
	Title, Content, Src, ImgSrc, ImgLocal, ThumbSrc, ThumbLocal   string
	Id, Date, Likes, ImgWidth, ImgHeight, ThumbWidth, ThumbHeight int64
}

//var COLLECT *mgo.Collection

// ---------------------------------------- Scheduler ----------------------------------------
type Scheduler struct {
	Crawlers map[string]ICrawler
	Locker   *sync.Mutex
	Queue    chan interface{}
}

func (this *Scheduler) Add(crawler ICrawler) {
	name := crawler.GetName()
	if _, ok := this.Crawlers[name]; ok != true {
		this.Crawlers[name] = crawler
	}
}

func (this *Scheduler) Run() {
	//var wg sync.WaitGroup
	for _, crawler := range this.Crawlers {
		//wg.Add(1)
		go func(c ICrawler) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
					//wg.Done()
				}
			}()
			c.Run()
			if c.Valid() {
				c.Visit(this.Queue, this.Locker)
				//wg.Done()
			}
		}(crawler)
	}
	//wg.Wait()
}

func (this *Scheduler) Loop() {
	collection, session := connectMongo()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			spew.Println(err)
		}
	}()
	for {
		select {
		case item := <-this.Queue:
			switch v := item.(type) {
			case Joke:
				spew.Println("[+] insert : ", v)
				v.Date = time.Now().Unix()
				v.Id = time.Now().UnixNano()
				spew.Println(v.Date, v.Id)
				collection.Insert(v)

			default:
				spew.Println("sth wrong?")
			}
			spew.Println("----------------------------------------------------------------------------")

			//case for other chan receive
			//...
		}
	}
}

func connectMongo() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session.DB("xiaohua").C("jokes"), session
}

// ---------------------------------------- Crawler ----------------------------------------

type CrawlerConfig struct {
	Url, Type, Host string
	isGB            bool
}

type CrawlerBase struct {
	Name string
	Cfg  *CrawlerConfig
	doc  *Document
}

func (this *CrawlerBase) GetName() string {
	return this.Name
}

func (this *CrawlerBase) GetDoc() *Document {
	return this.doc
}

func (this *CrawlerBase) GetType() string {
	return this.Cfg.Type
}

func (this *CrawlerBase) Visit(Q chan interface{}, L *sync.Mutex) {
	// do sth!
	fmt.Printf("[+] visit : ============== %s(%s) ==============\n", this.Name, this.Cfg.Url)
	if this.doc == nil {
		fmt.Printf("[-] http request failed, skip(%s)\n", this.Name)
		return
	}
	con, _ := this.doc.Find("meta").Eq(0).Attr("content")
	con = strings.ToUpper(con)
	if strings.Contains(con, "GBK") || strings.Contains(con, "GB2312") {
		this.Cfg.isGB = true
	}
}

func (this *CrawlerBase) Run() {
	var doc *Document
	var e error
	if doc, e = NewDocument(this.Cfg.Url); e != nil {
		this.doc = nil
		return
	}
	this.doc = doc
}

func (this *CrawlerBase) Valid() bool {
	doc := this.GetDoc()
	if doc == nil {
		return false
	}
	return true
}

// ------------------------------------ Crawler interface ------------------------------------

type ICrawler interface {
	GetName() string
	GetDoc() *Document
	GetType() string
	Valid() bool
	Visit(chan interface{}, *sync.Mutex)
	Run()
}

// ---------------------------------------- 百思不得姐（文字） ----------------------------------------

//http://www.budejie.com/duanzi/

type budejieCrawler struct {
	CrawlerBase
}

func (this *budejieCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	//spew.Dump(CACHEMAP)

	this.doc.Find(".web_size").Each(func(i int, s *Selection) {
		//title-5595718, detail-5595718.html
		src, ok := s.Attr("id")
		if !ok {
			src = this.doc.Url.String()
		} else {
			src = this.Cfg.Host + strings.Replace(src, "title", "detail", -1) + ".html"
		}
		//spew.Println("!!!!!!!!!!!!!!!!!!!!!!!!", src)
		cont := strings.TrimSpace(s.Text())
		_, ok = CACHEMAP[src]
		//spew.Println(ok)
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}
		joke := Joke{
			Content: cont,
			Src:     src,
		}
		Q <- joke
	})
	//spew.Dump(joke)
}

// ---------------------------------------- 百思不得姐（图文） ----------------------------------------

//http://www.budejie.com

type budejiePicCrawler struct {
	CrawlerBase
}

func (this *budejiePicCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	//spew.Dump(CACHEMAP)

	this.doc.Find(".web_size").Each(func(i int, s *Selection) {
		//title-5595718, detail-5595718.html
		src, ok := s.Attr("id")
		if !ok {
			src = this.doc.Url.String()
		} else {
			src = this.Cfg.Host + strings.Replace(src, "title", "detail", -1) + ".html"
		}
		//spew.Println("!!!!!!!!!!!!!!!!!!!!!!!!", src)
		cont := strings.TrimSpace(s.Text())
		_, ok = CACHEMAP[src]
		//spew.Println(ok)
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}
		imgsrc, _ := s.Next().Find("img").Eq(0).Attr("src")
		//spew.Println("!!!!!!!!!!", imgsrc)
		joke := Joke{
			Content: cont,
			Src:     src,
			ImgSrc:  imgsrc,
		}
		Q <- joke
	})
	//spew.Dump(joke)
}

// ---------------------------------------- 糗事百科 ----------------------------------------
//http://www.qiushibaike.com/

type qiubaiCrawler struct {
	CrawlerBase
}

func (this *qiubaiCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	//spew.Dump(CACHEMAP)
	this.doc.Find(".content").Each(func(i int, s *Selection) {
		src, ok := s.Parent().Attr("id")
		if !ok {
			src = this.doc.Url.String()
		} else {
			src = this.Cfg.Host + strings.Replace(src, "qiushi_tag_", "article/", -1)
		}
		cont := strings.TrimSpace(s.Text())
		_, ok = CACHEMAP[src]
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}

		imgsrc := ""
		if thumb := s.Next(); thumb != nil {
			imgsrc, _ = thumb.Find("img").Eq(0).Attr("src")
		}
		joke := Joke{
			Content: cont,
			Src:     src,
			ImgSrc:  imgsrc,
		}
		Q <- joke
	})
	//spew.Dump(joke)
}

// ---------------------------------------- 内涵段子 ----------------------------------------
//http://www.neihanshequ.com/

type neihanduanziCrawler struct {
	CrawlerBase
}

func (this *neihanduanziCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	this.doc.Find(".share_url").Each(func(i int, s *Selection) {
		src, _ := s.Attr("href")
		src = this.Cfg.Host + src[1:]
		imgsrc := ""
		if img := s.Find("img"); img != nil {
			imgsrc, _ = img.Attr("data-original")
		}
		cont := strings.TrimSpace(s.Text())
		_, ok := CACHEMAP[src]
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}
		joke := Joke{
			Content: cont,
			Src:     src,
			ImgSrc:  imgsrc,
		}
		Q <- joke
		//spew.Printf("%d: %s\n", i, band)
	})
}

// ---------------------------------------- 来福岛 ----------------------------------------
//http://www.laifudao.com/

type laifuCrawler struct {
	CrawlerBase
}

func (this *laifuCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	this.doc.Find(".post-article").Each(func(i int, s *Selection) {
		titleLink := s.Find("a").Eq(0)
		src, _ := titleLink.Attr("href")
		src = this.Cfg.Host + src[1:]
		title := titleLink.Text()

		cont := ""
		if art := s.Find(".article-content"); art != nil {
			cont = strings.TrimSpace(art.Text())
		}

		imgsrc := ""
		if pic := s.Find(".pic-content"); pic != nil {
			imgsrc, _ = pic.Find("img").Eq(0).Attr("src")
		}

		_, ok := CACHEMAP[src]
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}
		joke := Joke{
			Title:   title,
			Content: cont,
			Src:     src,
			ImgSrc:  imgsrc,
		}
		Q <- joke
		//spew.Printf("%d: %s\n", i, band)
	})
}

// ---------------------------------------- 九妖内涵图 ----------------------------------------
//http://www.9yao.com

type yaoCrawler struct {
	CrawlerBase
}

func (this *yaoCrawler) Visit(Q chan interface{}, L *sync.Mutex) {
	//spew.Dump(this)
	this.CrawlerBase.Visit(Q, L)
	this.doc.Find(".box").Each(func(i int, s *Selection) {
		titleLink := s.Find("a").Eq(0)
		src, _ := titleLink.Attr("href")
		src = this.Cfg.Host + src[1:]
		title := titleLink.Text()
		imgsrc := ""
		if pic := s.Find(".box-content"); pic != nil {
			img := pic.Find("img").Eq(0)
			mySrc, ok := img.Attr("data-original")
			if !ok {
				mySrc, _ = img.Attr("src")
			}
			imgsrc = mySrc
		}

		_, ok := CACHEMAP[src]
		if ok {
			spew.Println("[!] Skip!")
			return
		} else {
			L.Lock()
			CACHEMAP[src] = true
			L.Unlock()
		}
		joke := Joke{
			Title:  title,
			Src:    src,
			ImgSrc: imgsrc,
		}
		Q <- joke
		//spew.Printf("%d: %s\n", i, band)
	})
}

// ---------------------------------------- 捧腹网 ----------------------------------------
//http://www.pengfu.com/xiaohua_1.html

// ---------------------------------------- 哈哈MX ----------------------------------------
//http://www.haha.mx/

func main() {
	myScheduler := &Scheduler{
		Crawlers: make(map[string]ICrawler),
		Locker:   new(sync.Mutex),
		Queue:    make(chan interface{}, 10),
	}

	myScheduler.Add(&budejieCrawler{
		CrawlerBase{
			Name: "百思不得姐-内涵段子",
			Cfg: &CrawlerConfig{
				Host: "http://www.budejie.com/",
				Url:  "http://www.budejie.com/duanzi/", Type: "JOKE",
			},
		},
	})
	myScheduler.Add(&budejiePicCrawler{
		CrawlerBase{
			Name: "百思不得姐-搞笑图片",
			Cfg: &CrawlerConfig{
				Host: "http://www.budejie.com/",
				Url:  "http://www.budejie.com/", Type: "PIC",
			},
		},
	})

	myScheduler.Add(&qiubaiCrawler{
		CrawlerBase{
			Name: "糗事百科",
			Cfg: &CrawlerConfig{
				Host: "http://www.qiushibaike.com/",
				Url:  "http://www.qiushibaike.com/", Type: "JOKE",
			},
		},
	})

	myScheduler.Add(&neihanduanziCrawler{
		CrawlerBase{
			Name: "内涵段子",
			Cfg: &CrawlerConfig{
				Host: "http://www.neihanshequ.com/",
				Url:  "http://www.neihanshequ.com/", Type: "PIC",
			},
		},
	})

	myScheduler.Add(&laifuCrawler{
		CrawlerBase{
			Name: "来福岛",
			Cfg: &CrawlerConfig{
				Host: "http://www.laifudao.com/",
				Url:  "http://www.laifudao.com/", Type: "PIC",
			},
		},
	})

	myScheduler.Add(&yaoCrawler{
		CrawlerBase{
			Name: "九妖内涵图",
			Cfg: &CrawlerConfig{
				Host: "http://www.9yao.com/",
				Url:  "http://www.9yao.com/", Type: "PIC",
			},
		},
	})

	go func() {
		for {
			myScheduler.Run()
			spew.Println("[+] Zzz~")
			time.Sleep(time.Second * 60 * 5)
		}
	}()
	myScheduler.Loop()
}
