package sam

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
)

// Returns a map from tag name to its parsed (typed) value.
func parseTags(values []string) (map[string]any, error) {
	result := make(map[string]interface{}, len(values))
	for _, f := range values {
		parts, err := splitTag(f)
		if err != nil {
			return nil, err
		}
		switch parts[1] {
		case "A":
			if len(parts[2]) != 1 {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a single character",
					parts[1], parts[2])
			}
			result[parts[0]] = parts[2][0]
		case "i":
			x, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want an integer",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "f":
			x, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a number",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "Z":
			result[parts[0]] = parts[2]
		case "H":
			x, err := hex.DecodeString(parts[2])
			if err != nil {
				return nil, fmt.Errorf("illegal value for tag type %v: %q, "+
					"want a hexadecimal sequence",
					parts[1], parts[2])
			}
			result[parts[0]] = x
		case "B":
			// TODO(amit): Not implemented yet. Treating like string for now.
			result[parts[0]] = parts[2]
		default:
			return nil, fmt.Errorf("unrecognized tag type: %v, in tag %v",
				parts[1], f)
		}
	}
	return result, nil
}

// Splits a SAM tag by colon. Used instead of strings.SpliN for performance.
func splitTag(tag string) ([3]string, error) {
	colon1, colon2 := -1, -1
	for i, c := range tag {
		if c == ':' {
			if colon1 == -1 {
				colon1 = i
			} else {
				colon2 = i
				break
			}
		}
	}
	var result [3]string
	if colon2 == -1 {
		return result, fmt.Errorf("tag doesn't have at least 3 colons: %q", tag)
	}
	result[0] = tag[:colon1]
	result[1] = tag[colon1+1 : colon2]
	result[2] = tag[colon2+1:]
	return result, nil
}

// Returns the given tags in SAM format, sorted and tab-separated.
func tagsToText(tags map[string]interface{}) []string {
	texts := make([]string, 0, len(tags))
	for tag, val := range tags {
		texts = append(texts, tagToText(tag, val))
	}
	sort.Strings(texts)
	return texts
}

// Returns the SAM format representation of the given tag.
func tagToText(tag string, val interface{}) string {
	switch val := val.(type) {
	case byte:
		return tag + ":A:" + string(val)
	case int:
		return tag + ":i:" + strconv.Itoa(val)
	case float64:
		return tag + ":f:" + strconv.FormatFloat(val, 'e', -1, 64)
	case string:
		return tag + ":Z:" + val
	case []byte:
		return tag + ":H:" + hex.EncodeToString(val)
	default:
		panic(fmt.Sprintf("unsupported type for value %v", val))
	}
}
