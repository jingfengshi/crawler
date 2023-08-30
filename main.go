package main

import (
	"fmt"
	"github.com/jasperjing/crawler/collect"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.Fetcher = collect.BrowserFetch{}

	body, err := f.Get(url)

	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}

	fmt.Printf("%s\n", string(body))

}
