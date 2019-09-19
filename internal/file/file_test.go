// Create by Yale 2019/9/19 12:11
package file

import "testing"

func TestXlsx_GetDataByIndex(t *testing.T) {

	xf := NewFile()
	p := `C:\Users\Yale\Desktop\易学卡分成和保底分成合作单位明细(1).xlsx`
	xf.GetDataByIndex(p, []int{1}, func(i int, strings map[int]string) error {

		return nil
	})

}
