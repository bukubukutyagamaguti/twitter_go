package controllers

import (
	"api/server/domain"
	"api/server/interfaces/database"
	"api/server/interfaces/token"
	"api/server/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type TwitterController struct {
	InteractorUser usecase.UserInteractor
}

func NewTwitterController(
	SqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *TwitterController {
	return &TwitterController{
		InteractorUser: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: SqlHandler,
			},
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

func (controller *TwitterController) Login(c echo.Context) (err error) {
	loginuser := domain.LoginUser{}
	if err := c.Bind(&loginuser); err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	user, token, err := controller.InteractorUser.Login(loginuser)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": user, "token": token})
}

func (controller *TwitterController) GetLoginUser(c echo.Context) (loginuser domain.LoginUser) {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(jwt.MapClaims)
	id := claims["uid"].(float64)
	userName := claims["name"].(string)
	loginuser = domain.LoginUser{Id: int(id), Name: userName}
	return
}
