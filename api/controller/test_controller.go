package controller

import (
	"net/http"

	"github.com/sanjayshr/login/api/responses"
)

func (server *Server) CreatTest(w http.ResponseWriter, r *http.Request) {
	testCreated := "Test Created"

	responses.JSON(w, http.StatusCreated, testCreated)

}
