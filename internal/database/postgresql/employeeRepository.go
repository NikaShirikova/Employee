package postgresql

import (
	"Employee/internal/module"
	"fmt"
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
		return 0, result.Error
	}

	fmt.Println("ID добавленного сотрудника: ", empl.ID)
	return empl.ID, nil
}

func (pstg *ListPostger) GetIDCOmpanyByName(company string) uint {
	var input module.Company
	result := pstg.db.Where("name = ?", company).First(&input)
	if result.Error != nil {
		return 0
	}
	return input.ID
}

func (pstg *ListPostger) GetIDCDepartmentByName(department string) uint {
	var input module.Department
	result := pstg.db.Where("name = ?", department).First(&input)
	if result.Error != nil {
		return 0
	}
	return input.ID
}

func (pstg *ListPostger) DeleteEmployee(id uint) error {
	result := pstg.db.Delete(&module.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Сотрудник с выбранным ID удален.")
	return nil
}

func (pstg *ListPostger) UpdateEmployee(empl *module.Employee) error {
	result := pstg.db.Model(&empl).Updates(&empl)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Данные сотрудника обновлены.")
	return nil
}

func (pstg *ListPostger) UpdateDepartment(depart *module.Department) error {
	result := pstg.db.Model(&depart).Updates(&depart)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Данные сотрудника обновлены.")
	return nil
}

func (pstg *ListPostger) UpdateCompany(comp *module.Company) error {
	result := pstg.db.Model(&comp).Updates(&comp)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Данные сотрудника обновлены.")
	return nil
}

func (pstg *ListPostger) UpdatePassport(pass *module.Passport) error {
	result := pstg.db.Model(&pass).Updates(&pass)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Данные сотрудника обновлены.")
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
		return input, result.Error
	}
	fmt.Println("Данные сотрудников выданы.")
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
		return input, result.Error
	}
	fmt.Println("Данные сотрудников выданы.")
	return input, nil
}
