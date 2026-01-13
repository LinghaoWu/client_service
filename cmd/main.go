package main

import (
	"client_service/api"
	"client_service/configs"
	"client_service/pkg/snowflake"
	"log"
)

func main() {
	configs.InitConfig()

	if err := snowflake.InitSnowflake(configs.Config.Node.NodeID); err != nil {
		log.Fatalf("Failed to initialize ID generator: %v", err)
	}

	r := api.InitRouter()
	r.Run(":" + configs.Config.Service.Port)
}
