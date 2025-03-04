package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (a *AuthorityService) CreateAuthorityList(authorities []*model.SysAuthority) error {
	q := query.SysAuthority.WithContext(context.Background())
	return q.Create(authorities...)
}

func (a *AuthorityService) AddAuthorityMenus(authorityMenus []*model.SysAuthorityMenu) error {
	q := query.SysAuthorityMenu
	return q.WithContext(context.Background()).Create(authorityMenus...)
}
