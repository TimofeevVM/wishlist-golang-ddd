package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	wishlist3 "wishlist/internal/wishlist"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

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

	wishlist3.InitWishlistModule(db, r)

	r.Run(":7001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
