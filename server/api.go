package server

func (s *Server) APIRoute() {
	api := s.engine.Group("/api/v1")
	api.POST("/setup", s.ctr.Setup)
	api.POST("/sensor/data", s.ctr.InputData)
	api.GET("/excel", s.ctr.ExportToexcel)
	api.GET("/sensors", s.ctr.GetSensors)
	api.GET("/monitoring/:id_sensor", s.ctr.Monitoring)
}
