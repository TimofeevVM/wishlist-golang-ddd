package wishlist

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WishlistModule struct {
	db        *sql.DB
	pgxPool   *pgxpool.Pool
	gin       *gin.Engine
	container *Container
}

func InitWishlistModule(pgxPool *pgxpool.Pool, db *sql.DB, gin *gin.Engine) WishlistModule {
	container := NewContainer(db, gin, pgxPool)

	module := WishlistModule{
		db:        db,
		pgxPool:   pgxPool,
		gin:       gin,
		container: &container,
	}

	GinInit(container)

	return module
}
