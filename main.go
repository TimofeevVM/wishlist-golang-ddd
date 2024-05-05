package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	wishlist "wishlist/internal/wishlist/app/command"
	wishlist_query "wishlist/internal/wishlist/app/query"
	wishlist2 "wishlist/internal/wishlist/infra/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type CreateWishlistRequest struct {
	Title string `json:"title"`
}

type AddItemRequest struct {
	Text string `json:"text"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/wishlist", func(c *gin.Context) {
		handler := wishlist.NewCreateWishlistHandler(
			wishlist2.NewPqRepository(db),
		)

		request := CreateWishlistRequest{}
		err := c.ShouldBindJSON(&request)
		if err != nil {
			log.Fatalln("unmarshal ", err.Error())
		}

		wl, err := handler.Handle(
			wishlist.CreateWishlistCommand{
				Title: request.Title,
			})

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": wl.Id,
		})
	})

	r.GET("/wishlist/:id", func(c *gin.Context) {
		handler := wishlist_query.NewGetWishlistHandler(
			wishlist2.NewPqRepository(db),
		)

		wl, err := handler.Handle(
			wishlist_query.GetWishlistQuery{
				Id: c.Param("id"),
			})

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    wl.Id,
			"title": wl.Title,
		})
	})

	r.POST("/wishlist/:id/item", func(c *gin.Context) {
		handler := wishlist.NewAddItemHandler(
			wishlist2.NewPqRepository(db),
		)

		request := AddItemRequest{}
		err := c.ShouldBindJSON(&request)
		if err != nil {
			log.Fatalln("unmarshal ", err.Error())
		}

		wl, err := handler.Handle(
			wishlist.AddItemCommand{
				IdWishlist: c.Param("id"),
				Text:       request.Text,
			})

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    wl.Id,
			"title": wl.Text,
			"done":  wl.Done,
		})
	})

	r.Run(":7001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
