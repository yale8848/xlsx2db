// Create by Yale 2019/9/16 13:48
package file

import (
	"github.com/plandem/xlsx"
	"net/url"
	"strings"
)

type Title struct {
	Title string `json:"title"`
	Index int    `json:"index"`
}
type TableFileInfo struct {
	Index     int
	FieldType string
	DefValue  string
}
type XFile interface {
	GetFileTitles(filePath string) ([]Title, error)
	GetDataByIndex(filePath string, tbfInfo []TableFileInfo, rowFun RowFun) error
	GetSheetNames(filePath string) ([]Title, error)
	GetFileTitlesBySheetIndex(filePath string, index int) ([]Title, error)
}

type Xlsx struct {
}

func (this *Xlsx) GetSheetNames(filePath string) ([]Title, error) {
	xl, err := this.open(filePath)
	if err != nil {
		return nil, err
	}
	defer xl.Close()

	t := make([]Title, 0)

	for i, v := range xl.SheetNames() {
		t = append(t, Title{v, i})
	}
	return t, nil
}

func (this *Xlsx) GetFileTitlesBySheetIndex(filePath string, index int) ([]Title, error) {
	xl, err := this.open(filePath)
	if err != nil {
		return nil, err
	}
	defer xl.Close()

	t := make([]Title, 0)

	st := xl.Sheet(index)

	ci := st.Cols()
	for {
		if ci.HasNext() {
			i, co := ci.Next()
			cv := co.Cell(0).String()
			if len(cv) > 0 {
				t = append(t, Title{cv, i})
			}

		} else {
			break
		}
	}

	return t, nil
}

type RowFun func(int, map[int]string) error

func NewFile() XFile {
	return &Xlsx{}
}
func (this *Xlsx) open(filePath string) (*xlsx.Spreadsheet, error) {
	filePath, err := url.QueryUnescape(filePath)
	if err != nil {
		return nil, err
	}
	filePath = strings.Replace(filePath, "file://", "", -1)

	xl, err := xlsx.Open(filePath)
	if err != nil {
		return nil, err
	}
	return xl, nil
}
func (this *Xlsx) GetDataByIndex(filePath string, tbfInfo []TableFileInfo, rowFun RowFun) error {

	xl, err := this.open(filePath)
	if err != nil {
		return err
	}
	defer xl.Close()
	st := xl.Sheet(0)

	finish := false
	ri := st.Rows()

	for {
		if ri.HasNext() {
			i, row := ri.Next()
			if i == 0 {
				continue
			}
			mp := make(map[int]string)
			for i, v := range tbfInfo {
				vv := row.Cell(v.Index).String()

				if len(v.DefValue) > 0 {
					vv = v.DefValue
				}

				if i == 0 && len(vv) == 0 {
					finish = true
					break
				}

				sf := strings.ToLower(v.FieldType)

				t, er := row.Cell(v.Index).Date()
				if er != nil {
					sf = ""
				}
				if len(v.DefValue) > 0 {
					sf = ""
				}
				switch sf {
				case "datetime":
				case "timestamp":
					vv = t.Format("2006-01-02 15:04:05")
					break
				case "date":
					vv = t.Format("2006-01-02")
					break
				case "time":
					vv = t.Format("15:04:05")
					break
				case "year":
					vv = t.Format("2006")
					break
				}

				mp[v.Index] = vv
			}

			if finish {
				break
			}
			err := rowFun(i, mp)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil

}
func (this *Xlsx) GetFileTitles(filePath string) ([]Title, error) {

	xl, err := this.open(filePath)
	if err != nil {
		return nil, err
	}
	defer xl.Close()

	t := make([]Title, 0)

	st := xl.Sheet(0)

	ci := st.Cols()
	for {
		if ci.HasNext() {
			i, co := ci.Next()
			cv := co.Cell(0).String()
			if len(cv) > 0 {
				t = append(t, Title{cv, i})
			}

		} else {
			break
		}
	}

	return t, nil
}
