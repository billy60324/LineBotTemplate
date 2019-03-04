package main

type OperationCodeInfo struct {
	opCode  int
	keyword string
	complete bool
}

var OpCodeDefine = []OperationCodeInfo{
	{opCode: 1, keyword: "抽", complete: true},
	{opCode: 2, keyword: "484", complete: false},
	{opCode: 3, keyword: "是不是", complete: false},
}

const (
	GetPhoto = 1;
)

type OperationCode int

var OperationCodeGetPhoto = 1
