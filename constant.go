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
	{opCode: 2, keyword: "學習", complete: true},
	{opCode: 3, keyword: "忘記", complete: true},
	{opCode: 4, keyword: "什麼是", complete: true},
	{opCode: 5, keyword: "是什麼", complete: true},
	{opCode: 99, keyword: "DB", complete: true}, // ------------------  testing for DB ------------------
	{opCode: 201, keyword: "484", complete: false},
	{opCode: 202, keyword: "是不是", complete: false},
	{opCode: 203, keyword: "要不要", complete: false},
}

// OpCodeSyntaxDefine : define operation syntax
var OpCodeSyntaxDefine = []OpCodeSyntaxInfo{
	{opCode: 1, location: 0, length: 1},
	{opCode: 2, location: 0, length: 3},
	{opCode: 3, location: 0, length: 2},
	{opCode: 4, location: 0, length: 2},
	{opCode: 5, location: 1, length: 2},
	{opCode: 99, location: 0, length: 2},
	{opCode: 201, location: NoNeed, length: NoNeed},
	{opCode: 202, location: NoNeed, length: NoNeed},
	{opCode: 203, location: NoNeed, length: NoNeed},
}

const (
	// NoNeed : some parameter no need
	NoNeed = -1
	// GetRandomPicture : get random picture from database
	GetRandomPicture = 1
	// TeachKeyword : insert keyword information to LearnTable
	TeachKeyword = 2
	// ForgetKeyword : delete keyword information from LearnTable
	ForgetKeyword = 3
	// WhatIs : get detial information from keyword
	WhatIs = 4
	// IsWhat : get detial information from keyword
	IsWhat = 5
)
