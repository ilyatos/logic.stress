package helpers

import "fmt"

func PrintStructureWithFields(v interface{}) {
	fmt.Printf("%+v\n", v)
}
