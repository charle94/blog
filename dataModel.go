package blog

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var engine *xorm.Engine

type User struct {
	Id      int64
	Name    string `xorm:"unique"`
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Token   string
}
type Article struct {
	Id       int64
	Title    string
	Content  string `xorm:"Text"`
	Author   int64
	Category int64
	Created  time.Time `xorm:"created"`
	Update   time.Time `xorm:"updated"`
}
type Category struct {
	Id       int64
	CateName string
}

func BeginModel() {
	var err error
	//此处不能:=赋值否则全局变量相当于一个局部变量
	engine, err = xorm.NewEngine("mysql", DBCONFIG)
	if err != nil {
		fmt.Println("打开数据库错误")
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
func Migratiton() {

	BeginModel()
	err := engine.Sync2(new(Article))
	err = engine.Sync2(new(User))
	err = engine.Sync2(new(Category))
	if err != nil {

	}
}
