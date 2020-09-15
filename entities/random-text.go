package entities

// RandomText ...
type RandomText struct {
	ID      uint   `gorm:"column:id;primaryKey"`
	Text    string `gorm:"column:random_text"`
	TextSum string `gorm:"column:random_text_sum"`
}

// TableName ...
func (RandomText) TableName() string {
	return "example.random_text"
}
