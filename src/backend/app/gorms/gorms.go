package gorms

import (
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorms struct {
	Gorm *gorm.DB
}

func NewGorms() Gorms {
	host := "ec2-52-221-241-44.ap-southeast-1.compute.amazonaws.com"
	port := "5432"
	user := "stima"
	password := "stima"
	dbname := "stima"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	conf, _ := db.DB()
	conf.SetMaxIdleConns(3)
	conf.SetMaxOpenConns(10)
	fmt.Println("db on")

	return Gorms{Gorm: db}
}

var Module = fx.Options(
	fx.Provide(NewGorms),
)
