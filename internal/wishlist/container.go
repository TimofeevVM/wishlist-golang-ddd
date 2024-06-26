package wishlist

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	wishlist "wishlist/internal/wishlist/app/command"
	wishlist_query "wishlist/internal/wishlist/app/query"
	"wishlist/internal/wishlist/domain"
	"wishlist/internal/wishlist/endpoints/ginhttp/actions"
	"wishlist/internal/wishlist/infra/repository"
)

type Container interface {
	GinEngine() *gin.Engine
	GetPgPool() *pgxpool.Pool

	/* ginhttp */

	GinHttpAddItemAction() *actions.AddItemAction
	GinHttpGetWishlistAction() *actions.GetWishlistAction
	GinHttpCreateWishlistAction() *actions.CreateWishlistAction

	/* App */

	AppAddItemHandler() *wishlist.AddItemHandler
	AppCreateWishlistHandler() *wishlist.CreateWishlistHandler

	AppGetWishlistHandler() *wishlist_query.GetWishlistHandler

	/* Domain */

	DomainWishlistRepository() domain.Repository
}

type ImplContainer struct {
	pgxpool *pgxpool.Pool
	gin     *gin.Engine

	createItemAction     *actions.AddItemAction
	getWishlistAction    *actions.GetWishlistAction
	createWishlistAction *actions.CreateWishlistAction

	addItemHandler        *wishlist.AddItemHandler
	createWishlistHandler *wishlist.CreateWishlistHandler

	getWishlistHandler *wishlist_query.GetWishlistHandler

	wishlistRepository domain.Repository
}

func NewContainer(gin *gin.Engine, pgxpool *pgxpool.Pool) Container {
	container := &ImplContainer{
		gin:     gin,
		pgxpool: pgxpool,
	}

	return container
}

func (c *ImplContainer) GinEngine() *gin.Engine {
	return c.gin
}

func (c *ImplContainer) GetPgPool() *pgxpool.Pool {
	return c.pgxpool
}

func (c *ImplContainer) DomainWishlistRepository() domain.Repository {
	if c.wishlistRepository == nil {
		c.wishlistRepository = repository.NewPqRepository(c.GetPgPool())
	}

	return c.wishlistRepository
}

func (c *ImplContainer) GinHttpAddItemAction() *actions.AddItemAction {
	if c.createItemAction == nil {
		c.createItemAction = actions.NewCreateItemAction(c.AppAddItemHandler())
	}

	return c.createItemAction
}

func (c *ImplContainer) AppAddItemHandler() *wishlist.AddItemHandler {
	if c.addItemHandler == nil {
		c.addItemHandler = wishlist.NewAddItemHandler(c.DomainWishlistRepository())
	}

	return c.addItemHandler
}

func (c *ImplContainer) GinHttpGetWishlistAction() *actions.GetWishlistAction {
	if c.getWishlistAction == nil {
		c.getWishlistAction = actions.NewGetWishlistAction(c.AppGetWishlistHandler())
	}

	return c.getWishlistAction
}

func (c *ImplContainer) AppGetWishlistHandler() *wishlist_query.GetWishlistHandler {
	if c.getWishlistHandler == nil {
		c.getWishlistHandler = wishlist_query.NewGetWishlistHandler(c.DomainWishlistRepository())
	}

	return c.getWishlistHandler
}

func (c *ImplContainer) GinHttpCreateWishlistAction() *actions.CreateWishlistAction {
	if c.createWishlistAction == nil {
		c.createWishlistAction = actions.NewCreateWishlistAction(c.AppCreateWishlistHandler())
	}

	return c.createWishlistAction
}

func (c *ImplContainer) AppCreateWishlistHandler() *wishlist.CreateWishlistHandler {
	if c.createWishlistHandler == nil {
		c.createWishlistHandler = wishlist.NewCreateWishlistHandler(c.DomainWishlistRepository())
	}

	return c.createWishlistHandler
}
