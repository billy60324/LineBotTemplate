package main

type OperationCodeInfo struct {
	opCode  int
	keyword string
}

var OpCodeDefine = []OperationCodeInfo{
	{opCode: 1, keyword: "抽"},
	{opCode: 2, keyword: "484"},
	{opCode: 3, keyword: "是不是"},
}

const (
	GetPhoto = 1;
)

type OperationCode int

var OperationCodeGetPhoto = 1
