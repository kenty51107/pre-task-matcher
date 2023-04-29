package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// userRepositoryのインターフェースを埋め込んでいる
// controllerに依存させるためのインターフェース
type IUserUsecase interface {
	// userを値として受け取り、model.UserResponseとerrorを返す
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// usecaseのソースコードはrepositoryのインターフェースにだけ依存している
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// useCaseにrepositoryのインターフェースを注入する
// コンストラクタ
// 引数には外部で生成したrepositoryのインスタンスを渡す
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	// ここでusecaseのインスタンスを生成している
	// この時点でusecaseはrepositoryの実装に依存している
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// 入力値のバリデーションを行う
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// 新しくuserオブジェクトを作成する
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// repositoryのCreateUserメソッドを呼び出す
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// レスポンスとして返すUserResponseオブジェクトを作成する
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	// 入力値のバリデーションを行う
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	// Emailを元に取得したユーザーを格納するための空のUserオブジェクトを作成する
	storedUser := model.User{}
	// Emailを元にユーザーを取得する
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードが一致しているかを確認する
	// 第1引数にはdbに保存されているハッシュ化されたパスワードを、第2引数にはユーザーが入力したパスワードを渡す
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// jwtを生成する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
