package main

type OperationCodeInfo struct {
	opCode  int
	keyword string
}

var opCodeDefine = []OperationCodeInfo{
	{opCode: 1, keyword: "照片"},
	{opCode: 2, keyword: "111"},
	{opCode: 3, keyword: "111"},
}
/*
var (
	GetPhoto := &OperationCodeInfo{ opCode : 1, keyword : "抽" } ;
)*/

type OperationCode int

var OperationCodeGetPhoto = 1
