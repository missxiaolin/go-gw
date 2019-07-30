package model

type Category struct {
	ID			uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name		string		`gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Description	string		`gorm:"type:varchar(255);not null;" json:"description"`
	Pid			string		`gorm:"type:tinyint;not null;" json:"pid"`
	Timestamps
}

func (c *Category) getTable() string {
	return "category"
}