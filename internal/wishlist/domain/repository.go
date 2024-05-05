package wishlist

type Repository interface {
	Persist(wishlist *Wishlist) error
	GetById(id *WishlistId) (*Wishlist, error)
}
