package wishlist

func GinInit(container Container) {
	container.GinEngine().GET("/wishlist/:id", container.GinHttpGetWishlistAction().Handle)

	container.GinEngine().POST("/wishlist", container.GinHttpCreateWishlistAction().Handle)

	container.GinEngine().POST("/wishlist/:id/item", container.GinHttpAddItemAction().Handle)

}
