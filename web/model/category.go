package model

type Category struct {
	ID			uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name		string		`gorm:"type:varchar(50);unique_index;not null;" json:"name"`
	Pid			uint		`gorm:"type:tinyint;not null;" json:"pid"`
	Status		uint		`gorm:"type:tinyint;not null;" json:"status"`
	Timestamps
}

func (c *Category) getTable() string {
	return "category"
}