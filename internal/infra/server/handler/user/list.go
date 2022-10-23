package user

import (
	"net/http"

	domain "github.com/badcode256/example_go_hexagonal/internal/domain"
	"github.com/badcode256/example_go_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
)

func ListHandler(userService service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		usersFound, err := userService.ListUsers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: "not_found"})
			return
		}

		ctx.JSON(http.StatusOK, usersFound)
	}
}
