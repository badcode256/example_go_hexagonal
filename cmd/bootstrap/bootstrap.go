package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/badcode256/example_go_hexagonal/internal/infra/database/sql/user"
	"github.com/badcode256/example_go_hexagonal/internal/infra/server"
	"github.com/badcode256/example_go_hexagonal/internal/service"
	_ "github.com/denisenkom/go-mssqldb"
)

func Start() error {
	mySqlURI := fmt.Sprintf("server=localhost;user id=sa;password=123-abc;port=1433;database=Business;", os.Getenv("SQL_USER"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_SERVER"))
	db, err := sql.Open("sqlserver", mySqlURI)

	if err != nil {
		return err
	}

	userRepository := user.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	server := server.New(context.Background(), "localhost", 3000, userService)

	return server.Run()
}
