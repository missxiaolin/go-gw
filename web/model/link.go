package model

type Link struct {
	ID		uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name	string		`gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Url		string		`gorm:"type:varchar(255);not null;" json:"url"`
	Sort	uint		`gorm:"type:tinyint;not null;" json:"sort"`
	Timestamps
}