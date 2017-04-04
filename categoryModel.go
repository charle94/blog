package blog

func CateGet() (cat []Category) {
	err := engine.Find(&cat)
	if err != nil {
		return cat
	}
	return
}

func CateFindById(id int64) (cate Category, has bool) {
	has, _ = engine.Where("id=?", id).Get(&cate)
	return cate, has
}
