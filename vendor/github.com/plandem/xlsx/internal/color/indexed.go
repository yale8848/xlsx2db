// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package color

var (
	//list of built-in indexed colors
	indexed []string
)

func init() {
	indexed = []string{
		Normalize("#000000"),
		Normalize("#FFFFFF"),
		Normalize("#FF0000"),
		Normalize("#00FF00"),
		Normalize("#0000FF"),
		Normalize("#FFFF00"),
		Normalize("#FF00FF"),
		Normalize("#00FFFF"),
		Normalize("#000000"),
		Normalize("#FFFFFF"),
		Normalize("#FF0000"),
		Normalize("#00FF00"),
		Normalize("#0000FF"),
		Normalize("#FFFF00"),
		Normalize("#FF00FF"),
		Normalize("#00FFFF"),
		Normalize("#800000"),
		Normalize("#008000"),
		Normalize("#000080"),
		Normalize("#808000"),
		Normalize("#800080"),
		Normalize("#008080"),
		Normalize("#C0C0C0"),
		Normalize("#808080"),
		Normalize("#9999FF"),
		Normalize("#993366"),
		Normalize("#FFFFCC"),
		Normalize("#CCFFFF"),
		Normalize("#660066"),
		Normalize("#FF8080"),
		Normalize("#0066CC"),
		Normalize("#CCCCFF"),
		Normalize("#000080"),
		Normalize("#FF00FF"),
		Normalize("#FFFF00"),
		Normalize("#00FFFF"),
		Normalize("#800080"),
		Normalize("#800000"),
		Normalize("#008080"),
		Normalize("#0000FF"),
		Normalize("#00CCFF"),
		Normalize("#CCFFFF"),
		Normalize("#CCFFCC"),
		Normalize("#FFFF99"),
		Normalize("#99CCFF"),
		Normalize("#FF99CC"),
		Normalize("#CC99FF"),
		Normalize("#FFCC99"),
		Normalize("#3366FF"),
		Normalize("#33CCCC"),
		Normalize("#99CC00"),
		Normalize("#FFCC00"),
		Normalize("#FF9900"),
		Normalize("#FF6600"),
		Normalize("#666699"),
		Normalize("#969696"),
		Normalize("#003366"),
		Normalize("#339966"),
		Normalize("#003300"),
		Normalize("#333300"),
		Normalize("#993300"),
		Normalize("#993366"),
		Normalize("#333399"),
		Normalize("#333333"),
	}
}
