package models

import "time"

type Song struct {
	Id          string    `json:"id,omitempty" gorm:"column:id;type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"song,omitempty" example:"Supermassive Black Hole" gorm:"column:title"`
	Group       string    `json:"group,omitempty" example:"Muse" gorm:"column:group_song"`
	ReleaseDate time.Time `json:"releaseDate,omitempty" example:"Oh baby don't you know I suffer?" gorm:"column:release_date;type:date"`
	Text        string    `json:"text,omitempty" example:"19.06.2006" gorm:"column:text_song"`
	Link        string    `json:"link,omitempty" example:"https://www.youtube.com/watch?v=N-_mHedypEU" gorm:"column:link_song"`
}
