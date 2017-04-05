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

func ApiArticle(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	var (
		method string
		a      Article
	)
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	page, _ := strconv.Atoi(r.FormValue("page"))
	pagesize, _ := strconv.Atoi(r.FormValue("pagesize"))
	if pagesize <= 0 {
		pagesize = PAGEGSIZE
	}
	article := r.FormValue("article")
	json.Unmarshal([]byte(article), &a)
	switch r.Method {
	case "POST":
		//更新文章

	case "GET":
		//根据条件查询文章
		as := ArticleFind(a, page, pagesize)
		total := ArticleTotalCount(a)
		data["total"] = total
		data["size"] = len(as)
		data["articles"] = as
	case "PUT":
		//新增文章
	case "DELETE":

	default:
		//get

	}
	out, _ := json.Marshal(data)
	fmt.Fprintf(w, method)
	fmt.Fprintf(w, string(out))
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
		//改分类
	case "PUT":
		//新增
	case "DELETE":
		//删除
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
	http.HandleFunc("/article", ApiArticle)
	http.HandleFunc("/category", ApiCategory)
	http.ListenAndServe(":8090", nil)
}
