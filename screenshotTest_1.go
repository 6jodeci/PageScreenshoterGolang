package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://afisha.yandex.ru/info/cinema`, `#main`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	if err := chromedp.Run(ctx, fullScreenshot(`https://afisha.yandex.ru/info/cinema`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
}

func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}