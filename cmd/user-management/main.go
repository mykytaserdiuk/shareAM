package main

import (
	"fmt"
	"github.com/mykytaserdiuk/shaream/pkg/conf"
	pg "github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/repository/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/route"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/service"
	"net/http"
)

func main() {
	var config config
	err := conf.UnmarshalYAML(&config, `/configs/config.yaml`)
	if err != nil {
		panic(err)
	}

	db, err := pg.NewDB(config.DB.URL())
	if err != nil {
		fmt.Println("connecting db err: ", err)
		return
	}

	repoManager := postgres.NewRepoManager()

	svc := service.NewServices(db, repoManager)

	router := route.NewRouter(svc)
	err = http.ListenAndServe(config.Port, router)
	if err != nil {
		panic(err)
	}
}
