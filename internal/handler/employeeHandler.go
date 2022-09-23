package handler

import (
	"Employee/internal/module"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) AddEmpl(c *gin.Context) {
	var input module.Employee
	if err := c.BindJSON(&input); err != nil {
		fmt.Println("Error get data")
		return
	}

	emplId, err := h.services.ListServ.AddEmployee(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusResponse{"error"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": emplId,
	})
}

func (h *Handler) DeleteEmpl(c *gin.Context) {
	emplId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error get data")
	}

	err = h.services.ListServ.DeleteEmployee(uint(emplId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusResponse{"error"})
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) UpdateEmpl(c *gin.Context) {
	var input module.Employee
	if err := c.BindJSON(&input); err != nil {
		fmt.Println("Error get data")
		return
	}

	err := h.services.ListServ.UpdateEmployee(&input)
	errPass := h.services.ListServ.UpdatePassport(&input.Passport)
	errComp := h.services.ListServ.UpdateCompany(&input.Company)
	errDep := h.services.ListServ.UpdateDepartment(&input.Department)

	if err != nil && errPass != nil && errComp != nil && errDep != nil {
		c.JSON(http.StatusInternalServerError, statusResponse{"error"})
		return
	}

	c.JSON(http.StatusOK, statusResponse{"update ok"})
}

func (h *Handler) GetCompany(c *gin.Context) {
	name := c.Param("nameCompany")

	empl, err := h.services.ListServ.GetListEmployeeByCompany(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusResponse{"error"})
	} else if empl.ID == 0 {
		c.JSON(http.StatusNotFound, statusResponse{fmt.Sprintf("Employees not found with company name %s", name)})
	} else {
		c.IndentedJSON(http.StatusOK, empl)
	}
}

func (h *Handler) GetDepartment(c *gin.Context) {
	name := c.Param("nameDepartment")

	empl, err := h.services.ListServ.GetListEmployeeByDepartment(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusResponse{"error"})
	} else if empl.ID == 0 {
		c.JSON(http.StatusNotFound, statusResponse{fmt.Sprintf("Employees not found with department name %s", name)})
	} else {
		c.IndentedJSON(http.StatusOK, empl)
	}

}
