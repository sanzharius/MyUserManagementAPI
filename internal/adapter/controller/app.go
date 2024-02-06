package controller

type AppController struct {
	UserController interface{ IUserController }
}
