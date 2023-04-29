package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// controllerのインターフェース
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// usecaseのインターフェースを注入する
// 外部で生成されるusecaseのインスタンスを引数に渡す
// userControllerのインスタンスを生成している
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	// リクエストボディを格納するためのUserオブジェクトを作成する
	user := model.User{}
	// リクエストボディをUserオブジェクトにバインドする
	// そうすると、userオブジェクトにリクエストボディの値が格納される
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// バインドに成功したら、usecaseのSignUpメソッドを呼び出す
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc  *userController) LogIn(c echo.Context) error {
	// リクエストボディを格納するためのUserオブジェクトを作成する
	user := model.User{}
	// リクエストボディをUserオブジェクトにバインドする
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// バインドに成功したら、usecaseのLoginメソッドを呼び出す
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 取得したjwtトークンをサーバーサイドでクッキーに設定する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// バックエンドとフロントエンドのドメインが異なる場合は、SameSite属性をNoneに設定する
	cookie.SameSite = http.SameSiteNoneMode
	// context.SetCookie()で今作成したcookieをHttpレスポンスに設定する
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)

}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// バックエンドとフロントエンドのドメインが異なる場合は、SameSite属性をNoneに設定する
	cookie.SameSite = http.SameSiteNoneMode
	// context.SetCookie()で今作成したcookieをHttpレスポンスに設定する
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"CSRF_token": token,
	})
}
