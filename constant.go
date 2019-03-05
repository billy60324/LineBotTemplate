package main

// OperationCodeInfo : struct of operation code infomation
type OperationCodeInfo struct {
	opCode   int
	keyword  string
	complete bool
}

// OpCodeSyntaxInfo : struct of operation syntax
type OpCodeSyntaxInfo struct {
	opCode   int
	location int
	length   int
}

// OpCodeDefine : define operation code and value
var OpCodeDefine = []OperationCodeInfo{
	{opCode: 1, keyword: "抽", complete: true},
	{opCode: 99, keyword: "DB", complete: true}, // ------------------  testing for DB ------------------
	{opCode: 201, keyword: "484", complete: false},
	{opCode: 202, keyword: "是不是", complete: false},
	{opCode: 203, keyword: "要不要", complete: false},
}

// OpCodeSyntaxDefine : define operation syntax
var OpCodeSyntaxDefine = []OpCodeSyntaxInfo{
	{opCode: 1, location: 0, length: 1},
	{opCode: 99, location: 0, length: 2},
	{opCode: 201, location: NoNeed, length: NoNeed},
	{opCode: 202, location: NoNeed, length: NoNeed},
	{opCode: 203, location: NoNeed, length: NoNeed},
}

const (
	// NoNeed : some parameter no need
	NoNeed = -1
)
