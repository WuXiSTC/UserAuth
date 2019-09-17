package Service

import (
	"../Dao"
)

func VerifyCachedUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	return Dao.VerifyCachedUser(user)
}
