package wishlist

import (
	"context"
	wishlist "wishlist/internal/wishlist/domain"
)

type CreateWishlistCommand struct {
	Title   string
	Context context.Context
}

type CreateWishlistHandler struct {
	WishlistRepository wishlist.Repository
}

func NewCreateWishlistHandler(repository wishlist.Repository) *CreateWishlistHandler {
	return &CreateWishlistHandler{
		WishlistRepository: repository,
	}
}

func (h *CreateWishlistHandler) Handle(command CreateWishlistCommand) (*wishlist.Wishlist, error) {
	newWishlist := wishlist.NewWishlist(command.Title)
	err := h.WishlistRepository.Persist(command.Context, newWishlist)

	if err != nil {
		return nil, err
	}

	return newWishlist, nil
}
