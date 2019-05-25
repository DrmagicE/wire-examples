package service

import (
	"log"
	"database/sql"

	"github.com/google/wire"
)

// UserRepository
type UserRepository interface {
	AddUser()
}

// userRepo UserRepository接口实现
type userRepo struct {
	*sql.DB
}

// AddUser 新增user
func (u *userRepo) AddUser() {
	log.Println("add user")
}

// NewUserRepo provider: 根据*sql.DB初始化 *userRepo
func NewUserRepo(db *sql.DB) *userRepo{
	return &userRepo{}
}

// UserSet 声明NewUserRepo的返回值是UserRepository接口类型
var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*userRepo)))