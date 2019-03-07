package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

type User struct {
	code    int
	keyword string
}

func dbtesting(command string) string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)
	log.Print("connected")

	rows, err := db.Query(
		"SELECT code, keyword FROM public.operationlist",
	)

	checkErr(err)
	log.Print("already get query")

	defer rows.Close()

	var total string = ""
	user := User{}
	for rows.Next() {
		log.Print("In for loop")
		err = rows.Scan(&user.code, &user.keyword)
		checkErr(err)
		log.Print("scan success")

		total += strconv.Itoa(user.code) + ":" + user.keyword + "\n"
	}

	//_, err = db.Exec(command)
	/*
		_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id       SERIAL,
			username VARCHAR(64) NOT NULL UNIQUE,
			CHECK (CHAR_LENGTH(TRIM(username)) > 0)
		);
		`)
	*/
	db.Close()
	if err != nil {
		log.Print(err)
	}
	return total
}

/*
func getOperationCode(messageToken []string) int {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM OperationList")

	checkErr(err)


	operationCode := opCodeDefine[0].opCode
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
*/
