package interfaces

import (
	"cleango/app/usecases"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController(conn *sql.DB) *UserController {
	return &UserController{
		UserInteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{ConnMySQL: conn},
			JSONResponse:   &JSONResponse{},
		},
	}
}

func (uc *UserController) Save(c *gin.Context) {

	uc.UserInteractor.Store(c.Writer, c.Request)

}

func (uc *UserController) List(c *gin.Context) {

	uc.UserInteractor.List(c.Writer, c.Request)

}
