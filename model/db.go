package model

import (
	"database/sql"
	"fmt"
	"tzogcolly/data"
)

// 定义一个全局对象db
var Db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/tzog_page"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 插入列表数据
func InsertTopicList(data data.ListPage) {
	sqlStr := "insert into topic_list(page_id, topic_id,topic_title,topic_href,topic_cat,topic_easy) values (?,?,?,?,?,?)"
	ret, err := Db.Exec(sqlStr, data.PageId, data.TopicId, data.TopicTitle, data.TopicHref, data.TopicCat, data.TopicEasy)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 插入列表数据
func InsertTopicDetail(data data.TopicPage) {
	sqlStr := "insert into topic_detail(topic_id, topic_desc,topic_input,topic_output,topic_demo_input,topic_demo_output) values (?,?,?,?,?,?)"
	ret, err := Db.Exec(sqlStr, data.TopicId, data.TopicDesc, data.TopicInput, data.TopicOutput, data.TopicExampleInput, data.TopicExampleOutput)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
