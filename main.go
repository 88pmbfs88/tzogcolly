package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"strconv"
	"tzogcolly/data"
)

func main() {

	var listPages []data.ListPage

	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
	//文章列表
	c.OnHTML("table[class='TABLE1'] > tbody", func(e *colly.HTMLElement) {
		//列表中每一项
		e.ForEach("tr", func(i int, item *colly.HTMLElement) {
			//id
			id := item.ChildText("td:nth-child(1)")
			//文章链接
			href := item.ChildAttr("td:nth-child(2) > div > a", "href")
			//文章标题
			title := item.ChildText("td:nth-child(2) > div > a")
			//文章类型
			cat := item.ChildText("td:nth-child(3) > div > div > a")
			//文章难度
			easy := item.ChildText("td:nth-child(5) > div ")

			idInt, err := strconv.Atoi(id)
			if err != nil {
				idInt = 0
			}
			easyInt, err := strconv.Atoi(easy)
			if err != nil {
				easyInt = 0
			}
			catLen := len(cat) / 4
			catTmp := []rune(cat)
			var catSince []string
			for j := 0; j < catLen; j++ {
				cat = string(catTmp[j*4:(j+1)*4])
				catSince = append(catSince, cat)
			}

			topicList := data.ListPage{
				PageId:     1,
				TopicId:    idInt,
				TopicHref:  href,
				TopicTitle: title,
				TopicCat:   catSince,
				TopicEasy:  easyInt,
			}
			fmt.Printf("%v\r\n",topicList)
			listPages = append(listPages, topicList)
			fmt.Println("---------------------")
		})
	})

	err := c.Visit("http://www.tzcoder.cn/acmhome/classify.do?method=show&type=19&page=1")
	if err != nil {
		fmt.Println(err.Error())
	}
}
