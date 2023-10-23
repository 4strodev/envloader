package envloader

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type environmentField struct {
	name     string
	required bool
}

// Loads environment variables to struct fields
// struct fields cannot be complex data: slices, maps, structs, pointers, etc.
func Marshal(v any) (error) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Pointer {
		return errors.New("struct pointer expected")
	}

	structVal := val.Elem()
	if structVal.Kind() != reflect.Struct {
		return errors.New("struct pointer expected")
	}

	structType := structVal.Type()
	var envField environmentField = environmentField{name: "", required: true}
	for i := 0; i < structVal.NumField(); i++ {
		vfield := structVal.Field(i)
		tfield := structType.Field(i)

		// Cechking if field has env tag
		tagString, exists := tfield.Tag.Lookup("env")
		if !exists {
			envField.name = tfield.Name
		} else {
			var err error
			envField, err = tagToEnvField(tagString)
			if err != nil {
				return err
			}

			if envField.name == "" {
				envField.name = tfield.Name
			}
		}

		// Reading env variable
		envValue := os.Getenv(envField.name)
		if envValue == "" {
			if envField.required {
				return fmt.Errorf("'%s' is not defined and is required", envField.name)
			}
			continue
		}
		// Parse string to corresponding type
		switch tfield.Type.Kind() {
		case reflect.Int:
			parsedVal, err := strconv.ParseInt(envValue, 10, 32)
			if err != nil {
				return err
			}
			vfield.SetInt(parsedVal)
		case reflect.Int8:
			parsedVal, err := strconv.ParseInt(envValue, 10, 8)
			if err != nil {
				return err
			}
			vfield.SetInt(parsedVal)
		case reflect.Int16:
			parsedVal, err := strconv.ParseInt(envValue, 10, 16)
			if err != nil {
				return err
			}
			vfield.SetInt(parsedVal)
		case reflect.Int32:
			parsedVal, err := strconv.ParseInt(envValue, 10, 32)
			if err != nil {
				return err
			}
			vfield.SetInt(parsedVal)
		case reflect.Int64:
			parsedVal, err := strconv.ParseInt(envValue, 10, 64)
			if err != nil {
				return err
			}
			vfield.SetInt(parsedVal)
		case reflect.Uint:
			parsedVal, err := strconv.ParseUint(envValue, 10, 32)
			if err != nil {
				return err
			}
			vfield.SetUint(parsedVal)
		case reflect.Uint8:
			parsedVal, err := strconv.ParseUint(envValue, 10, 8)
			if err != nil {
				return err
			}
			vfield.SetUint(parsedVal)
		case reflect.Uint16:
			parsedVal, err := strconv.ParseUint(envValue, 10, 16)
			if err != nil {
				return err
			}
			vfield.SetUint(parsedVal)
		case reflect.Uint32:
			parsedVal, err := strconv.ParseUint(envValue, 10, 32)
			if err != nil {
				return err
			}
			vfield.SetUint(parsedVal)
		case reflect.Uint64:
			parsedVal, err := strconv.ParseUint(envValue, 10, 64)
			if err != nil {
				return err
			}
			vfield.SetUint(parsedVal)
		case reflect.Bool:
			parsedVal, err := strconv.ParseBool(envValue)
			if err != nil {
				return err
			}
			vfield.SetBool(parsedVal)
		case reflect.Float32:
			parsedVal, err := strconv.ParseFloat(envValue, 32)
			if err != nil {
				return err
			}
			vfield.SetFloat(parsedVal)
		case reflect.Float64:
			parsedVal, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				return err
			}
			vfield.SetFloat(parsedVal)
		case reflect.String:
			vfield.SetString(envValue)
		default:
			return fmt.Errorf("field '%s' of type %s' is not admited", tfield.Name, tfield.Type.Kind())
		}
	}

	return nil
}

// Parse `env` tag
func tagToEnvField(tag string) (environmentField, error) {
	envField := environmentField{}
	splitTag := strings.Split(tag, ",")

	if len(splitTag) == 0 {
		return envField, errors.New("empty tag not allowed")
	}

	for i, v := range splitTag {
		splitTag[i] = strings.Trim(v, "\\s\\n\\t")
	}

	if len(splitTag) == 2 {
		envField.name = splitTag[0]
		if splitTag[1] != "required" {
			return envField, fmt.Errorf("env tag not valid expected 'required' found '%s'", splitTag[1])
		}
		envField.required = true
	} else if len(splitTag) == 1 {
		if splitTag[0] == "required" {
			envField.required = true
		} else {
			envField.name = splitTag[0]
		}
	} else {
		return envField, fmt.Errorf("invalid env tag")
	}

	return envField, nil
}
