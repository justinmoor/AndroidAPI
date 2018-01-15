package main

type Score struct{
	Name string `json:"naam"`
	Score int64 `json:"score"`
}

type Scores [] Score