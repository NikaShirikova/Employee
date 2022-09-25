package postgresql

import (
	"Employee/internal/handler"
	"Employee/internal/module"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ListPostger struct {
	db *gorm.DB
}

func NewListPostgres(db *gorm.DB) *ListPostger {
	return &ListPostger{db: db}
}

func (pstg *ListPostger) AddEmployee(empl *module.Employee) (uint, error) {
	result := pstg.db.Debug()
	if empl.Department.ID == 0 {
		result.Omit("Department")
	}
	if empl.CompanyID == 0 {
		result.Omit("Company")
	}
	result.Create(&empl)
	if result.Error != nil {
		handler.LoggerZap.Error("Error when trying to add an employee",
			zap.Error(result.Error))
		return 0, result.Error
	}

	handler.LoggerZap.Info(
		"ID add employee %s ",
		zap.Uint("ID ", empl.ID))
	return empl.ID, nil
}

func (pstg *ListPostger) GetIDCOmpanyByName(company string) uint {
	var input module.Company
	result := pstg.db.Where("name = ?", company).First(&input)
	if result.Error != nil {
		handler.LoggerZap.Info(
			"Employee of such a company not found",
			zap.String("Company ", company))
		return 0
	}
	handler.LoggerZap.Info(
		"Employee of such a company has been found",
		zap.String("Company ", company),
		zap.Uint("ID ", input.ID))
	return input.ID
}

func (pstg *ListPostger) GetIDCDepartmentByName(department string) uint {
	var input module.Department
	result := pstg.db.Where("name = ?", department).First(&input)
	if result.Error != nil {
		handler.LoggerZap.Info(
			"Employee of such a department not found",
			zap.String("Department ", department))
		return 0
	}
	handler.LoggerZap.Info(
		"Employee of such a department has been found",
		zap.String("Department ", department),
		zap.Uint("ID ", input.ID))
	return input.ID
}

func (pstg *ListPostger) DeleteEmployee(id uint) error {
	result := pstg.db.Delete(&module.Employee{}, id)
	if result.Error != nil {
		handler.LoggerZap.Error("Error when trying to delete an employee",
			zap.Uint("ID ", id))
		return result.Error
	}
	handler.LoggerZap.Info(
		"Employee with the selected ID has been deleted",
		zap.Uint("ID ", id))
	return nil
}

func (pstg *ListPostger) UpdateEmployee(empl *module.Employee) error {
	result := pstg.db.Model(&empl).Updates(&empl)
	if result.Error != nil {
		handler.LoggerZap.Error(
			"Error when trying to update an employee data",
			zap.Error(result.Error))
		return result.Error
	}
	handler.LoggerZap.Info("Employee data has been updated")
	return nil
}

func (pstg *ListPostger) UpdateDepartment(depart *module.Department) error {
	result := pstg.db.Model(&depart).Updates(&depart)
	if result.Error != nil {
		handler.LoggerZap.Error(
			"Department data has not been updated",
			zap.String("Department ", depart.Name),
			zap.Error(result.Error))
		return result.Error
	}
	handler.LoggerZap.Info("Department data has been updated",
		zap.String("Department ", depart.Name))
	return nil
}

func (pstg *ListPostger) UpdateCompany(comp *module.Company) error {
	result := pstg.db.Model(&comp).Updates(&comp)
	if result.Error != nil {
		handler.LoggerZap.Error(
			"Company data has not been updated",
			zap.String("Company ", comp.Name),
			zap.Error(result.Error))
		return result.Error
	}
	handler.LoggerZap.Info("Company data has been updated",
		zap.String("Company ", comp.Name))
	return nil
}

func (pstg *ListPostger) UpdatePassport(pass *module.Passport) error {
	result := pstg.db.Model(&pass).Updates(&pass)
	if result.Error != nil {
		handler.LoggerZap.Error(
			"Passport data has not been updated",
			zap.Error(result.Error))
		return result.Error
	}
	handler.LoggerZap.Info("Passport data has been updated")
	return nil
}

func (pstg *ListPostger) GetListEmployeeByCompany(company string) (module.Employee, error) {
	var input module.Employee
	result := pstg.db.Debug().
		Joins("Join empl.passports on passports.id = employees.passport_id", &module.Passport{}).
		Joins("Join empl.departments on departments.id = employees.department_id", &module.Department{}).
		Joins("Join empl.companies on companies.id = employees.company_id", &module.Company{}).
		Where("companies.name =?", company).
		Preload("Passport").
		Preload("Company").
		Preload("Department").
		Find(&input)
	if result.Error != nil && result.RowsAffected == 0 {
		handler.LoggerZap.Info(
			"Error when trying to issue Employee Data",
			zap.String("Company", company),
			zap.Error(result.Error))
		return input, result.Error
	}
	handler.LoggerZap.Info(
		"Employees data issued",
		zap.String("Company", company))
	return input, nil
}

func (pstg *ListPostger) GetListEmployeeByDepartment(department string) (module.Employee, error) {
	var input module.Employee
	result := pstg.db.Debug().
		Joins("Join empl.passports on passports.id = employees.passport_id", &module.Passport{}).
		Joins("Join empl.departments on departments.id = employees.department_id", &module.Department{}).
		Joins("Join empl.companies on companies.id = employees.company_id", &module.Company{}).
		Where("departments.name =?", department).
		Preload("Passport").
		Preload("Company").
		Preload("Department").
		Find(&input)
	if result.Error != nil && result.RowsAffected == 0 {
		handler.LoggerZap.Info(
			"Error when trying to issue Employee Data",
			zap.String("Department", department),
			zap.Error(result.Error))
		return input, result.Error
	}
	handler.LoggerZap.Info(
		"Employees data issued",
		zap.String("Department", department))
	return input, nil
}
