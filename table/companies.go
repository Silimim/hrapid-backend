package table

import (
	"encoding/json"
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

	table, err := buildCompanyTable(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, table)
}

func fetchCompanies() ([]model.Company, error) {
	var companies []model.Company
	result := db.GetDB().Find(&companies)
	return companies, result.Error
}

func buildCompanyTable(companies []model.Company) (*AutoTable, error) {
	var companyModel model.Company
	headers := buildHeaders(reflect.ValueOf(&companyModel).Elem())

	table := &AutoTable{
		Headers: headers,
		Data:    make([]interface{}, len(companies)),
	}

	for i, company := range companies {
		table.Data[i] = company
	}

	return table, nil
}

func buildHeaders(val reflect.Value) []AutoTableHeader {
	headers := make([]AutoTableHeader, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		header := buildHeader(field)
		headers = append(headers, header)
	}

	return headers
}

func buildHeader(field reflect.StructField) AutoTableHeader {
	fieldName := field.Name
	fieldType := field.Type.String()

	header := AutoTableHeader{
		Header:     utils.SplitCamelCase(fieldName),
		Field:      utils.ToSnakeCase(fieldName),
		Type:       fieldType,
		InputType:  "text",
		Required:   !isPointerType(fieldType),
		FormatType: getFormatType(fieldName),
	}

	if fieldName == "Status" {
		header.InputType = "select"
	} else if fieldName == "Sales" {
		header.InputType = "number"
	}

	return header
}

func getFormatType(fieldName string) AutoTableFormat {
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

func isPointerType(fieldType string) bool {
	return len(fieldType) > 0 && fieldType[0] == '*'
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
