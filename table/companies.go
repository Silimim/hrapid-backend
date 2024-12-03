package table

import (
	"net/http"
	"reflect"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func CompaniesHandler(w http.ResponseWriter, r *http.Request) {
	companies, err := fetchCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	table, err := buildCompaniesTable(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, table)
}

func fetchCompanies() ([]model.Company, error) {
	var companies []model.Company
	result := db.GetDB().Find(&companies)
	return companies, result.Error
}

func buildCompaniesTable(companies []model.Company) (*AutoTable, error) {
	var companyModel model.Company
	headers := buildCompaniesHeaders(reflect.ValueOf(&companyModel).Elem())

	table := &AutoTable{
		Headers: headers,
		Data:    make([]interface{}, len(companies)),
	}

	for i, company := range companies {
		table.Data[i] = company
	}

	return table, nil
}

func buildCompaniesHeaders(val reflect.Value) []AutoTableHeader {
	headers := make([]AutoTableHeader, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		header := buildCompaniesHeader(field)
		headers = append(headers, header)
	}

	return headers
}

func buildCompaniesHeader(field reflect.StructField) AutoTableHeader {
	fieldName := field.Name
	fieldType := field.Type.String()

	header := AutoTableHeader{
		Header:     utils.SplitCamelCase(fieldName),
		Field:      utils.ToSnakeCase(fieldName),
		Type:       fieldType,
		InputType:  "text",
		Required:   !utils.IsPointerType(fieldType),
		FormatType: getCompaniesFormatType(fieldName),
	}

	if fieldName == "Status" {
		header.InputType = "select"
	} else if fieldName == "Sales" {
		header.InputType = "number"
	}

	return header
}

func getCompaniesFormatType(fieldName string) AutoTableFormat {
	var formatType AutoTableFormat

	switch fieldName {
	case "Status":
		enumType := "enum"
		formatType.Type = &enumType
		formatType.Enum = &[]EnumType{
			{"ACTIVE", "success"},
			{"INACTIVE", "danger"},
			{"PENDING", "warning"},
			{"TERMINATED", "secondary"},
		}
	case "Sales":
		currencyType := "currency"
		formatType.Type = &currencyType
	}

	return formatType
}
