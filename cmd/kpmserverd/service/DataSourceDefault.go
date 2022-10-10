package service

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"kpm/cmd/kpmserverd/dao/mysql"
)

type DataSourceDefault struct {
	mysql mysql.Mysql
}

func (d DataSourceDefault) SearchName(name string) string {
	//TODO implement me
	panic("implement me")
}

func (d DataSourceDefault) SearchSubPkgNames(SubPkgNames []string) string {
	//TODO implement me
	panic("implement me")
}

func (d DataSourceDefault) GetPublishVersion() {
	//TODO implement me
	panic("implement me")
}

// Search 搜索
func (d DataSourceDefault) Search(id []byte) string {
	pkgs, err := d.mysql.SearchPkg(string(id))
	if err != nil {
		return "1"
	}
	marshal, err := json.Marshal(pkgs)
	if err != nil {
		return "2"
	}
	return string(marshal)
}

// Publish 发布
func (d DataSourceDefault) Publish(id []byte) string {
	//TODO implement me
	//d.mysql.AddPkg()
	panic("implement me")
}

// Publish 发布

func NewDefault(db *sqlx.DB) (d DataSourceDefault) {
	d.mysql = mysql.NewMysql(db)
	return
}
