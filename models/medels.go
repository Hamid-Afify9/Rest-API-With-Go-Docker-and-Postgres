package models
import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model

	Questions string `json:"Questions" gorm:"text;not null;default:NULL"`  
	Answer 	  string `json:"Answer" gorm:"text;not null;default:NULL"`
}
