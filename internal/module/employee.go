package module

type Employee struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Phone        string     `gorm:"unique" json:"phone"`
	PassportID   uint       `json:"passid"`
	Passport     Passport   `json:"passport"`
	CompanyID    uint       `json:"companyid"`
	Company      Company    `json:"company"`
	DepartmentID uint       `json:"departid"`
	Department   Department `json:"department"`
}
