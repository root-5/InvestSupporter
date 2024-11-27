// 主にスプレッドシートからの利用を想定したAPIを提供する
package scraping

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// const OUTPUT_DIR = "./output"

func GetAnnounceDate() [][]string {
	maxPageNum := 10

	var resultData [][]string

	for pageNum := 1; pageNum <= maxPageNum; pageNum++ {
		url := fmt.Sprintf("https://www.traders.co.jp/market_jp/earnings_calendar/all/all/%d?term=future", pageNum)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			return nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil
		}

		doc, err := html.Parse(strings.NewReader(string(body)))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			return nil
		}

		var tdEles []string
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "td" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						tdEles = append(tdEles, c.Data)
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)

		if len(tdEles) == 0 {
			break
		}

		var tdArray []string
		for _, td := range tdEles {
			tdArray = append(tdArray, strings.TrimSpace(td))
		}
		fmt.Println(tdArray)

		column := 7
		var tdSquareArray [][]string
		for i := 0; i < len(tdArray); i += column {
			tdSquareArray = append(tdSquareArray, tdArray[i:i+column])
		}
		fmt.Println(tdSquareArray)

		for _, tdRow := range tdSquareArray {
			tdRow = tdRow[:len(tdRow)-1]
			tdRow[5] = strings.ReplaceAll(tdRow[5], ",", "")
			tdRow[2] = strings.ReplaceAll(tdRow[2], "\n", "")
			tdRow[2] = strings.ReplaceAll(tdRow[2], " ", "")

			stockName := tdRow[2][:strings.Index(tdRow[2], "(")]
			stockCode := tdRow[2][strings.Index(tdRow[2], "(")+1 : strings.Index(tdRow[2], ")")]
			market := tdRow[2][strings.Index(tdRow[2], "/")+1 : strings.Index(tdRow[2], "/")+3]

			tdRow = append(tdRow[:3], append([]string{market, stockCode, stockName}, tdRow[3:]...)...)
			tdRow = append(tdRow[:2], tdRow[3:]...)
			resultData = append(resultData, tdRow)
		}
	}

	return resultData

	// today := time.Now()
	// year := today.Year()
	// month := fmt.Sprintf("%02d", int(today.Month()))
	// date := fmt.Sprintf("%02d", today.Day())
	// fileName := fmt.Sprintf("%s/決算発表日一覧_%d%s%s.csv", OUTPUT_DIR, year, month, date)

	// file, err := os.Create(fileName)
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// for _, row := range resultData {
	// 	if err := writer.Write(row); err != nil {
	// 		fmt.Println("Error writing to CSV:", err)
	// 		return
	// 	}
	// }
}
