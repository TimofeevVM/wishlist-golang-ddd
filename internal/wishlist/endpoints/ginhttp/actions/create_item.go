package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	wishlist "wishlist/internal/wishlist/app/command"
	"wishlist/internal/wishlist/endpoints/ginhttp/response"
)

type AddItemRequest struct {
	WishListId string
	Text       string `json:"text"`
}

type AddItemResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type AddItemAction struct {
	addItemHandler *wishlist.AddItemHandler
}

func NewCreateItemAction(addItemHandler *wishlist.AddItemHandler) *AddItemAction {
	return &AddItemAction{
		addItemHandler: addItemHandler,
	}
}

func (c *AddItemAction) Handle(gin *gin.Context) {
	request := AddItemRequest{}
	err := gin.ShouldBindJSON(&request)
	if err != nil {
		gin.JSON(http.StatusBadRequest, &response.Error{
			Message: "Неккоректный json",
		})

		return
	}

	wl, err := c.addItemHandler.Handle(
		wishlist.AddItemCommand{
			IdWishlist: gin.Param("id"),
			Text:       request.Text,
		})

	if err != nil {
		gin.JSON(http.StatusBadRequest, &response.Error{
			Message: err.Error(),
		})

		return
	}

	gin.JSON(http.StatusOK, &AddItemResponse{
		Id:   wl.Id.String(),
		Text: wl.Text,
		Done: wl.Done,
	})
}
