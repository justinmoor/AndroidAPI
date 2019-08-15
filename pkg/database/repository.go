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
	Server, Port, User, Password, Database string
}

func NewScoreRepository(c DbConfig) *ScoreRepository {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		c.Server, c.User, c.Password, c.Port, c.Database)

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

func (repo *ScoreRepository) GetScore() (score.Scores, error) {
	var scores score.Scores
	stmt, err := repo.db.Prepare("SELECT TOP(10) name, score FROM scores ORDER BY score DESC")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var scoreEntry score.Score
		err := rows.Scan(&scoreEntry.Name, &scoreEntry.Score)

		if err != nil {
			return nil, err
		}

		scores = append(scores, scoreEntry)
	}

	return scores, nil
}

func (repo *ScoreRepository) SubmitScore(score score.Score) error {
	stmt, err := repo.db.Prepare("INSERT INTO scores (name, score) VALUES(@Name, @Score)")

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(
		sql.Named("Name", score.Name),
		sql.Named("Score", score.Score),
	); err != nil {
		return err
	}

	return nil
}
