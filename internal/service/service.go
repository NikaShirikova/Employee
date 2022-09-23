package service

import (
	"Employee/internal/database/postgresql"
	"Employee/internal/module"
)

type ListServ interface {
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

type Service struct {
	ListServ
}

func NewService(repos *postgresql.Repository) *Service {
	return &Service{
		ListServ: NewListService(repos.ListDB),
	}
}
