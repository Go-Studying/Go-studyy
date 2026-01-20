package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Go-Studying/Go-studyy/crawler"
)

const divideBar = "=============="

func main() {
	var confirm string
	useGoroutine := false
	fmt.Print("고루틴을 사용하시겠습니까? (y/N): ")
	fmt.Scanf("%s\n", &confirm)

	if strings.ToUpper(confirm) == "y" {
		useGoroutine = true
	}

	index, err := crawler.Indexing()
	if err != nil {
		panic(err)
	}

	fmt.Println(divideBar)
	fmt.Printf("총 %d건의 뉴스 기사를 찾았습니다.\n", len(index.Link))
	fmt.Printf("뉴스 내용을 불러옵니다")

	start := time.Now()
	err = crawler.Crawler(index, useGoroutine)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("모든 뉴스 내용을 불러왔습니다.")
	fmt.Println(divideBar)
	for _, link := range index.Link {
		fmt.Printf("TITLE: %v\n", link.Title)
		fmt.Printf("LINK : %v\n", link.URI)
		fmt.Printf("CONTENT: %d 글자의 내용\n", len(link.Content))
		fmt.Println(divideBar)
	}
	fmt.Printf("소요시간: %v\n", elapsed)
}
