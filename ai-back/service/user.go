// user.go

package service

import (
	"awesomeProject3/api"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	HasNameOrEmail(name string, email string) error
	RegisterUser(username, email, password string) error
	LoginUser(user *api.User) error
	FindUser(user *api.User) error
	FindAllUser(users *[]api.User) error
	DeleteUser(id string) error
}

type userService struct {
	// 可以在这里注入其他依赖，例如数据库连接、缓存等
}

func (u userService) HasNameOrEmail(name string, email string) error {
	var user api.User

	api.Db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return errors.New("用户名已存在")
	}

	api.Db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return errors.New("邮箱已存在")
	}

	fmt.Println(user)

	return nil
}

func (u userService) RegisterUser(username, email, password string) error {
	err := u.HasNameOrEmail(username, email)
	if err != nil {
		return err
	}

	//密码加密
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	api.Db.Create(&api.User{Name: username, Email: email, Password: string(fromPassword)})
	return nil
}

func (u userService) LoginUser(user *api.User) error {
	fmt.Println("账号" + user.Name)
	fmt.Println("密码" + user.Password)
	password := user.Password
	api.Db.Where("name = ?", user.Name).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("密码错误")
	}
	return nil
}

func (u userService) FindUser(user *api.User) error {
	// 通过id查找用户
	err := api.Db.Model(&api.User{}).Preload("Chats").Find(&user)
	//如果找不到user
	if err.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	return nil
}

func (u userService) FindAllUser(users *[]api.User) error {
	api.Db.Find(&users)
	return nil
}

func (u userService) DeleteUser(id string) error {
	//通过id删除用户，没有用户则返回错误
	var user api.User
	api.Db.Find(&user, id)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	api.Db.Delete(&user)
	return nil
}

func NewUserService() UserService {
	return &userService{}
}
