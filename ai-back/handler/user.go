// handler.go

package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RegisterUser 注册接口
func (f *UserHandler) RegisterUser(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.userService.RegisterUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "注册成功", nil))
}

// LoginUser 登录接口
func (f *UserHandler) LoginUser(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil || user.Name == "" || user.Password == "" {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.userService.LoginUser(&user)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "登录成功", user))
}

// FindUser 查找用户
func (f *UserHandler) FindUser(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.userService.FindUser(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "查找成功", user))
}

// FindAllUser GET 查找所有用户
func (f *UserHandler) FindAllUser(c *gin.Context) {
	var users []api.User
	err := f.userService.FindAllUser(&users)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "查找成功", users))
}

// DeleteUser GET 删除用户
func (f *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}
	fmt.Println("id is :"+id, "location is :handler/user.go")
	err := f.userService.DeleteUser(id)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "删除成功", nil))
}
