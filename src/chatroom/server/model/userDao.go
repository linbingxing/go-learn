package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

func (ud *UserDao) GetUserById(conn redis.Conn, userId int) (user User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	json.Unmarshal([]byte(res), &user)
	return
}

func (ud *UserDao) Login(userId int, userPwd string) (user User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
