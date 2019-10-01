package KeyValueDB

import (
	"reflect"
)

/*
 * Abdulazeez Abdulazeez Adeshina
 */

func (k *ValueObject) IsNumber(arg int) bool {
	return reflect.ValueOf(arg).Kind() == reflect.Int
}

func (k *ValueObject) IsString(arg string) bool {
	return reflect.ValueOf(arg).Kind() == reflect.String
}

func (k *ValueObject) IsObject(arg struct{}) bool {
	return reflect.ValueOf(arg).Kind() == reflect.Struct
}

func (k *ValueObject) IsBoolean(arg bool) bool {
	return reflect.ValueOf(arg).Kind() == reflect.Bool
}

func (k *ValueObject) IsDefined(arg string) bool {
	if arg != "" {
		return true
	}
	return false
}

func (k *ValueObject) IsChar(arg string) bool {
	return reflect.ValueOf(arg).Kind() == reflect.String && len(arg) == 1
}
