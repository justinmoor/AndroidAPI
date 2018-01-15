package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var DB *sql.DB

func databaseConnection(){
	db, err := sql.Open("mysql", "root:ipsen123@tcp(127.0.0.1:3306)/android")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	DB = db
}

func getScore() Scores{
	var scores Scores
	stmt, err := DB.Prepare("SELECT * FROM spel ORDER BY score DESC LIMIT 10")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next(){
		var score Score
		err := rows.Scan(&score.Name, &score.Score)

		if err != nil {
			log.Fatal(err)
		}

		scores = append(scores, score)
	}

	return scores
}

func submitScoreToDatabase(score Score){
	stmt, err := DB.Prepare("INSERT INTO spel (naam, score) VALUES(?, ?) ")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	stmt.Exec(score.Name, score.Score)

	if err != nil {
		log.Fatal(err)
	}
}