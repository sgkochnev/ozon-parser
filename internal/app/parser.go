package app

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"ozon-parser/model"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"

func parseOzon(_url string) ([]model.Goods, error) {
	doc, err := getHTMLNode(_url)
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, "//*[@class='uj uj0']")

	googs := []model.Goods{}
	g := model.Goods{}
	for _, v := range list {
		priceNode := htmlquery.FindOne(v, "//span[@class='ui-q5 ui-q9 ui-r1']")
		if priceNode == nil {
			priceNode = htmlquery.FindOne(v, "//span[@class='ui-q5 ui-q9']")
		}
		nameNode := htmlquery.FindOne(v, "//span[@class='md8 dm9 m9d dn1 tsBodyL js4']/*")
		urlIMGNode := htmlquery.FindOne(v, "//div[@class='js9']/img")
		urlNode := htmlquery.FindOne(v, "//a@href")
		price := htmlquery.InnerText(priceNode)

		g.Price = convertPriceToInt(price)
		g.Name = htmlquery.InnerText(nameNode)
		g.Url = baseURL + htmlquery.SelectAttr(urlNode, "href")
		g.UrlIMG = htmlquery.SelectAttr(urlIMGNode, "src")
		googs = append(googs, g)
	}
	return googs, err
}

func convertPriceToInt(price string) int {
	var p int
	price = strings.ReplaceAll(price, fmt.Sprintf("%c", 0x2009), "")
	fmt.Fscanf(strings.NewReader(price), "%d", &p)

	return p
}

func getHTMLNode(_url string) (*html.Node, error) {
	return htmlquery.Parse(getHtml(_url))
}

func getHtml(_url string) io.Reader {
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", userAgent)
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil && data == nil {
		log.Fatalln(err)
	}
	return bytes.NewReader(data)
}
