package csvton

import (
	"fmt"
	"reflect"
)

type Option struct {
	value map[string]bool
}

const tagName string = "csv"

func ParseOption(opt any) (*Option, error) {
	count := reflect.TypeOf(opt).NumField()
	parse := make(map[string]bool, count)
	value := reflect.ValueOf(opt)
	for i := 0; i < count; i++ {
		tagValue := reflect.TypeOf(opt).Field(i).Tag.Get(tagName)
		if tagValue == "" {
			continue
		}
		optionFiledType := value.Field(i).Kind()
		if optionFiledType != reflect.Bool {
			return nil, fmt.Errorf("option filed expected type is bool. has type %s", optionFiledType)
		}
		parse[tagValue] = value.Field(i).Bool()
	}
	return &Option{
		value: parse,
	}, nil
}
