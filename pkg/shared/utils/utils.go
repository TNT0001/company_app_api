package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"go-api/internal/pkg/domain/domain_model/dto"
	"io/ioutil"
	"reflect"
	"strconv"
)

// ConvertImportantToLevel func
func ConvertImportantToLevel(total float32) float32 {
	if (float32(0.0) < total) && (float32(0.2) >= total) {
		return 1
	}
	if (float32(0.2) < total) && (float32(0.4) >= total) {
		return 2
	}
	if (float32(0.4) < total) && (float32(0.6) >= total) {
		return 3
	}
	if (float32(0.6) < total) && (float32(0.8) >= total) {
		return 4
	}
	if (float32(0.8) < total) && (float32(1.0) >= total) {
		return 5
	}
	return 0
}

// GetValueFieldByName func
func GetValueFieldByName(v interface{}, field string) interface{} {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}

// EncryptPassword method SHA256 func
func EncryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}

// parse []project to csv file
func CSVFromPojects(projects dto.UserProjectsResponse) (string, error) {
	file, err := ioutil.TempFile("", "projects.csv")
	if err != nil {
		return "", err
	}

	defer file.Close()

	// get name of field in struct
	fileName := []string{"name", "category", "projected_spend", "projected_variance", "revenue_recognised"}

	// make slice
	temp := make([][]string, 0)
	temp = append(temp, fileName)

	for _, project := range projects.Projects {
		fieldValue := []string{project.Name, project.Category, strconv.Itoa(project.ProjectedSpend),
			strconv.Itoa(project.ProjectedVariance), strconv.Itoa(project.RevenueRecognised)}
		temp = append(temp, fieldValue)
	}

	w := csv.NewWriter(file)

	err = w.WriteAll(temp)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
