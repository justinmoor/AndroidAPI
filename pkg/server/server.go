package server

import (
	"CanYouGetTo20_REST-API/pkg/database"
	"CanYouGetTo20_REST-API/pkg/route"
	"CanYouGetTo20_REST-API/pkg/score"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	scoreRepo *database.ScoreRepository
	router    *mux.Router
	routes    route.Routes
}

func NewServer(repo *database.ScoreRepository) *Server {
	var s Server
	s.routes = registerRoutes(&s)

	s.scoreRepo = repo
	s.router = route.NewRouter(s.routes)

	return &s
}

func registerRoutes(s *Server) route.Routes {
	return route.Routes{
		route.Route{
			"SubmitScore",
			"POST",
			"/submitscore",
			s.submitScore,
		},
		route.Route{
			"ShowScore",
			"GET",
			"/showscore",
			s.showScore,
		},
	}
}

func (s *Server) Run(port int) {
	log.Printf("Server running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.router))
}

func (s *Server) submitScore(w http.ResponseWriter, r *http.Request) {
	var submittedScore score.Score

	name := r.FormValue("name")
	score := r.FormValue("score")

	i, _ := strconv.ParseInt(score, 10, 64)

	submittedScore.Name = name
	submittedScore.Score = i

	s.scoreRepo.SubmitScore(submittedScore)
}

func (s *Server) showScore(w http.ResponseWriter, r *http.Request) {
	scores := s.scoreRepo.GetScore()

	data, err := json.Marshal(scores)

	if err != nil {
		SendError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)

	if err != nil {
		SendError(w)
	}
}

func SendError(w http.ResponseWriter) {
	http.Error(w, "Something went wrong...", http.StatusInternalServerError)
}
