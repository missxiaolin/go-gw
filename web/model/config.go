package model

type Config struct {
	ID		uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name	string		`gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Value	string		`gorm:"type:varchar(255);not null;" json:"value"`
	Timestamps
}