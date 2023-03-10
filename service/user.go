package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/model"
	"ltt-gc/serializer"
	"ltt-gc/utils"
	"strconv"
)

type UserService struct {
	ID       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// GetUserByEmail
// @Tags user-service
// @Router /user/:email [get]
func (service *UserService) GetUserByEmail(ctx context.Context, email string) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByEmail(email)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// GetUserList
// @Tags user-service
// @Router /user/list [get]
func (service *UserService) GetUserList(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	users, err := userDao.GetUserList()
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(users)

}

// CreateUser
// @Tags user-service
// @Router /user [post]
func (service *UserService) CreateUser(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	snowFlake := utils.SnowFlake{}
	id := snowFlake.Generate()
	user := &model.User{
		ID:       strconv.FormatInt(id, 10),
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err := userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user := &model.User{
		Email:    service.Email,
		Password: service.Password,
	}
	err := userDao.Login(user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// UpdateUser
// @Tags user-service
// @Router /user [put]
func (service *UserService) UpdateUser(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user := &model.User{
		ID:       service.ID,
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err := userDao.UpdateUser(service.ID, user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// DeleteUserById
// @Tags user-service
// @Router /user/:id [delete]
func (service *UserService) DeleteUserById(ctx context.Context, id string) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	err := userDao.DeleteUserById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(nil)
}