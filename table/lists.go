package table

import (
	"net/http"
	"reflect"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func ListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := fetchLists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	table, err := buildListsTable(lists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, table)
}

func fetchLists() ([]model.List, error) {
	var lists []model.List
	result := db.GetDB().Find(&lists)
	return lists, result.Error
}

func buildListsTable(lists []model.List) (*AutoTable, error) {
	var listModel model.List
	headers := buildListsHeaders(reflect.ValueOf(&listModel).Elem())

	table := &AutoTable{
		Headers: headers,
		Data:    make([]interface{}, len(lists)),
	}

	for i, list := range lists {
		table.Data[i] = list
	}

	return table, nil
}

func buildListsHeaders(val reflect.Value) []AutoTableHeader {
	headers := make([]AutoTableHeader, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		header := buildListsHeader(field)
		headers = append(headers, header)
	}

	return headers
}

func buildListsHeader(field reflect.StructField) AutoTableHeader {
	fieldName := field.Name
	fieldType := field.Type.String()

	header := AutoTableHeader{
		Header:     utils.SplitCamelCase(fieldName),
		Field:      utils.ToSnakeCase(fieldName),
		Type:       fieldType,
		InputType:  "text",
		Required:   !utils.IsPointerType(fieldType),
		FormatType: getListsFormatType(fieldName),
	}

	return header
}

func getListsFormatType(_ string) AutoTableFormat {
	var formatType AutoTableFormat

	return formatType
}
