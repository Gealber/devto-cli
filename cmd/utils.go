package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//extractTagsName make sure ptr is a pointer
func extractTagsName(ptr interface{}) []string {
	fields := make([]string, 0)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) //a reflect.StructField
		tag := fieldInfo.Tag
		name := tag.Get("json")
		fields = append(fields, name)
	}
	return fields
}

//processInput make sure ptr is a pointer
func processInput(ptr interface{}) error {
	tagsFields := extractTagsName(ptr)
	v := reflect.ValueOf(ptr).Elem()
	if len(tagsFields) != v.NumField() {
		return errors.New("Number of fields differ from extracted")
	}

	reader := bufio.NewReader(os.Stdin)
	for i, field := range tagsFields {
		displayFancy(field, v.Field(i).Type().String())
		variable, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		//ignoring delimiter '\n'
		variable = variable[:len(variable)-1]

		if field == "body_markdown" {
			data, err := loadFromFile(variable)
			if err != nil {
				return err
			}
			variable = data
		}

		//store variable just read into desired structure
		if err := populate(v.Field(i), variable); err != nil {
			return err
		}
	}

	return nil
}

func displayFancy(field, t string) {
	fmt.Printf("  \033[0;36m> %s\033[0m\033[0;32m[%s]\033[0m: ", field, t)
}

//populate takes care of setting a single field v
//taken from The Go Programming Language book
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int32:
		//only put if a number is provided
		if len(value) != 0 {
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return errors.New(fmt.Sprintf("Invalid input type expected int32: %v", err))
			}
			v.SetInt(i)
			return nil
		}
	case reflect.Int64:
		//only put if a number is provided
		if len(value) != 0 {
			i, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Invalid input type expected int32: %v", err))
			}
			v.SetInt(i)
			return nil
		}
	case reflect.Bool:
		if len(value) == 0 {
			v.SetBool(false)
			return nil
		}
		b, err := strconv.ParseBool(value)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid input type expected bool: %v", err))
		}
		v.SetBool(b)

	case reflect.Slice:
		//in case of a slice we assume
		//the values are provided as a coma separated
		values := strings.Split(value, ",")
		elem := reflect.New(v.Type().Elem()).Elem()
		for _, element := range values {
			//recursion here shouldn't cause problem
			//at least you have several nested slice
			if err := populate(elem, element); err != nil {
				return err
			}
			v.Set(reflect.Append(v, elem))
		}

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//loadFromFile load the content from a file
func loadFromFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(b[:]), nil
}
