package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
)

const YahooFinance = "https://finance.yahoo.com"

const Apple = "AAPL"
const Amazon = "AMZN"

// 구글 파이낸스 애플 기사 스크래핑(괄호열고 크롤링 괄호닫고) - 영빈이가 지랄할수도 있으니
func main() {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// 크롤링 할 주소
	url := fmt.Sprintf("%s/quote/%s/news?p=%s", YahooFinance, Apple, Apple)

	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		panic(err)
	}

	scrollScript := `
	   window.scrollTo(0, document.body.scrollHeight);
	`

	for i := 0; i < 2; i++ {
		if err := chromedp.Run(ctx, chromedp.Evaluate(scrollScript, nil)); err != nil {
			panic(err)
		}
	}

	if err := chromedp.Run(ctx, chromedp.Nodes(".js-content-viewer", &nodes, chromedp.ByQueryAll)); err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		//fmt.Println(node)
		var title, url string
		if err := chromedp.Run(ctx,
			chromedp.Text(node.FullXPath(), &title, chromedp.NodeVisible),
		); err != nil {

			panic(err)
		}

		// 임시로 만듬..
		for _, attribute := range node.Attributes {
			if attribute[0] == '/' {
				url = attribute
			}
		}

		fmt.Println("제목: ", title)
		fmt.Println(fmt.Sprintf("기사 url: %s%s", YahooFinance, url))
	}

}
