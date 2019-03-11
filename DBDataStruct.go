package main

// LearnTable : Keyword answer table
type LearnTable struct {
	Keyword   string `form:"id" json:"id" gorm:"id"`
	Response  string `form:"log_name" json:"log_name" gorm:"column:log_name"`
	Teacher   string `form:"create_time" json:"create_time" gorm:"column:create_time"`
	Timestamp string `form:"create_date" json:"create_date" gorm:"column:create_date"`
}
