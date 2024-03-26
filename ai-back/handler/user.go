// handler.go

package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetEmailCode GET  参数: email(api.Email)
func (f *UserHandler) GetEmailCode(c *gin.Context) {
	var err error
	var userEmail string
	userEmail = c.Query("email")

	redisService := service.NewRedisService()
	emailService := service.NewEmailService()

	//生成、redis储存code
	var code string
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	fmt.Println("给"+userEmail+"的验证码为："+code, "   location is :handler/user.go  GetEmailCode")
	err = redisService.Set("registerCode", code)
	var content = fmt.Sprintf(api.EmailTemplate, code)
	err = emailService.SendEmail(userEmail, api.EmailTitle, content)

	if err != nil {
		c.JSON(200, api.M(api.FAIL, "发送失败", err.Error()))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "发送成功", nil))
}

// CheckRegisterCode  参数: email(用户邮箱) code
// 因为register要post入整个user，不太好融入，所以分开了
func (f *UserHandler) CheckRegisterCode(c *gin.Context) {
	//获取get参数，这里是email和code
	email := c.Query("email")
	code := c.Query("code")

	redisService := service.NewRedisService()
	registerCode, _ := redisService.Get("registerCode")
	if code == registerCode {
		redisService.Set(email, "1")
		c.JSON(200, api.M(api.SUCCESS, "验证成功", nil))
	} else {
		c.JSON(200, api.M(api.FAIL, "验证失败", "验证码错误"))
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
	if err != nil || user.Email == "" || user.Password == "" {
		c.JSON(200, api.M(api.FAIL, "邮箱或密码为空", nil))
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
