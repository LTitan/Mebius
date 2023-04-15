package dag

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/itchyny/gojq"
)

var (
	templateRegexp = regexp.MustCompile(`\$\{\{.*?\}\}`)
)

const (
	LT          = "<"
	LET         = "<="
	Equal       = "="
	DoubleEqual = "=="
	GT          = ">"
	GET         = ">="
	Regex       = "regexp"
	IN          = "in"
	NotIn       = "notin"
)

func lint(str string) string {
	str = strings.Trim(str, "$")
	str = strings.Trim(str, "{")
	return strings.Trim(str, "}")
}

func warpMapValueToStr(ctx DagContext, raw string) (string, error) {
	matchedStr := templateRegexp.FindAllString(raw, -1)
	for _, str := range matchedStr {
		target := lint(str)
		value := ctx[target]
		if value == nil {
			continue
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Array, reflect.Map, reflect.Ptr:
			data, err := json.Marshal(value)
			if err != nil {
				return "", err
			}
			value = string(data)
		default:
			value = fmt.Sprint(value)
		}
		raw = strings.ReplaceAll(raw, str, value.(string))
	}
	return raw, nil
}

func warpMapValueToStrByJQ(ctx map[string]any, raw string) (string, error) {
	query, err := gojq.Parse(raw)
	if err != nil {
		return "", err
	}
	iter := query.Run(ctx)
	valueSet := []string{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return "", err
		}
		strv := fmt.Sprint(v)
		if v == nil {
			strv = ""
		}
		valueSet = append(valueSet, strv)
	}
	return valueSet[0], nil
}

func warpCondition(value any, expression string, expected string) (bool, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Map, reflect.Ptr:
		data, err := json.Marshal(value)
		if err != nil {
			return false, err
		}
		value = string(data)
	default:
		value = fmt.Sprint(value)
	}
	strValue := value.(string)
	switch expression {
	case LT, LET, GT, GET:
		left, _ := strconv.ParseFloat(strValue, 32)
		right, _ := strconv.ParseFloat(expected, 32)
		switch expression {
		case LT:
			return left < right, nil
		case LET:
			return left <= right, nil
		case GT:
			return left > right, nil
		case GET:
			return left >= right, nil
		}
	case IN, NotIn:
		flag := false
		for _, item := range strings.Split(expected, ",") {
			if item == strValue {
				flag = true
				break
			}
		}
		if expected == IN {
			return flag, nil
		}
		return !flag, nil
	case Regex:
		return regexp.MatchString(expected, strValue)
	case Equal, DoubleEqual:
		return strValue == expected, nil
	}
	return false, nil
}
