package server

func Start() {
	router := SetRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
