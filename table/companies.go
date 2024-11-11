package table

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

type CompanyHeader struct {
	Header string `json:"header"`
	Field  string `json:"field"`
	Type   string `json:"type"`
}

type CompanyTable struct {
	Headers []CompanyHeader `json:"headers"`
	Data    []model.Company `json:"data"`
}

func Companies(w http.ResponseWriter, r *http.Request) {

	var companyTable CompanyTable
	var companies []model.Company

	var companyModel model.Company

	db.GetDB().Find(&companies)

	val := reflect.ValueOf(&companyModel).Elem()

	headers := []CompanyHeader{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i).Name
		fieldName := utils.ToSnakeCase(field)
		headerName := utils.SplitCamelCase(field)
		fieldType := val.Type().Field(i).Type.String()

		headers = append(headers, CompanyHeader{
			Header: headerName,
			Field:  fieldName,
			Type:   fieldType,
		})
	}

	companyTable.Headers = headers
	companyTable.Data = companies

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(companyTable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
