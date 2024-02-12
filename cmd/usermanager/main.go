package main

import (
	"github.com/labstack/echo/v4"
	"myAPIProject/internal/apperrors"
	"myAPIProject/internal/config"
	"myAPIProject/internal/infrastructure/datastore"
	"myAPIProject/internal/infrastructure/logger"
	"myAPIProject/internal/infrastructure/router"
	"myAPIProject/internal/registry"
)

const Env = "C:\\Users\\sanzh\\go\\src\\myAPIProject\\configs\\.env"

func main() {
	log := logger.NewLogger()
	cfg, err := config.NewConfig(Env)
	if err != nil {
		log.Fatal(err)
	}

	db, err := datastore.SetUpDatabase(cfg, log)
	if err != nil {
		log.Fatal(err)
	}

	reg := registry.NewRegistry(db)
	e := echo.New()
	e = router.NewRouter(e, reg.NewAppController())

	log.Println("app starting")

	err = e.Start(cfg.Port)
	if err != nil {
		log.Fatal(apperrors.ServerStartErr.AppendMessage(err))
	}
}
