package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var class = []string{"体育", "教育", "财经", "社会", "娱乐", "军事", "国内", "科技", "互联网", "房产", "国际", "女人", "汽车", "游戏"}

type Article struct {
	ID      int    `json:"-"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Comment string `json:"comment"`
	Link    string `json:"link"`
}

type Count struct {
	ID    int    `json:"-"`
	Link  string `json:"link"`
	Count int    `json:"count"`
}

type Tag struct {
	ID   int    `json:"-"`
	Link string `json:"link"`
	Tag  string `json:"tag"`
}

type Object struct {
	ID      int      `json:"-"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Comment string   `json:"comment"`
	Link    string   `json:"link"`
	Tag     []string `json:"tag"`
	Count   int      `json:"count"`
}

func TagsToString(tags []Tag, delim string) string {
	if len(tags) == 0 {
		return ""
	} else if len(tags) == 1 {
		if tags[0].Tag == "" {
			return ""
		} else {
			return "[" + tags[0].Tag + "]"
		}
	}
	ans := "["
	for i := 0; i < len(tags); i++ {
		if tags[i+1].Tag == "" {
			ans += "\"" + tags[i].Tag + "\""
			break;
		} else {
			ans += "\"" + tags[i].Tag + "\"" + delim
		}
	}
	ans += "]"
	return ans
}

func TagTOS(tags []Tag) []string {
	str := make([]string, len(tags))
	for i, value := range tags {
		str[i] = value.Tag
	}
	return str
}

// 保存分享的内容
func Send_share(summary, comment, link, title string) (o Object) {
	a, err := GetLink(link)
	if err != nil {
		//如果本来没有存该摘录，就创建
		a := CreateLink(summary, comment, link, title)
		o.Title = a.Title
		o.Summary = a.Summary
		o.Comment = a.Comment
		o.Link = a.Link
		//然后创建 tags 和 count
		tags, count := CreateTagAndCount(link, summary, 1)
		o.Tag = TagTOS(tags)
		o.Count = count
		return o
	} else {
		// 先创建第二个 Article
		a = CreateLink(summary, comment, link, title)
		//如果有该摘录，则更新对应的 count
		count := Count{}
		db.Where("link = ?", link).First(&count)
		db.Model(&count).Update("count", count.Count+1)
		//获取对应的标签
		tags := []Tag{}
		db.Where("link = ?", link).Find(&tags)
		o.Tag = TagTOS(tags)
		//同时赋值
		o.Title = a.Title
		o.Summary = a.Summary
		o.Comment = a.Comment
		o.Link = a.Link
		o.Count = count.Count + 1
		return o
	}
}

func GetLink(link string) (a Article, err error) {
	db.Where("link = ?", link).First(&a)
	if a.Link == "" {
		return a, errors.New("未找到 id 为 1 的对象")
	} else {
		return a, nil
	}
}

func httpDo(url string, summary []string) (ans string) {
	client := &http.Client{}

	payloadBytes, err := json.Marshal(summary)
	if err != nil {
		// handle err
	}
	b := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Token", "8OOZIpf5.34925.CodVj8Q7vxjb")

	//req.PostForm()
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i - 1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func GetTagOnline(link string, summary []string) (tags []Tag) {
	url := "http://api.bosonnlp.com/classify/analysis"
	tag := httpDo(url, summary)
	//fmt.Println(strs)
	//过滤掉相同的标签
	var i []int
	json.Unmarshal([]byte(tag), &i)
	sets := Duplicate(i)

	// create a secondary slice of ints, same length as selected
	ret := make([]int, len(sets))
	// copy one by one
	for i, x := range sets {
		ret[i] = x.(int) //provided it's indeed int. you can add a check here
	}
	tags = make([]Tag, len(ret))
	for i, value := range ret {
		tags[i].Link = link
		tags[i].Tag = class[value]
	}
	return tags
}

func CreateLink(summary, comment, link, title string) (a Article) {
	a = Article{
		Summary: summary,
		Comment: comment,
		Link:    link,
		Title:   title,
	}
	db.Create(&a)
	return a
}

func CreateTagAndCount(link string, summary string, count int) (tags []Tag, c int) {
	// 获取标签
	strs := strings.Split(summary, "。")
	tags = GetTagOnline(link, strs)
	// 插入到对应的数据库
	for _, value := range tags {
		db.Create(&value)
	}
	// 创建对应的计数标记
	db.Create(&Count{
		Link:  link,
		Count: count,
	})
	return tags, count
}

func GetHot(tag string) (o []Object) {
	//articles []Article
	var tags []Tag
	db.Where("tag Like ?", "%"+tag+"%").Find(&tags).Order("count")
	o = make([]Object, len(tags))
	for i, value := range tags {
		//查找标题等
		var article Article
		db.Where("link = ?", value.Link).First(&article)
		o[i].Title = article.Title
		o[i].Summary = article.Summary
		o[i].Comment = article.Comment
		o[i].Link = article.Link
		//查找 count
		var count Count
		db.Where("link = ?", value.Link).First(&count)
		o[i].Count = count.Count
		//查找 Tag
		var tags1 []Tag
		db.Where("link = ?", value.Link).Find(&tags1)
		o[i].Tag = TagTOS(tags1)
	}
	// 冒泡排序
	o = BubbleSorting(o)
	return o
}

func BubbleSorting(o []Object) (ob []Object) {
	leng := len(o)
	tmp := Object{}
	for i := 0; i < leng-1; i++ {
		for j := i + 1; j < leng-1; j++ {
			//交换
			if o[i].Count < o[j+1].Count {
				tmp = o[i]
				o[i] = o[j+1]
				o[j+1] = tmp
			}
		}
	}
	return o
}

func GetSummary(link string) (article []Article) {
	db.Where("link = ?",link).Find(&article)
	return article
}
