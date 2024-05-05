package wishlist_query

import wishlist "wishlist/internal/wishlist/domain"

type GetWishlistQuery struct {
	Id string
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
	foundWishlist, err := h.WishlistRepository.GetById(&id)

	if err != nil {
		return nil, err
	}

	return foundWishlist, nil
}
