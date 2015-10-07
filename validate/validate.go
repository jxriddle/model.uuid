// Package validate provides methods for validating params passed from the database row as interface{} types
package validate

import (
	"fmt"
	"time"
)

// Float returns the float value of param or 0.0
func Float(param interface{}) float64 {
	var v float64
	if param != nil {
		switch param.(type) {
		case int64:
			v = float64(param.(int64))
		default:
			v = param.(float64)
		}
	}
	return v
}

// Boolean returns the bool value of param or false
func Boolean(param interface{}) bool {
	var v bool
	if param != nil {
		v = param.(bool)
	}
	return v
}

// Int returns the int value of param or 0
func Int(param interface{}) int64 {
	var v int64
	if param != nil {
		v = param.(int64)
	}
	return v
}

// String returns the string value of param or ""
func String(param interface{}) string {
	var v string
	if param != nil {
		v = param.(string)
	}
	return v
}

// Time returns the time value of param or the zero value of time.Time
func Time(param interface{}) time.Time {
	var v time.Time
	if param != nil {
		v = param.(time.Time)
	}
	return v
}

// Length validates a param by min and max length
func Length(param string, min int, max int) error {
	length := len(param)
	if min != -1 && length < min {
		return fmt.Errorf("Length of string %s %d, expected > %d", param, length, min)
	}
	if max != -1 && length > max {
		return fmt.Errorf("Length of string %s %d, expected < %d", param, length, max)
	}
	return nil
}
