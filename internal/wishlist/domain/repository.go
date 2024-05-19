package domain

import "context"

type Repository interface {
	Persist(ctx context.Context, wishlist *Wishlist) error
	GetById(ctx context.Context, id *WishlistId) (*Wishlist, error)
}
