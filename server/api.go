package server

func (s *Server) APIRoute() {
	api := s.engine.Group("/api/v1")
	api.POST("/setup", s.ctr.Setup)
}
