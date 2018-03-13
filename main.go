// profilepics generates nice looking profile pictures for a user based on his/her name
package main

import (
	"net/http"
	"time"

	"github.com/emman27/profilepics/imaging"
	"github.com/sirupsen/logrus"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	logrus.Error(err)
}

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "OPTIONS" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	if r.Method == "OPTIONS" {
		return
	}

	name := r.FormValue("name")
	logrus.Info(name)
	w.Header().Set("Content-Type", "image/svg+xml")
	imaging.Generate(name, w)
}
