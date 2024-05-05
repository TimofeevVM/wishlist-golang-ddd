package wishlist

import (
	"github.com/google/uuid"
)

type WishlistId string

func NewWishlistId() WishlistId {
	return WishlistId(uuid.New().String())
}

func (id WishlistId) String() string {
	return string(id)
}
