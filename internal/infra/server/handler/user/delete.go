package user

import (
	"net/http"
	"strconv"

	domain "github.com/badcode256/example_go_hexagonal/internal/domain"
	"github.com/badcode256/example_go_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(userService service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, errP := strconv.ParseInt(ctx.Param("id"), 0, 64)

		if errP != nil {
			ctx.JSON(http.StatusNotFound, &domain.Response{Message: "invalid id"})
			return
		}
		err := userService.DeleteUser(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
