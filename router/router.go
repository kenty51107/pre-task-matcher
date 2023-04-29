package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// routerのなかでtaskControllerを使えるようにするために
// 引数でtaskControllerを受け取るようにする
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
    // routerの中でuserControllerを使えるようにするために
    // 引数でuserControllerを受け取るようにする

    // Echoのインスタンスを生成
    e := echo.New()

    // corsのミドルウェアを設定する
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        // アクセスを許可するフロントエンドのドメインを設定する
        AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
            echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
        AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
        AllowCredentials: true,
    }))

    // csrfのミドルウェアを設定する
    e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
        CookiePath: "/",
        CookieDomain: os.Getenv("API_DOMAIN"),
        CookieHTTPOnly: true,
        CookieSameSite: http.SameSiteNoneMode,
        // CookieSameSite: http.SameSiteDefaultMode,
        // CookieMaxAge: 60,
    }))
    // echoのインスタンスを渡してルーティングを設定する
    e.POST("/signup", uc.SignUp)
    e.POST("/login", uc.LogIn)
    e.POST("/logout", uc.LogOut)
    e.GET("/csrf", uc.CsrfToken)
    // インスタンスeに対して、グループを作成する
    t := e.Group("/tasks")
    // グループtに対して、jwtのミドルウェアを設定する
    // useキーワードを使うこと、エンドポイントに対してミドルウェアを設定することができる
    // echojwt: jwtのミドルウェア
    t.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(os.Getenv("SECRET")),
        TokenLookup: "cookie:token",
    }))

    // task関係のエンドポイントを設定する
    t.GET("", tc.GetAllTasks)
    t.GET("/:taskId", tc.GetTaskById)
    t.POST("", tc.CreateTask)
    t.PUT("/:taskId", tc.UpdateTask)
    t.DELETE("/:taskId", tc.DeleteTask)
    return e
}
