package main

import (
	"Employee/internal/database/postgresql"
	"Employee/internal/module"
	"fmt"
	"gorm.io/gorm"
)

func main() {
	db := postgresql.Init()
	employee := &module.Employee{
		Name:         "Roman",
		Surname:      "Toropkin",
		Phone:        "89995550202",
		Passport:     module.Passport{PassType: "РФ", Number: "8989565656"},
		DepartmentID: 1,
	}

	result := AddEmployee(employee, db)
	fmt.Println("ID добавленного сотрудника: ", result)
}

func AddEmployee(employee *module.Employee, db *gorm.DB) uint {
	result := db.Create(&employee)
	if result.Error != nil {
		panic("Failed to add employee")
	}
	return employee.ID
}
