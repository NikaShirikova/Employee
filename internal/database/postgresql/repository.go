package postgresql

import (
	"employee/internal/module"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DB interface {
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
	DB
}

func NewRepository(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		DB: NewListPostgres(db, log),
	}
}
