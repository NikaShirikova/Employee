package module

type Company struct {
	ID    uint   `json:"id"`
	Name  string `gorm:"unique" json:"nameCompany"`
	Phone string `gorm:"unique" json:"phoneCompany"`
}
