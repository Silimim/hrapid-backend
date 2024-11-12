package table

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func Companies(w http.ResponseWriter, r *http.Request) {

	var companyTable AutoTable
	var companies []model.Company

	var companyModel model.Company

	db.GetDB().Find(&companies)

	val := reflect.ValueOf(&companyModel).Elem()

	headers := headerDescriptor(val)

	companyTable.Headers = headers
	companyTable.Data = make([]interface{}, len(companies))
	for i, c := range companies {
		companyTable.Data[i] = c
	}

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(companyTable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func headerDescriptor(val reflect.Value) []AutoTableHeader {
	headers := []AutoTableHeader{}
	for i := 0; i < val.NumField(); i++ {

		field := val.Type().Field(i).Name

		var fieldName = utils.ToSnakeCase(field)
		var headerName = utils.SplitCamelCase(field)
		var fieldType = val.Type().Field(i).Type.String()
		var formatType AutoTableFormat
		var inputType = "text"
		var required bool
		if fieldType[0] == '*' {
			required = false
		} else {
			required = true
		}

		if field == "Status" {
			enumType := "enum"

			formatType.Type = &enumType
			enum := []EnumType{
				{"ACTIVE", "success"},
				{"INACTIVE", "danger"},
				{"PENDING", "warning"},
				{"TERMINATED", "secondary"},
			}
			formatType.Enum = &enum

			inputType = "select"

		} else if field == "Sales" {
			currencyType := "currency"
			formatType.Type = &currencyType
			inputType = "number"
		}

		headers = append(headers, AutoTableHeader{
			Header:     headerName,
			Field:      fieldName,
			Type:       fieldType,
			FormatType: formatType,
			InputType:  inputType,
			Required:   required,
		})
	}
	return headers
}
