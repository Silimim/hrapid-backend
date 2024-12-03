package table

import (
	"net/http"
	"reflect"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := fetchEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	table, err := buildEmployeesTable(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, table)
}

func fetchEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	result := db.GetDB().Find(&employees)
	return employees, result.Error
}

func buildEmployeesTable(employees []model.Employee) (*AutoTable, error) {
	var employeeModel model.Employee
	headers := buildEmployeesHeaders(reflect.ValueOf(&employeeModel).Elem())

	table := &AutoTable{
		Headers: headers,
		Data:    make([]interface{}, len(employees)),
	}

	for i, employee := range employees {
		table.Data[i] = employee
	}

	return table, nil
}

func buildEmployeesHeaders(val reflect.Value) []AutoTableHeader {
	headers := make([]AutoTableHeader, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		header := buildEmployeesHeader(field)
		headers = append(headers, header)
	}

	return headers
}

func buildEmployeesHeader(field reflect.StructField) AutoTableHeader {
	fieldName := field.Name
	fieldType := field.Type.String()

	header := AutoTableHeader{
		Header:     utils.SplitCamelCase(fieldName),
		Field:      utils.ToSnakeCase(fieldName),
		Type:       fieldType,
		InputType:  "text",
		Required:   !utils.IsPointerType(fieldType),
		FormatType: getEmployeesFormatType(fieldName),
	}

	return header
}

func getEmployeesFormatType(_ string) AutoTableFormat {
	var formatType AutoTableFormat

	return formatType
}
