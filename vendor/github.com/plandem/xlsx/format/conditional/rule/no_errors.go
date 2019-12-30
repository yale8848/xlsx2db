// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

type noErrorsRule struct {
	baseRule
}

//NoErrors is helper object to set specific options for rule
var NoErrors noErrorsRule

func (x noErrorsRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = NoErrors
		r.rule = &ml.ConditionalRule{
			Type:    primitives.ConditionTypeNotContainsErrors,
			Formula: []ml.Formula{`NOT(ISERROR(:cell:))`},
		}
	}
}

func (x noErrorsRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
