package wishlist

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type WishlistModule struct {
	db        *sql.DB
	gin       *gin.Engine
	container *Container
}

func InitWishlistModule(db *sql.DB, gin *gin.Engine) WishlistModule {
	container := NewContainer(db, gin)

	module := WishlistModule{
		db:        db,
		gin:       gin,
		container: &container,
	}

	GinInit(container)

	return module
}
