package module

type Department struct {
	ID    uint   `json:"id"`
	Name  string `gorm:"unique" json:"nameDepartment"`
	Phone string `gorm:"unique" json:"phoneDepartment"`
}
