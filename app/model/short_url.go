package model

type ShortUrl struct {
	Id        int64  `gorm:"column:id; primaryKey"`
	SourceUrl string `gorm:"column:source_ur"`
}

func (s *ShortUrl) TableName() string {
	return "short_url"
}
