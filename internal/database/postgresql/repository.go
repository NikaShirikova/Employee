package postgresql

import (
	"Employee/internal/module"
	"gorm.io/gorm"
)

type ListDB interface {
	AddEmployee(*module.Employee) (uint, error)
	GetIDCOmpanyByName(string) uint
	GetIDCDepartmentByName(string) uint
	DeleteEmployee(uint) error
	UpdateEmployee(*module.Employee) error
	UpdateDepartment(*module.Department) error
	UpdateCompany(*module.Company) error
	UpdatePassport(*module.Passport) error
	GetListEmployeeByCompany(string) (module.Employee, error)
	GetListEmployeeByDepartment(string) (module.Employee, error)
}

type Repository struct {
	ListDB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ListDB: NewListPostgres(db),
	}
}
