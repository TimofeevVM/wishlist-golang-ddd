package wishlist

type Wishlist struct {
	Id    WishlistId
	Title string
	Items Items
}

func NewWishlist(title string) *Wishlist {
	return &Wishlist{
		Id:    NewWishlistId(),
		Items: make(Items, 0),
		Title: title,
	}
}

func (wl *Wishlist) AddItem(text string) *Item {
	item := NewItem(text)
	wl.Items = append(wl.Items, item)
	return item
}

func (wl *Wishlist) DeleteItem(id ItemId) bool {
	for i, item := range wl.Items {
		if item.Id.equals(id) {
			wl.Items = append(wl.Items[:i], wl.Items[i+1:]...)
			return true
		}
	}

	return false
}

func (wl *Wishlist) markAsDone(id ItemId) bool {
	for _, item := range wl.Items {
		if item.Id.equals(id) {
			item.markDone()
			return true
		}
	}

	return false
}

func (wl *Wishlist) refreshItem(id ItemId) bool {
	for _, item := range wl.Items {
		if item.Id.equals(id) {
			item.refresh()
			return true
		}
	}

	return false
}
