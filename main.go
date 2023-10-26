package main

import (
	"log"

	ordersController "github.com/Gealber/re-partners-challenge/controllers/orders"
	dimensionsRepo "github.com/Gealber/re-partners-challenge/repositories/dimensions"
	ordersService "github.com/Gealber/re-partners-challenge/services/orders"
	badger "github.com/dgraph-io/badger/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	defaultDimensions = []int{5000, 2000, 1000, 500, 250}
)

type Controller struct {
}

func main() {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// service orders
	srv := ordersService.New()
	// dimensions repository
	repo := dimensionsRepo.New(db)
	// initialize default dimensions
	err = repo.Update(defaultDimensions)
	if err != nil {
		log.Fatal(err)
	}

	// controller
	ctr := ordersController.New(repo, srv)

	r := gin.Default()
	// CORS middleware with an insecure way to set it up
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/orders/pack", ctr.GetOrderPacking)
		api.PUT("/orders/dimensions", ctr.PutOrdersDimensions)
		api.GET("/orders/dimensions", ctr.GetOrdersDimensions)
	}

	r.Run()
}

