package main

import (
	"discounts/api"
	"discounts/bootstrap"
	"discounts/domain/service"
	"discounts/store"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatal("config:", err)
	}

	db, err := bootstrap.ConnectDB(c)
	if err != nil {
		log.Fatal("connect Db:", err)
	}

	defer db.Close()

	s := store.NewStore(db)

	ds := service.NewDiscountService(s)

	srv := api.NewServer(ds)
	err = srv.Start(c.HttpPort)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
