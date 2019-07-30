package model

type Users struct {
	ID			uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name		string		`gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Password	string		`gorm:"type:varchar(255);not null;" json:"password"`
	Status		uint		`gorm:"type:tinyint;not null;" json:"status"`
	Timestamps
}