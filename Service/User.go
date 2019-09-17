package Service

import (
	"crypto/md5"
	"fmt"
)

type User struct {
	ID   string
	PASS string
}

func (u *User) SetID(ID string) {
	u.ID = ID
}

func (u *User) SetPASS(PASS string) {
	u.PASS = fmt.Sprintf("%x", md5.Sum([]byte(PASS))) //MD5加密
}
func (u User) GetID() string {
	return u.ID
}

func (u User) GetPASS() string {
	return u.PASS
}
