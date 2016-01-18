package gojq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elgs/gosplitargs"
	"strconv"
	"strings"
)

// JQ (JSON Query) struct
type JQ struct {
	Data interface{}
}

// NewStringQuery - Create a new &JQ from a raw JSON string.
func NewStringQuery(jsonString string) (*JQ, error) {
	var data = new(interface{})
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return nil, err
	}
	return &JQ{*data}, nil
}

// NewQuery - Create a &JQ from an interface{} parsed by json.Unmarshal
func NewQuery(jsonObject interface{}) *JQ {
	return &JQ{Data: jsonObject}
}

// Query queries against the JSON with the expression passed in. The exp is separated by dots (".")
func (jq *JQ) Query(exp string) (interface{}, error) {
	paths, err := gosplitargs.SplitArgs(exp, "\\.", false)
	if err != nil {
		return nil, err
	}
	var context interface{} = jq.Data
	for _, path := range paths {
		if len(path) >= 3 && strings.HasPrefix(path, "[") && strings.HasSuffix(path, "]") {
			// array
			index, err := strconv.Atoi(path[1 : len(path)-1])
			if err != nil {
				return nil, err
			}
			switch v := context.(type) {
			case []interface{}:
				{
					if len(v) <= index {
						return nil, errors.New(fmt.Sprint(path, " index out of range."))
					}
					context = v[index]
				}
			default:
				return nil, errors.New(fmt.Sprint(path, " is not an array. ", v))
			}

		} else {
			// map
			switch v := context.(type) {
			case map[string]interface{}:
				if val, ok := v[path]; ok {
					context = val
				} else {
					return nil, errors.New(fmt.Sprint(path, " does not exist."))
				}
			default:
				return nil, errors.New(fmt.Sprint(path, " is not an object. ", v))
			}
		}
	}
	return context, nil
}
