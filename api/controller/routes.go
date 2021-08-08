package controller

import "github.com/sanjayshr/login/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/test", middlewares.SetMiddleWareJSON(s.CreatTest)).Methods("POST")

}
