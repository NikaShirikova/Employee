package service

import (
	"employee/internal/database/postgresql"
	"employee/internal/module"
)

type ListService struct {
	repos postgresql.DB
}

func NewListService(repos postgresql.DB) *ListService {
	return &ListService{repos: repos}
}

func (serv *ListService) AddEmployee(empl *module.Employee) (uint, error) {
	empl.Company.ID = serv.repos.GetIDCOmpanyByName(empl.Company.Name)
	empl.Department.ID = serv.repos.GetIDCDepartmentByName(empl.Department.Name)
	return serv.repos.AddEmployee(empl)
}

func (serv *ListService) DeleteEmployee(id uint) error {
	return serv.repos.DeleteEmployee(id)
}

func (serv *ListService) UpdateEmployee(empl *module.Employee) error {
	return serv.repos.UpdateEmployee(empl)
}

func (serv *ListService) UpdateDepartment(depart *module.Department) error {
	return serv.repos.UpdateDepartment(depart)
}

func (serv *ListService) UpdateCompany(comp *module.Company) error {
	return serv.repos.UpdateCompany(comp)
}

func (serv *ListService) UpdatePassport(pass *module.Passport) error {
	return serv.repos.UpdatePassport(pass)
}

func (serv *ListService) GetListEmployeeByCompany(company string) (module.Employee, error) {
	return serv.repos.GetListEmployeeByCompany(company)
}

func (serv *ListService) GetListEmployeeByDepartment(department string) (module.Employee, error) {
	return serv.repos.GetListEmployeeByDepartment(department)
}

func (serv *ListService) GetIDCOmpanyByName(company string) uint {
	return serv.repos.GetIDCOmpanyByName(company)
}

func (serv *ListService) GetIDCDepartmentByName(department string) uint {
	return serv.repos.GetIDCDepartmentByName(department)
}
