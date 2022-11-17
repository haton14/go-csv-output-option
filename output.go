package csvton

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

var errUnexpectedType = "expected type is struct or struct slice. has type :%s"

func (opt Option) Output(fileName string, data any) error {
	dataType := reflect.TypeOf(data).Kind()
	var csvWriteData [][]string
	switch dataType {
	case reflect.Struct:
		csvWriteData = append(csvWriteData, opt.convertRow(data))
	case reflect.Slice:
		dataSliceType := reflect.TypeOf(data).Elem().Kind()
		if dataSliceType != reflect.Struct {
			return fmt.Errorf(errUnexpectedType, fmt.Sprintf("%s slice", dataSliceType))
		}
		dataLen := reflect.ValueOf(data).Len()
		csvWriteData = make([][]string, 0, dataLen)
		for i := 0; i < dataLen; i++ {
			csvWriteData = append(csvWriteData, opt.convertRow(reflect.ValueOf(data).Index(i)))
		}
	default:
		return fmt.Errorf(errUnexpectedType, dataType)
	}
	if err := output(fileName, csvWriteData); err != nil {
		return err
	}
	return nil
}

func (opt Option) convertRow(data any) []string {
	v := reflect.ValueOf(data)
	rowCount := reflect.TypeOf(data).NumField()
	rowCSV := make([]string, 0, rowCount)
	for i := 0; i < rowCount; i++ {
		tagValue := reflect.TypeOf(data).Field(i).Tag.Get(tagName)
		if tagValue == "" || opt.value[tagValue] {
			rowCSV = append(rowCSV, fmt.Sprint(v.Field(i)))
		}
	}
	return rowCSV
}

func output(fileName string, csvWriteData [][]string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("csv file create error. %w", err)
	}
	w := csv.NewWriter(f)
	for _, record := range csvWriteData {
		if err := w.Write(record); err != nil {
			return fmt.Errorf("csv file write error. %w", err)
		}
	}
	w.Flush()
	return nil
}
