package module

type Employee struct {
	ID           uint
	Name         string
	Surname      string
	Phone        string
	PassportID   uint
	Passport     Passport
	DepartmentID uint
	Department   Department
}
