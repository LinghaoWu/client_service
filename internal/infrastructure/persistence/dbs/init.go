package dbs

import (
	"client_service/configs"
	"client_service/internal/infrastructure/ent"

	_ "github.com/lib/pq"
)

var (
	DB_client *ent.Client
)

func InitDB() {
	dsn := "host" + configs.Config.Database.Host +
		" port=" + configs.Config.Database.Port +
		" user=" + configs.Config.Database.User +
		" dbname=" + configs.Config.Database.Db_name +
		" password=" + configs.Config.Database.Password +
		" sslmode=" + configs.Config.Database.SSLMode

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	DB_client = client
}
