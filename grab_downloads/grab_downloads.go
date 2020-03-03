package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

const (
	grabURL   = "https://vol.moe/"
	loginURL  = "https://vol.moe/login_do.php"
	email     = "1248210537@qq.com"
	passwd    = "1996617zjxD"
	keepalive = "on"
)

var searchInfo string

func main() {
	start()
}

// start is a function to start the app
func start() {
	search()
}

func search() {
	fmt.Println("Please input the comic that you want to search ...")
	fmt.Scanln(&searchInfo)
	searchURL := grabURL + "list.php?s=" + searchInfo
	grabInit(searchURL)
}

func grabInit(URL string) {
	c := colly.NewCollector()

	// login
	err := c.Post(loginURL, map[string]string{"email": email, "passwd": passwd, "keepalive": keepalive})
	if err != nil {
		log.Fatal(err)
	}

	// proxy
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7891")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// monitor
	c.OnHTML(".listbg td a[href]:has(img)", func(e *colly.HTMLElement) {
		name := e.ChildAttr("img", "alt")
		link := e.Attr("href")
		imgURL := e.ChildAttr("img", "src")
		fmt.Printf("Search Result = name: %s, link: %s, imgURL: %s \n", name, link, imgURL)
		if name == "一人之下" {
			c.Visit(link)
		}
	})

	c.OnHTML("#div_cbz td a[href]:contains(下載)", func(e *colly.HTMLElement) {
		fmt.Println("Downloads URLs", e.Attr("href"))
		// file, err := os.OpenFile(
		// 	"test.txt",
		// 	os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		// 	0666,
		// )
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file.Close()
		// // 写字节到文件中
		// byteSlice := []byte(e.Attr("href"))
		// bytesWritten, err := file.Write(byteSlice)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(bytesWritten)
	})
	// Called before the request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// called after the response
	// c.OnResponse(func(res *colly.Response) {
	// 	if res.StatusCode == 200 {
	// 		c.Visit(URL)
	// 	}
	// })
	c.Visit(URL)
}
