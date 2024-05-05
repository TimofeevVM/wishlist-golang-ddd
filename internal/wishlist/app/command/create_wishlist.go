package wishlist

import wishlist "wishlist/internal/wishlist/domain"

type CreateWishlistCommand struct {
	Title string
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
	err := h.WishlistRepository.Persist(newWishlist)

	if err != nil {
		return nil, err
	}

	return newWishlist, nil
}
