package types

var JAVA_TYPE_MAP = map[string]string{
	"string":      "String",
	"boolean":     "boolean",
	"date":        "Date",
	"dateTime":    "Date",
	"double":      "double",
	"hovedklasse": "FintMainObject",
	"referanse":   "FintReference",
	"abstrakt":    "FintAbstractObject",
	"datatype":    "FintComplexDatatypeObject",
}


func GetJavaType(t string) string {

	value, ok := JAVA_TYPE_MAP[t]
	if ok {
		return value
	} else {
		return t
	}
}

