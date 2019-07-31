package model

type Article struct {
	ID			uint		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Cid			uint		`gorm:"type:tinyint;not null;" json:"cid"`
	Title		string		`gorm:"type:varchar(255);not null;" json:"title"`
	Author		string		`gorm:"type:varchar(15);not null;" json:"author"`
	Content		string		`gorm:"type:text;not null;" json:"content"`
	Keywords	string		`gorm:"type:varchar(255);not null;" json:"keywords"`
	Description	string		`gorm:"type:varchar(255);not null;" json:"description"`
	Status		uint		`gorm:"type:tinyint;not null;" json:"status"`
	Timestamps
}

func (c *Article) getTable() string {
	return "article"
}