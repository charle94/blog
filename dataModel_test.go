package blog

import (
	"fmt"
	//"fmt"

	"testing"
)

func Test_user(*testing.T) {
	Migratiton()
}
func Test_article(t *testing.T) {
	BeginModel()
	a := Article{
		Title:    "欢迎3",
		Category: 2,
		Author:   1,
		Content:  "欢迎内容3",
	}
	fmt.Println(a)
	engine.Insert(&a)
}
func Test_articleget(t *testing.T) {
	BeginModel()
	as := ArticleFind(Article{Author: 0}, 3, 4)
	if true {
		fmt.Println(as)
	}
}
