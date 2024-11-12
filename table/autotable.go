package table

type EnumType struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AutoTableFormat struct {
	Type *string     `json:"type"`
	Enum *[]EnumType `json:"enum"`
}

type AutoTableHeader struct {
	Header     string          `json:"header"`
	Field      string          `json:"field"`
	Type       string          `json:"type"`
	FormatType AutoTableFormat `json:"formatType"`
}

type AutoTable struct {
	Headers []AutoTableHeader `json:"headers"`
	Data    []interface{}     `json:"data"`
}
