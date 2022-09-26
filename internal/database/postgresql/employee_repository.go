package postgresql

import (
	"employee/internal/module"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Postgres struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewListPostgres(db *gorm.DB, log *zap.Logger) *Postgres {
	return &Postgres{
		db:  db,
		log: log.Named("postgresql")}
}

func (pstg *Postgres) AddEmployee(empl *module.Employee) (uint, error) {
	result := pstg.db.Debug()
	if empl.Department.ID == 0 {
		result.Omit("Department")
	}
	if empl.CompanyID == 0 {
		result.Omit("Company")
	}
	result.Create(&empl)
	if result.Error != nil {
		pstg.log.Error("Error when trying to add an employee",
			zap.Error(result.Error))
		return 0, result.Error
	}

	pstg.log.Info(
		"ID add employee %s ",
		zap.Uint("ID ", empl.ID))
	return empl.ID, nil
}

func (pstg *Postgres) GetIDCOmpanyByName(company string) uint {
	var input module.Company
	result := pstg.db.Where("name = ?", company).First(&input)
	if result.Error != nil {
		pstg.log.Info(
			"Employee of such a company not found",
			zap.String("Company ", company))
		return 0
	}
	pstg.log.Info(
		"Employee of such a company has been found",
		zap.String("Company ", company),
		zap.Uint("ID ", input.ID))
	return input.ID
}

func (pstg *Postgres) GetIDCDepartmentByName(department string) uint {
	var input module.Department
	result := pstg.db.Where("name = ?", department).First(&input)
	if result.Error != nil {
		pstg.log.Info(
			"Employee of such a department not found",
			zap.String("Department ", department))
		return 0
	}
	pstg.log.Info(
		"Employee of such a department has been found",
		zap.String("Department ", department),
		zap.Uint("ID ", input.ID))
	return input.ID
}

func (pstg *Postgres) DeleteEmployee(id uint) error {
	result := pstg.db.Delete(&module.Employee{}, id)
	if result.Error != nil {
		pstg.log.Error("Error when trying to delete an employee",
			zap.Uint("ID ", id))
		return result.Error
	}
	pstg.log.Info(
		"Employee with the selected ID has been deleted",
		zap.Uint("ID ", id))
	return nil
}

func (pstg *Postgres) UpdateEmployee(empl *module.Employee) error {
	result := pstg.db.Model(&empl).Updates(&empl)
	if result.Error != nil {
		pstg.log.Error(
			"Error when trying to update an employee data",
			zap.Error(result.Error))
		return result.Error
	}
	pstg.log.Info("Employee data has been updated")
	return nil
}

func (pstg *Postgres) UpdateDepartment(depart *module.Department) error {
	result := pstg.db.Model(&depart).Updates(&depart)
	if result.Error != nil {
		pstg.log.Error(
			"Department data has not been updated",
			zap.String("Department ", depart.Name),
			zap.Error(result.Error))
		return result.Error
	}
	pstg.log.Info("Department data has been updated",
		zap.String("Department ", depart.Name))
	return nil
}

func (pstg *Postgres) UpdateCompany(comp *module.Company) error {
	result := pstg.db.Model(&comp).Updates(&comp)
	if result.Error != nil {
		pstg.log.Error(
			"Company data has not been updated",
			zap.String("Company ", comp.Name),
			zap.Error(result.Error))
		return result.Error
	}
	pstg.log.Info("Company data has been updated",
		zap.String("Company ", comp.Name))
	return nil
}

func (pstg *Postgres) UpdatePassport(pass *module.Passport) error {
	result := pstg.db.Model(&pass).Updates(&pass)
	if result.Error != nil {
		pstg.log.Error(
			"Passport data has not been updated",
			zap.Error(result.Error))
		return result.Error
	}
	pstg.log.Info("Passport data has been updated")
	return nil
}

func (pstg *Postgres) GetListEmployeeByCompany(company string) (module.Employee, error) {
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
		pstg.log.Info(
			"Error when trying to issue Employee Data",
			zap.String("Company", company),
			zap.Error(result.Error))
		return input, result.Error
	}
	pstg.log.Info(
		"Employees data issued",
		zap.String("Company", company))
	return input, nil
}

func (pstg *Postgres) GetListEmployeeByDepartment(department string) (module.Employee, error) {
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
		pstg.log.Info(
			"Error when trying to issue Employee Data",
			zap.String("Department", department),
			zap.Error(result.Error))
		return input, result.Error
	}
	pstg.log.Info(
		"Employees data issued",
		zap.String("Department", department))
	return input, nil
}
