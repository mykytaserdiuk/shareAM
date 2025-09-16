package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mykytaserdiuk/shaream/pkg/conf"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/repository/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/route"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/service"
	"github.com/mykytaserdiuk/shaream/pkg/minio"

	pg "github.com/mykytaserdiuk/shaream/pkg/db/postgres"
)

func main() {
	var config config
	err := conf.UnmarshalYAML(&config, `/configs/config.yaml`)
	if err != nil {
		panic(err)
	}

	minioDB, err := minio.NewMinio("minio:9000", "minioadmin", "minioadmin", false)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	bucket := "mybucket"

	err = minioDB.SetupBucket(ctx, bucket)
	if err != nil {
		fmt.Println("minio bucket err: ", err)
		return
	}

	db, err := pg.NewDB(config.DB.URL())
	if err != nil {
		fmt.Println("connecting db err: ", err)
		return
	}

	repoManager := postgres.NewRepoManager()

	svc := service.NewServices(db, repoManager, minioDB)

	router := route.NewRouter(svc)
	err = http.ListenAndServe(config.Port, router)
	if err != nil {
		panic(err)
	}
}
