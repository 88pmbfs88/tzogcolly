package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"strconv"
	"tzogcolly/data"
	"tzogcolly/model"
)

func main() {

	pageIndex := 3

	var listPages []data.ListPage

	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
	c2 := c.Clone()
	//c2.Async = true
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
				return
			}
			easyInt, err := strconv.Atoi(easy)
			if err != nil {
				easyInt = 0
			}
			//catLen := len(cat) / 4
			//catTmp := []rune(cat)
			//var catSince []string
			//for j := 0; j < catLen; j++ {
			//	cat = string(catTmp[j*4:(j+1)*4])
			//	catSince = append(catSince, cat)
			//}

			topicList := data.ListPage{
				PageId:     pageIndex,
				TopicId:    idInt,
				TopicHref:  href,
				TopicTitle: title,
				//TopicCat:   catSince,
				TopicCat:  cat,
				TopicEasy: easyInt,
			}
			fmt.Printf("%v\r\n", topicList)
			listPages = append(listPages, topicList)
			model.InsertTopicList(topicList)
			ctx := colly.NewContext()
			ctx.Put("id", id)
			//通过Context上下文对象将采集器1采集到的数据传递到采集器2
			c2.Request("GET", "http://www.tzcoder.cn"+href, nil, ctx, nil)
			fmt.Println("---------------------")
		})
	})

	//采集器2，获取文章详情
	c2.OnHTML("body", func(e *colly.HTMLElement) {
		topicDetailDesc := e.ChildText("table:nth-of-type(2) [align='left'] > div:nth-of-type(1) ")
		topicDetailInput := e.ChildText("table:nth-of-type(2) div:nth-of-type(2) ")
		topicDetailOutput := e.ChildText("table:nth-of-type(2) div:nth-of-type(3) ")
		topicDetailDemoInput := e.ChildText("#sample_input")
		topicDetailDemoOutput := e.ChildText("#sample_output")
		print("！！！！！！！！！！！！！！！！！！！")
		fmt.Println(topicDetailDesc)
		fmt.Println(topicDetailInput)
		fmt.Println(topicDetailOutput)
		fmt.Println(topicDetailDemoInput)
		fmt.Println(topicDetailDemoOutput)
		print("！！！！！！！！！！！！！！！！！！！")
		id := e.Request.Ctx.Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			idInt = 0
		}
		topicDetail := data.TopicPage{
			TopicId:            idInt,
			TopicDesc:          topicDetailDesc,
			TopicInput:         topicDetailInput,
			TopicOutput:        topicDetailOutput,
			TopicExampleInput:  topicDetailDemoInput,
			TopicExampleOutput: topicDetailDemoOutput,
		}
		model.InsertTopicDetail(topicDetail)
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("c2爬取页面：", r.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("c1爬取页面：", r.URL)
	})

	err := model.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v", err)
		return
	}

	err = c.Visit("http://www.tzcoder.cn/acmhome/classify.do?method=show&type=19&page=" + strconv.Itoa(pageIndex))
	if err != nil {
		fmt.Println(err.Error())
	}
}
