package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

/*
type User struct {
	code    int
	keyword string
}
*/

func connectDBQuery(queryString string) *sql.Rows {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	checkErr(err)
	log.Print("DB connected, query:" + queryString)
	rows, err := db.Query(queryString)
	checkErr(err)
	log.Print("get query Successed")
	return rows
}

func dbSearchKeywordDetail(keyword string) (string, string, string) {
	rows := connectDBQuery("SELECT response, teacher, timestamp FROM learn WHERE keyword='" + keyword + "'")
	defer rows.Close()
	learnTable := LearnTable{}
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&learnTable.Response, &learnTable.Teacher, &learnTable.Timestamp)
			checkErr(err)
		}
	}
	return learnTable.Response, learnTable.Teacher, learnTable.Timestamp
}

func dbSearchLearnTable(messageToken []string) string {
	response := ""

	rows := connectDBQuery("SELECT keyword, response FROM learn")
	defer rows.Close()
	learnTable := LearnTable{}
	for rows.Next() {
		err := rows.Scan(&learnTable.Keyword, &learnTable.Response)
		checkErr(err)

		for tokenIndex := 0; tokenIndex < len(messageToken); tokenIndex++ {
			if messageToken[tokenIndex] == learnTable.Keyword {
				response = learnTable.Response
				goto Response
			}
		}
	}
Response:
	return response
}

func dbInsertLearnTable(keyword string, response string, teacher string, timestamp string) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	checkErr(err)
	result, err := db.Exec("INSERT INTO learn VALUES ($1, $2, $3, $4)", keyword, response, teacher, timestamp)
	log.Print(result)
	checkErr(err)
}

func dbDeleteLearnTable(keyword string) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	checkErr(err)
	result, err := db.Exec("DELETE FROM learn WHERE keyword=$1", keyword)
	log.Print(result)
	checkErr(err)
}

func dbUserExist(tableName string, userid string) bool {
	query := "SELECT * FROM $1 WHERE userid='$2'"
	query = strings.Replace(query, "$1", tableName, 1)
	query = strings.Replace(query, "$2", userid, 1)
	rows := connectDBQuery(query)

	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			return true
		}
	}

	return false
}

func dbInsertUserStockTable(userid string, stockNumber string) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	checkErr(err)
	result, err := db.Exec("INSERT INTO userstock VALUES ($1, $2)", userid, stockNumber+",")
	log.Print(result)
	checkErr(err)
}

func dbUpdateUserStockTable(userid string, stockNumber string) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	checkErr(err)
	result, err := db.Exec("UPDATE userstock SET followstock = $1 WHERE userid = $2", stockNumber, userid)
	log.Print(result)
	checkErr(err)
}

func dbGetFollowStock(userid string) string {
	query := "SELECT followstock FROM userstock WHERE userid='$1'"
	query = strings.Replace(query, "$1", userid, 1)
	rows := connectDBQuery(query)
	followstock := ""

	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&followstock)
			checkErr(err)
		}
	}
	return followstock
}

func dbStockExist(userid string, stockNumber string) bool {
	query := "SELECT followstock FROM userstock WHERE userid='$1'"
	query = strings.Replace(query, "$1", userid, 1)
	rows := connectDBQuery(query)
	followStock := ""

	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&followStock)
			checkErr(err)
		}
	}
	strings.Replace(followStock, ",", " ", -1)
	stockArray := strings.Split(followStock, " ")
	for i := 0; i < len(stockArray); i++ {
		if stockArray[i] == stockNumber {
			return true
		}
	}

	return false
}

/*
func dbtesting(command string) string {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)
	log.Print("connected")

	rows, err := db.Query(
		"SELECT code, keyword FROM operationlist",
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

		_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id       SERIAL,
			username VARCHAR(64) NOT NULL UNIQUE,
			CHECK (CHAR_LENGTH(TRIM(username)) > 0)
		);
		`)

	db.Close()
	if err != nil {
		log.Print(err)
	}
	return total
}
*/

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
