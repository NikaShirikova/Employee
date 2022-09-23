package module

type Passport struct {
	ID       uint   `json:"id"`
	PassType string `json:"passportType"`
	Number   string `gorm:"unique" json:"passportNumber"`
}
