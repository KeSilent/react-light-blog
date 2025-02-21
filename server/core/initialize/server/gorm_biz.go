package server

import "github.com/kesilent/react-light-blog/global"

func BizModel() error {
	db := global.RLB_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
