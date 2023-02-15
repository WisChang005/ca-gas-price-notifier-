package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/juunini/simple-go-line-notify/notify"
)

func main() {
	c := colly.NewCollector()
	var notifyMessage string

	c.OnHTML("li[data-location=\"OTTAWA\"]", func(e *colly.HTMLElement) {
		dToday := e.Attr("data-today")
		dYesterday := e.Attr("data-yesterday")
		dLastWeek := e.Attr("data-lastweek")
		dArrow := getArrowSign(e.Attr("data-arrow"))

		notifyMessage = fmt.Sprintf(`
Gas Prise in Ottawa
Today: %v %v
Yesterday: %v
Lastweek: %v`, dArrow, dToday, dYesterday, dLastWeek)
		fmt.Println(notifyMessage)
	})

	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
	})

	err := c.Visit("https://www.caa.ca/gas-prices/")
	if err != nil {
		panic(err)
	}

	notifyToLine(notifyMessage)
}

func notifyToLine(msg string) {
	accessToken := os.Getenv("LINE_TOKEN")
	if err := notify.SendText(accessToken, msg); err != nil {
		panic(err)
	}
	fmt.Println("Notify to LINE completed!")
}

func getArrowSign(arrowName string) string {
	var arrow string
	switch arrowName {
	case "up":
		arrow = "↟"
	case "down":
		arrow = "↡"
	case "equal":
		arrow = "↭"
	default:
		arrow = "?"
	}
	return arrow
}
