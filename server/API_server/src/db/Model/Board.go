package model

import (
	"time"
)

type Board struct {
	Id            int       `json:"id"              gorm:"PRIMARY_KEY"`
	CreatedUserId string    `json:"created_user_id" gorm:"type:varchar(255); NOT NULL"`
	Title         string    `json:"title"           gorm:"type:varchar(255); NOT NULL"`
	Publish       bool      `json:"publish"         gorm:"type:boolean; NOT NULL"`
	BgImg         string    `json:"bg_img"          gorm:"type:varchar(255); NOT NULL"`
	LastAccess    time.Time `json:"last_access"     gorm:"type:datetime; NOT NULL"`
	InviteURL     string    `json:"invite_url"      gorm:"invite_url;type:varchar(255); NOT NULL"`
	Kanbans       []Kanban  `json:"-"               gorm:"foreignkey:BoardId"`
	Authors       []*User   `json:"-"               gorm:"many2many:user_boards;"`
}
