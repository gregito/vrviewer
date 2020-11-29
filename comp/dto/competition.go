package dto

import "github.com/gregito/vrviewer/comp/model"

type Competition struct {
	ID   int64
	Name string
	Year int64
	Type model.ClimbingType
}
