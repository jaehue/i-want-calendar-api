package models

import "github.com/go-xorm/xorm"

func Init(db *xorm.Engine) error {
	return db.Sync(new(Member), new(Application))
}
