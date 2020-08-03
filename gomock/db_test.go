package db

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

/*
参数(Eq, Any, Not, Nil)
	Eq(value) 表示与 value 等价的值
	Any() 可以用来表示任意的入参
	Not(value) 用来表示非 value 以外的值
	Nil() 表示 None 值

返回值(Return, DoAndReturn)
	Return 返回确定的值
	Do Mock 方法被调用时，要执行的操作吗，忽略返回值
	DoAndReturn 可以动态地控制返回值
*/
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockDB(ctrl)
	// m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, nil)
	// m.EXPECT().Get(gomock.Any()).Return(630, nil)
	// m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	// m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))
	// m.EXPECT().Get(gomock.Eq("ok")).Do(func(key string) {
	// 	t.Log(key)
	// })

	if v, err := GetFromDB(m, "Tom"); err != nil {
		t.Fatal("expected no error, but got error.", err)
	} else {
		t.Log("get value ", v)
	}
}

/*
调用次数(Times)
	Times() 断言 Mock 方法被调用的次数
	MaxTimes() 最大次数
	MinTimes() 最小次数
	AnyTimes() 任意次数（包括 0 次）
*/
func TestGetFromDBTimes(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")
}

/*
调用顺序(InOrder)
*/

func TestGetFromDBOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}
