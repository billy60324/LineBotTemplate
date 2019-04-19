package main

// LearnTable : Keyword answer table
type LearnTable struct {
	Keyword   string `form:"id" json:"id" gorm:"id"`
	Response  string `form:"log_name" json:"log_name" gorm:"column:log_name"`
	Teacher   string `form:"create_time" json:"create_time" gorm:"column:create_time"`
	Timestamp string `form:"create_date" json:"create_date" gorm:"column:create_date"`
}

// StockInformation : stock struct
type StockInformation struct {
	/*	z	當盤成交價
		tv	當盤成交量
		v	累積成交量
		b	揭示買價(從高到低，以_分隔資料)
		g	揭示買量(配合b，以_分隔資料)
		a	揭示賣價(從低到高，以_分隔資料)
		f	揭示賣量(配合a，以_分隔資料)
		o	開盤
		h	最高
		l	最低
		y	昨收
		u	漲停價
		w	跌停價
		tlong	epoch毫秒數
		d	最近交易日期(YYYYMMDD)
		t	最近成交時刻(HH:MI:SS)
		c	股票代號
		n	公司簡稱
		nf	公司全名
	*/
	Z string `json:"z"`
	O string `json:"o"`
	H string `json:"h"`
	L string `json:"l"`
	Y string `json:"y"`
	T string `json:"t"`
	C string `json:"c"`
	N string `json:"n"`
}
