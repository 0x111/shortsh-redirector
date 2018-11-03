package models

import (
	"time"
)

type ShortShUrl struct {
	Id        int64     `xorm:"'id' index pk autoincr" json:"id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
	UpdatedAt time.Time `xorm:"updated" json:"-"`

	Url         string `xorm:"url" json:"url,required"`
	ShortDomain string `xorm:"short_domain" json:"short_domain"`
	ShortId     string `xorm:"short_id" json:"short_id"`
}
