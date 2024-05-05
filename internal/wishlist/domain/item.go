package wishlist

type Item struct {
	Id   ItemId
	Text string
	Done bool
}

type Items []*Item

func NewItem(text string) *Item {
	return &Item{
		Id:   NewItemId(),
		Text: text,
		Done: false,
	}
}

func (item *Item) markDone() {
	item.Done = true
}

func (item *Item) refresh() {
	item.Done = false
}
