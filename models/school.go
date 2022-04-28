// models/school.go

package models

type School struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Rank       uint   `json:"rank"`
	Name       string `json:"title"`
	Conference string `json:"author"`
	Record     string `json:"record"`
	Road       string `json:"road"`
	Neutral    string `json:"neutral"`
	Home       string `json:"home"`
}
