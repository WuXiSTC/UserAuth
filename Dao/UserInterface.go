package Dao

//最基础的用户接口
type User interface {
	GetID() string   //获取ID
	SetID(string)    //设置ID
	GetPASS() string //获取密码
	SetPASS(string)  //设置密码
}
