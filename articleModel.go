package blog

import (
	"time"
)

func ArticleInsert(a Article) int64 {
	affected, err := engine.Insert(&a)
	if err != nil {
		return -1
	} else {
		return affected
	}

}
func ArticleUpdate(a Article) int64 {
	affected, err := engine.Where("id=?", a.Id).Update(&a)
	if err != nil {
		return -1
	} else {
		return affected
	}
}
func ArticleDelete(a Article) bool {
	_, err := engine.Where("id=?", a.Id).Delete(&a)
	if err != nil {
		return true
	} else {
		return false
	}
}
func ArticleFindById(id int64) (a Article, has bool) {
	has, _ = engine.Where("id=?", id).Get(&a)

	return a, has
}
func ArticleFindByCate(cateid int64) (as []Article) {
	err := engine.Where("category=?", cateid).Find(&as)
	if err != nil {
		return
	}
	return as
}
func ArticleFindByPage(pagenum int, pagesize int) (as []Article) {
	//limit(pagesize,offset)

	err := engine.Desc("id").Limit(pagesize, pagenum*pagesize).Find(&as)
	if err != nil {
		return
	}
	return as
}
func ArticleFindByUserByPage(u User, pagenum int, pagesize int) (as Article) {
	err := engine.Where("author=?", u.Id).Desc("id").Limit(pagesize, pagenum*pagesize).Find(&as)
	if err != nil {
		return
	}
	return as
}
func ArticleFindByCateAndPage(cateid int64, pagenum int, pagesize int) (as []Article) {
	err := engine.Where("category=?", cateid).Desc("id").Limit(pagesize, pagenum*pagesize).Find(&as)
	if err != nil {
		return
	}
	return as
}
func ArticleFindByCre(start time.Time, end time.Time) (as []Article) {
	engine.Where("created>=? and created <=?", start, end).Find(&as)
	return
}
func ArticleTotalCount(a interface{}) int {
	//等于零时约束视为不存在
	//a.Author = int64(0)
	count, _ := engine.Count(a)
	return int(count)
}
