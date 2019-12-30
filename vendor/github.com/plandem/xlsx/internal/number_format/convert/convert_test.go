// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToBool(t *testing.T) {
	v, err := ToBool("true")
	require.Nil(t, err)
	require.Equal(t, true, v)

	v, err = ToBool("1")
	require.Nil(t, err)
	require.Equal(t, true, v)

	v, err = ToBool("0")
	require.Nil(t, err)
	require.Equal(t, false, v)

	v, err = ToBool("false")
	require.Nil(t, err)
	require.Equal(t, false, v)

	_, err = ToBool("dsds3wr")
	require.NotNil(t, err)
}

func TestToInt(t *testing.T) {
	v, err := ToInt("1")
	require.Nil(t, err)
	require.Equal(t, 1, v)

	v, err = ToInt("-1")
	require.Nil(t, err)
	require.Equal(t, -1, v)

	_, err = ToInt("12345.12345")
	require.NotNil(t, err)

	_, err = ToInt("fgfert3")
	require.NotNil(t, err)
}

func TestToFloat(t *testing.T) {
	v, err := ToFloat("1")
	require.Nil(t, err)
	require.Equal(t, 1.0, v)

	v, err = ToFloat("-1")
	require.Nil(t, err)
	require.Equal(t, -1.0, v)

	v, err = ToFloat("12345.12345")
	require.Nil(t, err)
	require.Equal(t, 12345.12345, v)

	_, err = ToFloat("fgfert3")
	require.NotNil(t, err)
}

func TestToDate(t *testing.T) {
	v, err := ToDate("43308.7047106481")
	require.Nil(t, err)
	require.Equal(t, "2018-07-27 16:54:46 +0000 UTC", v.UTC().String())

	v, err = ToDate("2018-07-27T17:40:58")
	require.Nil(t, err)
	require.Equal(t, "2018-07-27 17:40:58 +0000 UTC", v.UTC().String())

	_, err = ToDate("dsdsds")
	require.NotNil(t, err)
}
