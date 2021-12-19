package data

type ListPage struct {
	PageId     int
	TopicId    int
	TopicHref  string
	TopicTitle string
	//TopicCat   []string
	TopicCat  string
	TopicEasy int
}

type TopicPage struct {
	TopicId            int
	TopicDesc          string
	TopicInput         string
	TopicOutput        string
	TopicExampleInput  string
	TopicExampleOutput string
}
