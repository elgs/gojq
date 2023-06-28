package gojq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/elgs/gosplitargs"
)

// JQ (JSON Query) struct
type JQ struct {
	Data any
}

// NewFileQuery - Create a new &JQ from a JSON file.
func NewFileQuery(jsonFile string) (*JQ, error) {
	raw, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}
	var data = new(any)
	err = json.Unmarshal(raw, data)
	if err != nil {
		return nil, err
	}
	return &JQ{*data}, nil
}

// NewStringQuery - Create a new &JQ from a raw JSON string.
func NewStringQuery(jsonString string) (*JQ, error) {
	var data = new(any)
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return nil, err
	}
	return &JQ{*data}, nil
}

// NewQuery - Create a &JQ from an any parsed by json.Unmarshal
func NewQuery(jsonObject any) *JQ {
	return &JQ{Data: jsonObject}
}

// Query queries against the JSON with the expression passed in. The exp is separated by dots (".")
func (jq *JQ) Query(exp string) (any, error) {
	if exp == "." {
		return jq.Data, nil
	}
	paths, err := gosplitargs.SplitArgs(exp, ".", false)
	if err != nil {
		return nil, err
	}
	var context any = jq.Data
	for _, path := range paths {
		if len(path) >= 3 && strings.HasPrefix(path, "[") && strings.HasSuffix(path, "]") {
			// array
			index, err := strconv.Atoi(path[1 : len(path)-1])
			if err != nil {
				return nil, err
			}
			if v, ok := context.([]any); ok {
				if len(v) <= index {
					return nil, errors.New(fmt.Sprint(path, " index out of range."))
				}
				context = v[index]
			} else {
				return nil, errors.New(fmt.Sprint(path, " is not an array. ", v))
			}
		} else {
			// map
			if v, ok := context.(map[string]any); ok {
				if val, ok := v[path]; ok {
					context = val
				} else {
					return nil, errors.New(fmt.Sprint(path, " does not exist."))
				}
			} else {
				return nil, errors.New(fmt.Sprint(path, " is not an object. ", v))
			}
		}
	}
	return context, nil
}

// QueryToMap queries against the JSON with the expression passed in, and convert to a map[string]any
func (jq *JQ) QueryToMap(exp string) (map[string]any, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return nil, errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.(map[string]any); ok {
		return ret, nil
	}
	return nil, errors.New("Failed to convert to map: " + exp)
}

// QueryToMap queries against the JSON with the expression passed in, and convert to a array: []any
func (jq *JQ) QueryToArray(exp string) ([]any, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return nil, errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.([]any); ok {
		return ret, nil
	}
	return nil, errors.New("Failed to convert to array: " + exp)
}

// QueryToMap queries against the JSON with the expression passed in, and convert to string
func (jq *JQ) QueryToString(exp string) (string, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return "", errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.(string); ok {
		return ret, nil
	}
	return "", errors.New("Failed to convert to string: " + exp)
}

// QueryToMap queries against the JSON with the expression passed in, and convert to int64
func (jq *JQ) QueryToInt64(exp string) (int64, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return 0, errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.(float64); ok {
		return int64(ret), nil
	}
	return 0, errors.New("Failed to convert to int64: " + exp)
}

// QueryToMap queries against the JSON with the expression passed in, and convert to float64
func (jq *JQ) QueryToFloat64(exp string) (float64, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return 0, errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.(float64); ok {
		return ret, nil
	}
	return 0, errors.New("Failed to convert to float64: " + exp)
}

// QueryToMap queries against the JSON with the expression passed in, and convert to bool
func (jq *JQ) QueryToBool(exp string) (bool, error) {
	r, err := jq.Query(exp)
	if err != nil {
		return false, errors.New("Failed to parse: " + exp)
	}
	if ret, ok := r.(bool); ok {
		return ret, nil
	}
	return false, errors.New("Failed to convert to float64: " + exp)
}
