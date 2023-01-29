package conf

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 解析预制配置
func presetConf(v interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(v))
	typ := val.Type()
	return presetStruct(val, typ)
}

func presetVal(fv reflect.Value, ft reflect.Type) error {
	if ft.Kind() == reflect.Struct {
		if err := presetStruct(fv, ft); err != nil {
			return err
		}

	} else if ft.Kind() == reflect.Slice {
		if err := presetSlice(fv, ft); err != nil {
			return err
		}
	} else if ft.Kind() == reflect.Map {
		if err := presetMap(fv, ft); err != nil {
			return err
		}
	}
	return nil
}

func presetStruct(val reflect.Value, typ reflect.Type) error {
	if typ.Kind() != reflect.Struct {
		return errors.New("preset struct parse fail,val is not `struct`")
	}
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		fv := reflect.Indirect(field)
		if fv.IsValid() {
			ft := fv.Type()
			tfield := typ.Field(i)
			requiredTag := tfield.Tag.Get("required")
			defaultTag, defaultOk := tfield.Tag.Lookup("default")
			if requiredTag != "" && !defaultOk && requiredTag != "0" && strings.ToLower(requiredTag) != "false" && fv.IsZero() {
				return errors.New("preset struct field tag exists `required` but value is zero")
			}
			if ft.Kind() == reflect.Struct || ft.Kind() == reflect.Slice || ft.Kind() == reflect.Map {
				if err := presetVal(fv, ft); err != nil {
					return err
				}
			} else {
				switch ft.Kind() {
				case reflect.String:
					if fv.IsZero() && defaultOk {
						field.SetString(defaultTag)
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if fv.IsZero() && defaultOk {
						defaultTagInt64, err := strconv.ParseInt(defaultTag, 10, 0)
						if err != nil {
							return err
						}
						field.SetInt(defaultTagInt64)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if fv.IsZero() && defaultOk {
						defaultTagInt64, err := strconv.ParseInt(defaultTag, 10, 0)
						if err != nil {
							return err
						}
						field.SetUint(uint64(defaultTagInt64))
					}
				case reflect.Float32, reflect.Float64:
					if fv.IsZero() && defaultOk {
						defaultTagFloat64, err := strconv.ParseFloat(defaultTag, 10)
						if err != nil {
							return err
						}
						field.SetFloat(defaultTagFloat64)
					}
				case reflect.Bool:
					defaultTagBool, err := strconv.ParseBool(defaultTag)
					if err != nil {
						return err
					}
					field.SetBool(defaultTagBool)
				default:
					return nil
				}
				if err := checkStructFieldTag(val, typ, i); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func presetSlice(val reflect.Value, typ reflect.Type) error {
	if typ.Kind() != reflect.Slice {
		return errors.New("preset slice parse fail,val is not `Slice`")
	}
	lens := val.Len()
	if lens > 0 {
		for i := 0; i < lens; i++ {
			fv := reflect.Indirect(val.Index(i))
			ft := fv.Type()
			if err := presetVal(fv, ft); err != nil {
				return err
			}
		}
	}

	return nil
}

func presetMap(val reflect.Value, typ reflect.Type) error {
	if typ.Kind() != reflect.Map {
		return errors.New("preset map parse fail,val is not `Map`")
	}
	if !val.IsNil() {
		keys := val.MapKeys()
		for _, k := range keys {
			fv := reflect.Indirect(val.MapIndex(k))
			ft := fv.Type()
			if err := presetVal(fv, ft); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkStructFieldTag(val reflect.Value, typ reflect.Type, structIndex int) error {
	field := val.Field(structIndex)
	ft := field.Type()
	isNumber := true
	var fieldVal float64
	switch ft.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldVal = float64(field.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldVal = float64(field.Uint())
	case reflect.Float32, reflect.Float64:
		fieldVal = field.Float()
	case reflect.String:
		isNumber = false
	default:
		return nil
	}

	tfield := typ.Field(structIndex)
	rangeTag := tfield.Tag.Get("range")
	if isNumber && rangeTag != "" {
		if !strings.Contains(rangeTag, ":") {
			return errors.New(fmt.Sprintf("struct field tag `range` format error,rangeTag [%s] not contain `:`", rangeTag))
		}
		rangeSlice := strings.Split(rangeTag, ":")
		if len(rangeSlice) != 2 {
			return errors.New(fmt.Sprintf("struct field tag `range` format error,rangeTag [%s] slice len is not 2", rangeTag))
		}
		var min, max float64
		var err error
		if rangeSlice[0] != "" {
			min, err = strconv.ParseFloat(rangeSlice[0], 10)
			if err != nil {
				return err
			}
		}
		if rangeSlice[1] != "" {
			max, err = strconv.ParseFloat(rangeSlice[1], 10)
			if err != nil {
				return err
			}
		}
		if min != 0 && fieldVal < min {
			return errors.New(fmt.Sprintf("struct field tag `range` format error,field value '%v' not in range [%s]", field, rangeTag))
		}
		if max != 0 && fieldVal > max {
			return errors.New(fmt.Sprintf("struct field tag `range` format error,field value '%v' not in range [%s]", field, rangeTag))
		}
	}

	optionsTag := tfield.Tag.Get("options")
	if optionsTag != "" {
		exist := false
		optionsSlice := strings.Split(optionsTag, ",")
		for _, val := range optionsSlice {
			if isNumber {
				vfloat, err := strconv.ParseFloat(val, 10)
				if err != nil {
					return err
				}
				if vfloat == fieldVal {
					exist = true
					break
				}
			} else {
				if val == field.String() {
					exist = true
					break
				}
			}
		}
		if !exist {
			return errors.New(fmt.Sprintf("struct field tag `range` format error,field value '%v' not contain [%s]", field, optionsTag))
		}
	}

	return nil
}
