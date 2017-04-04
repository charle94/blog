package blog

import (
	"fmt"
	"time"
	//"fmt"
	"testing"
)

func Test_api(t *testing.T) {
	//apiClient()
	var a Article
	BeginModel()
	a, has := ArticleFindById(int64(1))
	c, has := CateFindById(int64(1))
	fmt.Println(a, has, c, has)

}
func Test_insert(t *testing.T) {
	BeginModel()
	a := Article{
		Content:  "我的代码，插入",
		Author:   1,
		Title:    "yes this title",
		Category: 1,
	}
	id := ArticleInsert(a)
	fmt.Println(id)
}
func Test_findbycre(t *testing.T) {
	BeginModel()
	//const shortForm = "2006-Jan-02"
	ti, _ := time.Parse("2006-Jan-02", "2017-Mar-17")
	end, _ := time.Parse("2006-Jan-02", "2017-Mar-18")
	a := ArticleFindByCre(ti, end)
	fmt.Println(a, ti, end)
}
func Test_apiclient(t *testing.T) {
	//apiClient()
}
