package database

import (
	"CanYouGetTo20_REST-API/pkg/score"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type ScoreRepository struct {
	db *sql.DB
}

type DbConfig struct {
	server, port, user, password, database string
}

func NewScoreRepository(c DbConfig) *ScoreRepository {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		c.server, c.user, c.password, c.port, c.database)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected!")

	return &ScoreRepository{db}
}

func (repo *ScoreRepository) GetScore() score.Scores {
	var scores score.Scores
	stmt, err := repo.db.Prepare("SELECT TOP(10) name, score FROM scores ORDER BY score DESC")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var scoreEntry score.Score
		err := rows.Scan(&scoreEntry.Name, &scoreEntry.Score)

		if err != nil {
			log.Fatal(err)
		}

		scores = append(scores, scoreEntry)
	}

	return scores
}

func (repo *ScoreRepository) SubmitScore(score score.Score) {
	stmt, err := repo.db.Prepare("INSERT INTO scores (name, score) VALUES(@Name, @Score)")

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		sql.Named("Name", score.Name),
		sql.Named("Score", score.Score),
	)

	if err != nil {
		log.Println(err)
	}
}
