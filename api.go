package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Api interface {
	Resopnse(w http.ResponseWriter, r *http.Request)
}

func home(w http.ResponseWriter, r *http.Request) {
	var a Article
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	data := make(map[string]interface{})
	data["hasmore"] = true
	pagenum, _ := strconv.Atoi(r.FormValue("pagenum"))
	pagesize, _ := strconv.Atoi(r.FormValue("pagesize"))
	if pagesize <= 0 {
		pagesize = 6
	}
	as := ArticleFindByPage(pagenum, pagesize)
	if len(as) < pagesize || len(as) == 0 {
		data["hasmore"] = false
	}
	total := ArticleTotalCount(a)
	data["total"] = total
	data["size"] = len(as)
	data["articles"] = as
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "json err")

	}
	fmt.Fprintf(w, string(json))

}
func ApiArticleByCat(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	category := strconv.Atoi(r.FormValue("id"))
	page := strconv.Atoi(r.FormValue("page"))
	fmt.Println(page, category)
	fmt.Fprintf(w, "ok")

}
func ApiCategory(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	switch r.Method {
	case "GET":
		//获取所有分类或者获取单一分类
		if len(r.FormValue("id")) == 0 {
			cats := CateGet()
			data["categorys"] = cats

		} else {
			id, _ := strconv.Atoi(r.FormValue("id"))
			cat, _ := CateFindById(int64(id))
			data["category"] = cat
		}
	case "POST":

	default:

	}
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(w, "json error")
	}
	fmt.Fprintf(w, string(json))
}
func ApiClient() {
	BeginModel()
	http.HandleFunc("/home", home)
	http.HandleFunc("/article", ApiArticleByCat)
	http.HandleFunc("/category", ApiCategory)
	http.ListenAndServe(":8090", nil)
}
