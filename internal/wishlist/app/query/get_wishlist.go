package wishlist_query

import (
	"context"
	wishlist "wishlist/internal/wishlist/domain"
)

type GetWishlistQuery struct {
	Id      string
	Context context.Context
}

type GetWishlistHandler struct {
	WishlistRepository wishlist.Repository
}

func NewGetWishlistHandler(repository wishlist.Repository) *GetWishlistHandler {
	return &GetWishlistHandler{
		WishlistRepository: repository,
	}
}

func (h *GetWishlistHandler) Handle(command GetWishlistQuery) (*wishlist.Wishlist, error) {
	id := wishlist.WishlistId(command.Id)
	foundWishlist, err := h.WishlistRepository.GetById(command.Context, &id)

	if err != nil {
		return nil, err
	}

	return foundWishlist, nil
}
