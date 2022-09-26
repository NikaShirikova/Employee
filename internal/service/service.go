package service

import (
	"employee/internal/database/postgresql"
	"employee/internal/module"
)

type Serv interface {
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
	Serv
}

func NewService(repos *postgresql.Repository) *Service {
	return &Service{
		Serv: NewListService(repos.DB),
	}
}
