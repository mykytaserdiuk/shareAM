package main

import (
	"context"
	"fmt"
	"github.com/mykytaserdiuk/shaream/pkg/conf"
	"github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/route"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/service"
	"github.com/mykytaserdiuk/shaream/pkg/minio"
	"net/http"
)

func main() {
	var config config
	err := conf.UnmarshalYAML(&config, "./config.yaml")

	minioDB, err := minio.NewMinio("localhost:9000", "minioadmin", "minioadmin", false)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	bucket := "mybucket"

	err = minioDB.SetupBucket(ctx, bucket)
	if err != nil {
		fmt.Println("minio bucket err: ", err)
	}

	db, err := postgres.NewDB(config.DB.URL())
	if err != nil {
		fmt.Println("connecting db err: ", err)
	}

	svc := service.NewServices(db, minioDB)

	router := route.NewRouter(svc)
	err = http.ListenAndServe(config.Port, router)
	if err != nil {
		panic(err)
	}
}
