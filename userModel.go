package blog

import (
	"crypto/md5"
	"fmt"
	"time"
)

func UserLogin(name string, password string) bool {
	var u User
	has, _ := engine.Where("name=? and passwd=?", name, password).Get(&u)
	if has {
		return true
	} else {
		return false
	}
}
func GetToken(u *User) string {
	if UserLogin(u.Name, u.Passwd) {
		data := []byte(u.Salt + u.Name + fmt.Sprintf("%v", time.Now()))
		token := md5.Sum(data)
		u.Token = fmt.Sprintf("%v", token)
		engine.Where("name=?", u.Name).Update(&u)
		return u.Token
	} else {
		return ""
	}
}
