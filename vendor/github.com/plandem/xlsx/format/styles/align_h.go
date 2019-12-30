// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for HAlignType
const (
	_ primitives.HAlignType = iota
	HAlignGeneral
	HAlignLeft
	HAlignCenter
	HAlignRight
	HAlignFill
	HAlignJustify
	HAlignCenterContinuous
	HAlignDistributed
)

func init() {
	primitives.FromHAlignType = map[primitives.HAlignType]string{
		HAlignGeneral:          "general",
		HAlignLeft:             "left",
		HAlignCenter:           "center",
		HAlignRight:            "right",
		HAlignFill:             "fill",
		HAlignJustify:          "justify",
		HAlignCenterContinuous: "centerContinuous",
		HAlignDistributed:      "distributed",
	}

	primitives.ToHAlignType = make(map[string]primitives.HAlignType, len(primitives.FromHAlignType))
	for k, v := range primitives.FromHAlignType {
		primitives.ToHAlignType[v] = k
	}
}
