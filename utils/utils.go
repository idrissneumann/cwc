package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PromptUserForValue() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	value = strings.TrimSuffix(value, "\n")
	return value
}

func PrintJson(class interface{}) {
	marchal_class, err := json.Marshal(class)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", JsonPrettyPrint(string(marchal_class)))
}

func PrintHeader(class interface{}) {
	values := reflect.ValueOf(class)
	typesOf := values.Type()
	headerMsg := ""
	for i := 0; i < values.NumField(); i++ {
		headerMsg = headerMsg + typesOf.Field(i).Name + "\t"
	}
	fmt.Println(headerMsg)
}
func PrintRow(class interface{}) {
	PrintHeader(class)
	values := reflect.ValueOf(class)
	valuesMsg := ""
	for i := 0; i < values.NumField(); i++ {
		valuesMsg = valuesMsg + fmt.Sprintf("%v\t", values.Field(i).Interface())
	}
	fmt.Println(valuesMsg)
}

func PrintMultiRow(type_class interface{}, class interface{}) {
	PrintHeader(type_class)
	s := reflect.ValueOf(class)
	for i := 0; i < s.Len(); i++ {
		v := reflect.Indirect(s.Index(i))
		valuesMsg := ""
		for i := 0; i < v.NumField(); i++ {
			valuesMsg = valuesMsg + fmt.Sprintf("%v\t", v.Field(i).Interface())
		}
		fmt.Println(valuesMsg)
	}
}
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
