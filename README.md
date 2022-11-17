# go-csv-output-option
A simple Go library that optionally controls items to be output to CSV.

## Usage
```
package main

import csvton "github.com/haton14/go-csv-output-option"

type ExampleOption struct {
	HasName    bool `csv:"has_name"`
	HasAddress bool `csv:"has_address"`
}

type ExampleData struct {
	ID      int
	Name    string `csv:"has_name"`
	Address string `csv:"has_address"`
}

func main() {
	data := ExampleData{
		ID:      111,
		Name:    "taro",
		Address: "tokyo",
	}

	opt := csvton.ParseOption(
		ExampleOption{
			HasName:    true,
			HasAddress: false,
		},
	)

	opt.Output("examlpe.csv", data)
}
```
