// A generic TSV parser. Parses values using structs and reflection.
package tparser

import(
	"reflect"
	"fmt"
	"strconv"
	"bufio"
	"io"
	"strings"
)

// Populates a value's fields with the values in slice s. Value is assumed to be
// a struct.
func fill(value reflect.Value, s []string) error {
	// Check number of fields.
	if len(s) < value.NumField() {
		return fmt.Errorf("Not enough values to populate all fields (%d/%d).",
				len(s), value.NumField())
	}
	
	// Go over fields.
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		kind := field.Kind()
		
		if !field.CanSet() {
			return fmt.Errorf("Field %d: Cannot be set. Is it unexported?", i)
		}
		
		// Assign value according to type.
		switch {
		case kind == reflect.String:
			field.SetString(s[i])
		
		case kind >= reflect.Int && kind <= reflect.Int64:
			v, err := strconv.ParseInt(s[i], 0, 64)
			if err != nil {
				return fmt.Errorf("Field %d: %v", i, err)
			}
			field.SetInt(v)
		
		case kind >= reflect.Uint && kind <= reflect.Uint64:
			v, err := strconv.ParseUint(s[i], 0, 64)
			if err != nil {
				return fmt.Errorf("Field %d: %v", i, err)
			}
			field.SetUint(v)
		
		case kind == reflect.Float64 || kind == reflect.Float32:
			v, err := strconv.ParseFloat(s[i], 64)
			if err != nil {
				return fmt.Errorf("Field %d: %v", i, err)
			}
			field.SetFloat(v)
		
		case kind == reflect.Bool:
			v, err := strconv.ParseBool(s[i])
			if err != nil {
				return fmt.Errorf("Field %d: %v", i, err)
			}
			field.SetBool(v)
		
		default:
			return fmt.Errorf("Field %d: Unsupported field type: %s", i,
					kind.String())
		}
	}
	
	return nil
}

// A scanner like bufio.Scanner, that can also parse tables.
type Scanner struct {
	*bufio.Scanner
	typ reflect.Type
}

// Returns a new scanner. typ is a pointer to a struct, that represents the type
// that should be returned by the parser.
func NewScanner(r io.Reader, typ interface{}) *Scanner {
	if !isStructPtr(typ) {
		panic("Argument must be a pointer to a struct.")
	}
	return &Scanner{bufio.NewScanner(r), reflect.ValueOf(typ).Elem().Type()}
}

// Returns the parsed object from the last read line.
func (s *Scanner) Obj() (interface{}, error) {
	a := reflect.New(s.typ)
	err := fill(a.Elem(), strings.Split(s.Text(), "\t"))
	if err != nil {
		return nil, err
	}
	return a.Interface(), nil
}

// Scans all objects from the given reader. typ is a pointer to a struct of the
// type that should be returned. Will return a slice of the given type. Reading
// is buffered.
func ScanAll(r io.Reader, typ interface{}) (interface{}, error) {
	if !isStructPtr(typ) {
		panic("Argument must be a pointer to a struct.")
	}
	
	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(typ)), 0, 0)
	s := NewScanner(r, typ)
	for s.Scan() {
		obj, err := s.Obj()
		if err != nil {
			return nil, err
		}
		result = reflect.Append(result, reflect.ValueOf(obj))
	}
	
	return result.Interface(), nil
}

// Checks if the given thing is a pointer to a struct.
func isStructPtr(a interface{}) bool {
	value := reflect.ValueOf(a)
	return value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct
}
