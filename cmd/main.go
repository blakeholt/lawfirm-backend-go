package main

import (
	"fmt"
	"lawfirm-go-backend/cmd/api"
	"lawfirm-go-backend/config"
	"log"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
