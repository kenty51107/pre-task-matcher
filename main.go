package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
    // dbをインスタンス化
    db := db.NewDB()
    // userValidatorのコンストラクタを呼び出す
    userValidator := validator.NewUserValidator()
    // taskValidatorのコンストラクタを呼び出す
    taskValidator := validator.NewTaskValidator()
    // レポジトリパッケージで定義したコンストラクタを呼び出す
    userRepository := repository.NewUserRepository(db)
    // taskRepositoryのコンストラクタを呼び出す
    taskRepository := repository.NewTaskRepository(db)
    // usecaseパッケージで定義したコンストラクタを呼び出す
    userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
    // taskUsecaseのコンストラクタを呼び出す
    taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
    // controllerパッケージで定義したコンストラクタを呼び出す
    userController := controller.NewUserController(userUseCase)
    // taskControllerのコンストラクタを呼び出す
    taskController := controller.NewTaskController(taskUsecase)
    // routerパッケージで定義したコンストラクタを呼び出す
    e := router.NewRouter(userController, taskController)
    // サーバー起動
    e.Logger.Fatal(e.Start(":8080"))
}
