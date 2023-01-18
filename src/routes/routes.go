package routes

func (s *Server) initRoute() {
	s.Router.POST("/login", s.Controller.Login())
}

func (s *Server) StartServer() {
	s.initRoute()
	s.Router.Run("localhost:5000")
}
