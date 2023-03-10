package controllers

import (
	"api/server/domain"
	"api/server/interfaces/database"
	"api/server/interfaces/token"
	"api/server/usecase"
	"net/http"
	"sort"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type TwitterController struct {
	InteractorUser   usecase.UserInteractor
	InteractorPost   usecase.PostInteractor
	InteractorFollow usecase.FollowInteractor
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
		InteractorFollow: usecase.FollowInteractor{
			FollowRepository: &database.FollowRepository{
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

func (controller *TwitterController) RelatedPost(c echo.Context) (err error) {
	timeLine := domain.Posts{}

	loginUser := controller.GetLoginUser(c)

	follows, err := controller.InteractorFollow.SearchFollowByUserId("user_id = ?", loginUser.Id)
	if err != nil {
		return err
	}
	for _, v := range follows {
		posts, err := controller.InteractorPost.RelatedByUserId("User", "user_id = ?", v.FollowId)
		if err != nil {
			continue
		}
		timeLine = append(timeLine, posts...)
	}

	sort.SliceStable(timeLine, func(i, j int) bool {
		return timeLine[i].CreatedAt > timeLine[j].CreatedAt
	})
	return c.JSON(http.StatusOK, timeLine)
}

func (controller *TwitterController) CreatePost(c echo.Context) (err error) {
	loginUser := controller.GetLoginUser(c)
	p := domain.NewPost()
	p.UserId = loginUser.Id
	c.Bind(&p)
	post, err := controller.InteractorPost.Add(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusCreated, post)
}

func (controller *TwitterController) CreateFollow(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	loginUser := controller.GetLoginUser(c)
	f := domain.Follow{}
	f, _ = controller.InteractorFollow.SearchByFollowIdAndUserId("user_id = ? And follow_id = ?", loginUser.Id, id)
	if f.Id == 0 {
		f = domain.NewFollow()
		f.UserId = loginUser.Id
		f.FollowId = id
	} else {
		f.DeletedAt = gorm.DeletedAt{}
	}
	follow, err := controller.InteractorFollow.Update(f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusCreated, follow)
}

func (controller *TwitterController) DeleteFollow(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	loginUser := controller.GetLoginUser(c)
	follow, _ := controller.InteractorFollow.SearchByFollowIdAndUserId("user_id = ? And follow_id = ?", loginUser.Id, id)
	err = controller.InteractorFollow.DeleteById(follow)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, follow)
}

func (controller *TwitterController) GetLoginUser(c echo.Context) (loginuser domain.LoginUser) {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(jwt.MapClaims)
	id := claims["uid"].(float64)
	userName := claims["name"].(string)
	loginuser = domain.LoginUser{Id: int(id), Name: userName}
	return
}
