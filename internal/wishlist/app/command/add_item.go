package wishlist

import wishlist "wishlist/internal/wishlist/domain"

type AddItemCommand struct {
	IdWishlist string
	Text       string
}

type AddItemHandler struct {
	WishlistRepository wishlist.Repository
}

func NewAddItemHandler(repository wishlist.Repository) *AddItemHandler {
	return &AddItemHandler{
		WishlistRepository: repository,
	}
}

func (h *AddItemHandler) Handle(command AddItemCommand) (*wishlist.Item, error) {
	id := wishlist.WishlistId(command.IdWishlist)
	foundWishlist, err := h.WishlistRepository.GetById(&id)

	if err != nil {
		return nil, err
	}

	item := foundWishlist.AddItem(command.Text)

	err = h.WishlistRepository.Persist(foundWishlist)

	if err != nil {
		return nil, err
	}

	return item, nil
}
