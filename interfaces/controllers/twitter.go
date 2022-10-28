package controllers

import (
	"api/server/config"
	"api/server/domain"
	"api/server/interfaces/database"
	"api/server/interfaces/token"
	"api/server/usecase"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type TwitterController struct {
	InteractorUser usecase.UserInteractor
	InteractorPost usecase.PostInteractor
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
		InteractorPost: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SqlHandler: SqlHandler,
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

func (controller *TwitterController) CreatePost(c echo.Context) (err error) {
	loginUser := controller.GetLoginUser(c)
	p := domain.Post{
		UserId:    loginUser.Id,
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
	c.Bind(&p)
	post, err := controller.InteractorPost.Add(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusCreated, post)
}

func (controller *TwitterController) GetLoginUser(c echo.Context) (loginuser domain.LoginUser) {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(jwt.MapClaims)
	id := claims["uid"].(float64)
	userName := claims["name"].(string)
	loginuser = domain.LoginUser{Id: int(id), Name: userName}
	return
}
