package models

type Item struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Details string `json:"details"`
}
