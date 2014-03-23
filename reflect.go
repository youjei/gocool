package main

import (
	"fmt"
	"reflect" // 这里引入reflect模块
)

type User struct {
	Name   string "user name" //这引号里面的就是tag
	Passwd string "user passsword"
}

func main() {
	user := &User{{"chronos", "pass"}, {"chronos2", "pass2"}}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义

	//a, _ := s.FieldByName("Name") //将tag输出出来
	fmt.Println(s)

}
