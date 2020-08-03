//+build wireinject

package main

import (
	"time"

	"github.com/google/wire"
)

// ProviderSet
var mosterPlayerSet = wire.NewSet(NewMoster, NewPlayer, NewMission)

// 结构构造器
var endingASet = wire.NewSet(mosterPlayerSet, wire.Struct(
	new(EndingA), "Player", "Moster",
))
var endingBset = wire.NewSet(mosterPlayerSet, wire.Struct(
	new(EndingB), "Player", "Moster",
))

func InitMission(p PlayerParam, m MosterParam, t time.Time) (Mission, func(), error) {
	wire.Build(mosterPlayerSet)
	return Mission{}, nil, nil
}

func InitEndingA(p PlayerParam, m MosterParam) (EndingA, func(), error) {
	wire.Build(endingASet)
	return EndingA{}, nil, nil
}

func InitEndingB(p PlayerParam, m MosterParam) (EndingB, func(), error) {
	wire.Build(endingBset)
	return EndingB{}, nil, nil
}

// 绑定值, 值拷贝
var birdMoster = Moster{Name: "bird"}

func InitManMission(p PlayerParam, t time.Time) (*Mission, func(), error) {
	wire.Build(NewPlayer, wire.Value(birdMoster), NewMission2)
	return &Mission{}, nil, nil
}

// // 结构字段构造器, 通过 Mission 构造 Win
// func InitWin() Win {
// 	wire.Build(NewMission, wire.FieldsOf(
// 		new(Mission), "Win",
// 	))
// 	return Win{}
// }

func InitMissionResult(p PlayerParam, m MosterParam, t time.Time) (Win, func(), error) {
	wire.Build(mosterPlayerSet, wire.FieldsOf(
		new(Mission), "Win",
	))
	return Win{}, nil, nil
}
