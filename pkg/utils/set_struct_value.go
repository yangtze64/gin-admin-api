package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type StructFieldOption = func(tField reflect.StructField, vField reflect.Value) error

func SetStructValue(v interface{}, opts ...StructFieldOption) error {
	valueOf := reflect.Indirect(reflect.ValueOf(v))
	typeOf := valueOf.Type()
	if len(opts) == 0 {
		opts = append(opts, WithSetStructFieldDefault())
	}
	return setStructValue(typeOf, valueOf, opts...)
}

func setStructValue(typeOf reflect.Type, valueOf reflect.Value, opts ...StructFieldOption) error {
	if typeOf.Kind() == reflect.Struct {
		for i := 0; i < typeOf.NumField(); i++ {
			vField := valueOf.Field(i)
			tField := typeOf.Field(i)
			if vField.IsValid() {
				valof := reflect.Indirect(vField)
				typof := valof.Type()
				if typof.Kind() == reflect.Struct || typof.Kind() == reflect.Slice || typof.Kind() == reflect.Map {
					if err := setStructValue(typof, valof, opts...); err != nil {
						return err
					}
				} else {
					for _, opt := range opts {
						if err := opt(tField, vField); err != nil {
							return err
						}
					}
				}
			}
		}
	} else if typeOf.Kind() == reflect.Slice {
		lens := valueOf.Len()
		if lens > 0 {
			for i := 0; i < lens; i++ {
				valof := reflect.Indirect(valueOf.Index(i))
				typof := valof.Type()
				if err := setStructValue(typof, valof, opts...); err != nil {
					return err
				}
			}
		}
	} else if typeOf.Kind() == reflect.Map {
		if !valueOf.IsNil() {
			keys := valueOf.MapKeys()
			for _, k := range keys {
				valof := reflect.Indirect(valueOf.MapIndex(k))
				typof := valof.Type()
				if err := setStructValue(typof, valof, opts...); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func setStructFieldDefault(tField reflect.StructField, vField reflect.Value) error {
	valueOf := reflect.Indirect(vField)
	typeOf := valueOf.Type()
	tag, ok := tField.Tag.Lookup("default")
	if !ok || !valueOf.IsZero() {
		return nil
	}
	switch typeOf.Kind() {
	case reflect.String:
		vField.SetString(tag)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tagVal, err := strconv.ParseInt(tag, 10, 0)
		if err != nil {
			return err
		}
		vField.SetInt(tagVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tagVal, err := strconv.ParseInt(tag, 10, 0)
		if err != nil {
			return err
		}
		vField.SetUint(uint64(tagVal))
	case reflect.Float32, reflect.Float64:
		tagVal, err := strconv.ParseFloat(tag, 10)
		if err != nil {
			return err
		}
		vField.SetFloat(tagVal)
	case reflect.Bool:
		tagVal, err := strconv.ParseBool(tag)
		if err != nil {
			return err
		}
		vField.SetBool(tagVal)
	}
	return nil
}

func WithSetStructFieldDefault() func(tField reflect.StructField, vField reflect.Value) error {
	return setStructFieldDefault
}

func verifyStructFieldRequired(tField reflect.StructField, vField reflect.Value) error {
	valueOf := reflect.Indirect(vField)
	tag, ok := tField.Tag.Lookup("required")
	isBool, _ := strconv.ParseBool(tag)
	if ok && isBool && valueOf.IsZero() {
		return errors.New(fmt.Sprintf("`%s` tag exist `required` but its value is zero", tField.Name))
	}
	return nil
}

func WithVerifyStructFieldRequired() func(tField reflect.StructField, vField reflect.Value) error {
	return verifyStructFieldRequired
}

func verifyStructFieldRange(tField reflect.StructField, vField reflect.Value) error {
	valueOf := reflect.Indirect(vField)
	typeOf := valueOf.Type()
	val, isNumber := getNumberValue(typeOf, valueOf)
	if !isNumber {
		return nil
	}
	tag, ok := tField.Tag.Lookup("range")
	if !ok {
		return nil
	}
	if !strings.Contains(tag, ":") {
		return errors.New(fmt.Sprintf("struct field `%s` tag `range` [%s] format error,not contain `:`", typeOf.Name(), tag))
	}
	slice := strings.Split(tag, ":")
	if len(slice) != 2 {
		return errors.New(fmt.Sprintf("struct field `%s` tag `range` [%s] format error,slice len is not 2", typeOf.Name(), tag))
	}
	var min, max float64
	var err error
	if slice[0] != "" {
		min, err = strconv.ParseFloat(slice[0], 10)
		if err != nil {
			return err
		}
	}
	if slice[1] != "" {
		max, err = strconv.ParseFloat(slice[1], 10)
		if err != nil {
			return err
		}
	}
	if min != 0 && val < min {
		return errors.New(fmt.Sprintf("struct field `%s` tag `range` [%s] format error,'%v' not in range [%s]", typeOf.Name(), tag, val, tag))
	}
	if max != 0 && val > max {
		return errors.New(fmt.Sprintf("struct field `%s` tag `range` [%s] format error,'%v' not in range [%s]", typeOf.Name(), tag, val, tag))
	}
	return nil
}

func WithVerifyStructFieldRange() func(tField reflect.StructField, vField reflect.Value) error {
	return verifyStructFieldRange
}

func verifyStructFieldOptions(tField reflect.StructField, vField reflect.Value) error {
	valueOf := reflect.Indirect(vField)
	typeOf := valueOf.Type()
	tag, ok := tField.Tag.Lookup("options")
	if !ok || tag == "" {
		return nil
	}
	exist := false
	var valz interface{}
	value, isNumber := getNumberValue(typeOf, valueOf)
	valz = value
	tagSlice := strings.Split(tag, ",")
	for _, val := range tagSlice {
		if isNumber {
			vfloat, err := strconv.ParseFloat(val, 10)
			if err != nil {
				return err
			}
			if vfloat == value {
				exist = true
				break
			}
		} else if typeOf.Kind() == reflect.String {
			valz = valueOf.String()
			if val == valueOf.String() {
				exist = true
				break
			}
		}
	}
	if !exist {
		return errors.New(fmt.Sprintf("struct field `%s` tag `range` [%s] format error,field value '%v' not contain [%s]", typeOf.Name(), tag, valz, tag))
	}
	return nil
}

func WithVerifyStructFieldOptions() func(tField reflect.StructField, vField reflect.Value) error {
	return verifyStructFieldOptions
}

func getNumberValue(typeOf reflect.Type, valueOf reflect.Value) (float64, bool) {
	var (
		value float64 = 0
		ok            = true
	)
	switch typeOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = float64(valueOf.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = float64(valueOf.Uint())
	case reflect.Float32, reflect.Float64:
		value = valueOf.Float()
	default:
		ok = false
	}
	return value, ok
}
