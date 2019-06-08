package models

type Article struct {
	ID      int    `json:"-"`
	Summary string `json:"summary"`
	Comment string `json:"comment"`
	Link    string `json:link`
	//Tag     string `json:`
}

// 保存分享的内容
func Send_share(summary, comment, link string) {

}
