package csvton

import "reflect"

type Option struct {
	value map[string]bool
}

const tagName string = "csv"

func ParseOption(opt any) Option {
	count := reflect.TypeOf(opt).NumField()
	parse := make(map[string]bool, count)
	value := reflect.ValueOf(opt)
	for i := 0; i < count; i++ {
		parse[reflect.TypeOf(opt).Field(i).Tag.Get(tagName)] = value.Field(i).Bool()
	}
	return Option{
		value: parse,
	}
}
