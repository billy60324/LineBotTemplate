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
	{opCode: 6, keyword: "#", complete: true},
	{opCode: 7, keyword: "股票", complete: true},
	{opCode: 8, keyword: "股票", complete: true},
	{opCode: 9, keyword: "股票", complete: true},
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
	{opCode: 6, location: 0, length: 2},
	{opCode: 7, location: 0, length: 1},
	{opCode: 8, location: 0, length: 2},
	{opCode: 9, location: 0, length: 3},
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
	// SearchStock : get stock information
	SearchStock = 6
	// GetFollowStock : get following stock by usetID
	GetFollowStock = 7
	// SetFollowStock : set following stock by userID
	SetFollowStock = 8
	// DeleteFollowStock : delete following stock by userID
	DeleteFollowStock = 9
	// NumberYesNo : random response yes or no
	NumberYesNo = 201
	// YesNo : random response yes or no
	YesNo = 202
	// WantOrNot : random response yes or no
	WantOrNot = 203
)
