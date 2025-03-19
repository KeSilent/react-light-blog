package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (a *AuthorityService) CreateAuthorityList(authorities []*model.SysRole) error {
	q := query.SysRole.WithContext(context.Background())
	return q.Create(authorities...)
}

func (a *AuthorityService) AddAuthorityMenus(authorityMenus []*model.SysRoleMenu) error {
	q := query.SysRoleMenu
	return q.WithContext(context.Background()).Create(authorityMenus...)
}
