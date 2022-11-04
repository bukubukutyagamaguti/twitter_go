package controllers

import (
	"api/server/domain"
	"api/server/interfaces/database"
	"api/server/interfaces/token"
	"api/server/usecase"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(
	SqlHandler database.SqlHandler,
	tokenHandler token.TokenHandler,
) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: SqlHandler,
			},
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

func (controller *UserController) Show(c echo.Context) (err error) {
	loginuser := controller.GetLoginUser(c)
	user, err := controller.Interactor.UserById(loginuser.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := domain.NewUser()
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusCreated, user)
}

func (controller *UserController) Save(c echo.Context) (err error) {
	loginuser := controller.GetLoginUser(c)
	u, err := controller.Interactor.UserById(loginuser.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		Id: id,
	}
	err = controller.Interactor.DeleteById(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetLoginUser(c echo.Context) (loginuser domain.LoginUser) {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(jwt.MapClaims)
	id := claims["uid"].(float64)
	userName := claims["name"].(string)
	loginuser = domain.LoginUser{Id: int(id), Name: userName}
	return
}
