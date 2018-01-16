package main

type Score struct{
	Name string `json:"name"`
	Score int64 `json:"score"`
}

type Scores [] Score
