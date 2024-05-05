package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	wishlist_query "wishlist/internal/wishlist/app/query"
	wishlist "wishlist/internal/wishlist/domain"
	"wishlist/internal/wishlist/endpoints/ginhttp/response"
)

type GetWishlistItemResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type GetWishlistResponse struct {
	Id    string                     `json:"id"`
	Title string                     `json:"title"`
	Items []*GetWishlistItemResponse `json:"items"`
}

type GetWishlistAction struct {
	getWishlistHandler *wishlist_query.GetWishlistHandler
}

func NewGetWishlistAction(getWishlistAction *wishlist_query.GetWishlistHandler) *GetWishlistAction {
	return &GetWishlistAction{
		getWishlistHandler: getWishlistAction,
	}
}

func (c *GetWishlistAction) Handle(gin *gin.Context) {
	wl, err := c.getWishlistHandler.Handle(
		wishlist_query.GetWishlistQuery{
			Id: gin.Param("id"),
		})

	if err != nil {
		gin.JSON(http.StatusBadRequest, &response.Error{
			Message: err.Error(),
		})

		return
	}

	gin.JSON(http.StatusOK, createGetWishlistResponse(wl))
}

func createGetWishlistResponse(wl *wishlist.Wishlist) *GetWishlistResponse {
	items := make([]*GetWishlistItemResponse, 0)

	for _, item := range wl.Items {
		items = append(items, &GetWishlistItemResponse{
			Id:   item.Id.String(),
			Text: item.Text,
			Done: item.Done,
		})
	}

	return &GetWishlistResponse{
		Id:    wl.Id.String(),
		Title: wl.Title,
		Items: items,
	}
}
