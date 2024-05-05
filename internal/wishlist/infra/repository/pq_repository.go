package wishlist

import (
	"database/sql"
	"encoding/json"
	wishlist "wishlist/internal/wishlist/domain"
)

type PqRepository struct {
	db *sql.DB
}

func (r *PqRepository) GetById(id *wishlist.WishlistId) (*wishlist.Wishlist, error) {
	row := r.db.QueryRow(`select id, data from wishlist where id = $1`, id.String())

	var idRaw string
	var data string
	err := row.Scan(&idRaw, &data)

	if err != nil {
		return nil, err
	}

	wishlistData := WishlistData{}
	err = json.Unmarshal(
		[]byte(data),
		&wishlistData,
	)

	if err != nil {
		return nil, err
	}

	items := make(wishlist.Items, 0)

	for _, item := range wishlistData.Items {
		items = append(items, &wishlist.Item{
			Id:   wishlist.ItemId(item.Id),
			Text: item.Text,
			Done: item.Done,
		})
	}

	return &wishlist.Wishlist{
		Id:    wishlist.WishlistId(idRaw),
		Title: wishlistData.Title,
		Items: items,
	}, nil

}

func NewPqRepository(db *sql.DB) *PqRepository {
	return &PqRepository{
		db: db,
	}
}

type WishlistData struct {
	Id    string             `json:"id"`
	Title string             `json:"title"`
	Items []WishlistItemData `json:"items"`
}

type WishlistItemData struct {
	Id   string `json:"id"`
	Text string `json:"title"`
	Done bool   `json:"done"`
}

func (r *PqRepository) Persist(wishlist *wishlist.Wishlist) error {
	items := make([]WishlistItemData, 0)

	for _, item := range wishlist.Items {
		items = append(items, WishlistItemData{
			Id:   item.Id.String(),
			Text: item.Text,
			Done: item.Done,
		})
	}

	wishlistData := WishlistData{
		Id:    wishlist.Id.String(),
		Title: wishlist.Title,
		Items: items,
	}

	// Сериализация структуры в строку
	jsonBytes, err := json.Marshal(&wishlistData)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		`INSERT INTO wishlist (id, data) 
VALUES ($1, $2)
ON CONFLICT(id) DO UPDATE SET data = EXCLUDED.data`,
		wishlist.Id.String(),
		string(jsonBytes),
	)

	if err != nil {
		return err
	}

	return nil
}
