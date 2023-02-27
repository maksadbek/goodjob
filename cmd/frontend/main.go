package main

import (
	"fmt"
	"net/http"

	"goodjob/cmd/frontend/config"
)

func main() {
	config.Init()

	fmt.Println(config.Config)
}

func errorHandler(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	
}
