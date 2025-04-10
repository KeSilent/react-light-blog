/*
 * @Author: Yang
 * @Date: 2025-02-21 16:29:20
 * @Description: 请填写简介
 */
package request

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/kesilent/react-light-blog/dal/model"
)

// CustomClaims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID     string
	ID       model.SnowflakeID
	Username string
	NickName string
	RoleId   string
}
