package user

import (
	"net/http"

	domain "github.com/badcode256/example_go_hexagonal/internal/domain"
	"github.com/badcode256/example_go_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(userService service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req domain.UUser
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, &domain.Response{Message: "bad_request"})
			return
		}
		err := userService.UpdateUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
	}
}
