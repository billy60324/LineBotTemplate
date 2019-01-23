package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getOperationCode(messageToken []string) int {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM OperationList")
	defer rows.Close()
	checkErr(err)

	operationCode := -1
	for tokenIndex := 0; tokenIndex < len(messageToken); tokenIndex++ {
		for rows.Next() {
			var code int
			var keyword string
			err = rows.Scan(&code, &keyword)
			println(messageToken[tokenIndex] + "===" + keyword)
			checkErr(err)
			if messageToken[tokenIndex] == keyword {
				operationCode = code
			}
		}
		print(opCodeDefine[0].opCode)
	}

	return operationCode
}

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

func getOperationCodeT(messageToken []string) string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM OperationList")

	checkErr(err)

	operationCode := "11221"
	for tokenIndex := 0; tokenIndex < len(messageToken); tokenIndex++ {
		for rows.Next() {
			var code int
			var keyword string
			err = rows.Scan(&code, &keyword)
			println(messageToken[tokenIndex] + "===" + keyword)
			checkErr(err)
			operationCode = keyword
		}
		print(opCodeDefine[0].opCode)
		print(rows.Columns)
	}
	defer rows.Close()
	return operationCode
}
