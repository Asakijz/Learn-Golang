package _const

import "testing"

//const定义常量
//iota为自增，初始为0
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstDay(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}
