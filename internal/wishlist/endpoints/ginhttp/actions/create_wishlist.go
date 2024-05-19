package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	wishlist "wishlist/internal/wishlist/app/command"
	"wishlist/internal/wishlist/endpoints/ginhttp/response"
)

type CreateWishlistRequest struct {
	Title string `json:"title"`
}

type CreateWishlistResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type CreateWishlistAction struct {
	addItemHandler *wishlist.CreateWishlistHandler
}

func NewCreateWishlistAction(addItemHandler *wishlist.CreateWishlistHandler) *CreateWishlistAction {
	return &CreateWishlistAction{
		addItemHandler: addItemHandler,
	}
}

func (c *CreateWishlistAction) Handle(ginContext *gin.Context) {
	request := CreateWishlistRequest{}
	err := ginContext.ShouldBindJSON(&request)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, &response.Error{
			Message: "Неккоректный json",
		})

		return
	}

	wl, err := c.addItemHandler.Handle(
		wishlist.CreateWishlistCommand{
			Title:   request.Title,
			Context: ginContext,
		})

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, &response.Error{
			Message: err.Error(),
		})

		return
	}

	ginContext.JSON(http.StatusOK, &CreateWishlistResponse{
		Id:    wl.Id.String(),
		Title: wl.Title,
	})
}
