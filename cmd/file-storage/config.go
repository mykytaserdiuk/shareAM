package main

import "github.com/mykytaserdiuk/shaream/pkg/db"

type config struct {
	Port string     `yaml:"port"`
	DB   *db.Config `yaml:"db"`
}
