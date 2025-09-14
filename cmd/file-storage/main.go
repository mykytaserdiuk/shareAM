package main

import (
	"context"
	"fmt"
	"github.com/mykytaserdiuk/shaream/pkg/db"
	"github.com/mykytaserdiuk/shaream/pkg/inventory/route"
	"github.com/mykytaserdiuk/shaream/pkg/inventory/service"
	"net/http"
)

func main() {
	minioDB, err := db.NewMinio("localhost:9000", "minioadmin", "minioadmin", false)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	bucket := "mybucket"

	err = minioDB.ConfigureBucket(ctx, bucket)
	if err != nil {
		fmt.Println("minio bucket err: ", err)
	}

	svc := service.NewServices(minioDB)

	router := route.NewRouter(svc)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
