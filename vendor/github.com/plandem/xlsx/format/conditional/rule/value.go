// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"errors"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/plandem/xlsx/internal/number_format/convert"
	//"reflect"
	"strconv"
	"time"
)

type valueRule struct {
	baseRule
}

//Value is helper object to set specific options for rule
var Value valueRule

func (x valueRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Value
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeCellIs,
		}
	}
}

func (x valueRule) fromInt(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func (x valueRule) fromUInt(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func (x valueRule) fromFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func (x valueRule) fromBool(value bool) string {
	if value {
		return "1"
	}

	return "0"
}

//nolint
func (x valueRule) setValue(r *Info, values []interface{}, operator primitives.ConditionOperatorType, s *styles.Info) {
	x.initIfRequired(r)
	r.rule.Operator = operator

	for _, v := range values {
		var s string
		switch value := v.(type) {
		case int:
			s = x.fromInt(value)
		case int8:
			s = x.fromInt(int(value))
		case int16:
			s = x.fromInt(int(value))
		case int32:
			s = x.fromInt(int(value))
		case int64:
			s = x.fromInt(int(value))
		case uint:
			s = x.fromUInt(value)
		case uint8:
			s = x.fromUInt(uint(value))
		case uint16:
			s = x.fromUInt(uint(value))
		case uint32:
			s = x.fromUInt(uint(value))
		case uint64:
			s = x.fromUInt(uint(value))
		case float32:
			s = x.fromFloat(float64(value))
		case float64:
			s = x.fromFloat(float64(value))
		case []byte:
			s = string(value)
		case bool:
			s = x.fromBool(value)
		case time.Time:
			s = value.Format(convert.ISO8601)
		case string:
			s = value
		default:
			s = fmt.Sprintf("%v", value)
		}

		//if value has '=', then remove it
		if len(s) > 0 && s[0] == '=' {
			s = s[1:]
		}

		if len(s) > 0 {
			r.rule.Formula = append(r.rule.Formula, ml.Formula(s))
		}
	}

	r.style = s

}

func (x valueRule) Between(from, to interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		x.setValue(r, []interface{}{from, to}, primitives.ConditionOperatorBetween, s)
	}
}

func (x valueRule) NotBetween(from, to interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{from, to}, primitives.ConditionOperatorNotBetween, s)
	}
}

func (x valueRule) Equal(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorEqual, s)
	}
}

func (x valueRule) NotEqual(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorNotEqual, s)
	}
}

func (x valueRule) Greater(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorGreaterThan, s)
	}
}

func (x valueRule) Less(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorLessThan, s)
	}
}

func (x valueRule) GreaterOrEqual(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorGreaterThanOrEqual, s)
	}
}

func (x valueRule) LessOrEqual(value interface{}, s *styles.Info) Option {
	return func(r *Info) {
		x.setValue(r, []interface{}{value}, primitives.ConditionOperatorLessThanOrEqual, s)
	}
}

func (x valueRule) Validate(r *Info) error {
	if len(r.rule.Formula) == 0 || len(r.rule.Formula[0]) == 0 {
		return errors.New("value: no criteria or value for rule")
	}

	return nil
}
