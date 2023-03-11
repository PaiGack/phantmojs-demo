package main

import (
	"net/http"
	"os"

	"github.com/benbjohnson/phantomjs"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
		}
	}()

	var url = "https://www.163.com"
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer phantomjs.DefaultProcess.Close()

	page, err := phantomjs.CreateWebPage()
	if err != nil {
		panic(err)
	}

	requestHader := http.Header{"User-AgentF": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63"}}

	if err := page.SetCustomHeaders(requestHader); err != nil {
		panic(err)
	}

	if err := page.SetViewportSize(640, 960); err != nil {
		panic(err)
	}

	if err := page.Open(url); err != nil {
		panic(err)
	}

	if err := page.Render("hackernews4.png", "png", 50); err != nil {
		panic(err)
	}
}
