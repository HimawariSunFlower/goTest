package test

import (
	"bytes"
	r "crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/tealeg/xlsx"
)

//test 抽奖概率
func TestAddId() {
	//todo 优化
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}

	row1 := map[string]string{
		"A1": "id",
		"B1": "普通",
		"C1": "优秀",
		"D1": "精良",
		"E1": "史诗",
		"F1": "传说",
		"G1": "神话",
	}
	for k, v := range row1 {
		if err := streamWriter.SetRow(k, []interface{}{excelize.Cell{StyleID: styleID, Value: v}}); err != nil {
			fmt.Println(err)
			return
		}
	}

	itemRewards := fmt.Sprintf("%d:%d|", 1, 250)
	itemRewards += fmt.Sprintf("%d:%d|", 2, 270)
	itemRewards += fmt.Sprintf("%d:%d|", 3, 230)
	itemRewards += fmt.Sprintf("%d:%d|", 4, 150)
	itemRewards += fmt.Sprintf("%d:%d|", 5, 80)
	itemRewards += fmt.Sprintf("%d:%d|", 6, 20)

	taskSendRande := NewAttrRandId(1, itemRewards)
	z, x, c, v, b, n := 0, 0, 0, 0, 0, 0
	for rowID := 2; rowID <= 10000; rowID++ {
		row := make([]interface{}, 50)
		ll := taskSendRande.GenItemsCountTask(6)
		q, w, e, r, t, y := 0, 0, 0, 0, 0, 0
		for _, vv := range ll {
			switch vv {
			case 1:
				q++
				z++
			case 2:
				w++
				x++
			case 3:
				e++
				c++
			case 4:
				r++
				v++
			case 5:
				t++
				b++
			case 6:
				y++
				n++
			}
		}
		for colID := 0; colID < 7; colID++ {
			switch colID {
			case 0:
				row[colID] = rowID
			case 1:
				row[colID] = q
			case 2:
				row[colID] = w
			case 3:
				row[colID] = e
			case 4:
				row[colID] = r
			case 5:
				row[colID] = t
			case 6:
				row[colID] = y
			}
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	row := make([]interface{}, 50)
	for colID := 0; colID < 7; colID++ {
		switch colID {
		case 0:
			row[colID] = 10001
		case 1:
			row[colID] = z
		case 2:
			row[colID] = x
		case 3:
			row[colID] = c
		case 4:
			row[colID] = v
		case 5:
			row[colID] = b
		case 6:
			row[colID] = n
		}
	}
	cell, _ := excelize.CoordinatesToCellName(1, 10001)
	if err := streamWriter.SetRow(cell, row); err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
	}
}

type AttrRandId struct {
	sync.RWMutex
	mold int //1权重抽，2概率
	data map[int]int
}

func NewAttrRandId(mold int, source string) *AttrRandId {
	a := new(AttrRandId)
	a.mold = mold
	a.data = make(map[int]int, 0)
	if source == "" {
		return nil
	}
	a.Unmarshal(source)
	return a
}

func (a *AttrRandId) Unmarshal(source string) error {
	a.Lock()
	defer a.Unlock()
	//id:500|id:300|id:200

	item := strings.Split(source, "|")
	a.data = make(map[int]int, 0)
	for _, v := range item {
		line := strings.Split(v, ":")
		if len(line) != 2 {
			//todo error
		} else {
			id, _ := strconv.Atoi(line[0])
			rate, _ := strconv.Atoi(line[1])
			a.data[id] = rate
		}
	}

	return nil
}

//6,5等级 单独抽奖,单次概率,4级有数量上限
const (
	hyperSix  = 6
	hyperFive = 5
	hyperFour = 4
	purpleMax = 3
)

//4,5,6高级任务,每次刷新最多一个
func (a *AttrRandId) GenItemsCountTask(count int) []int {
	a.RLock()
	defer a.RUnlock()
	ret := make([]int, 0)
	if a.mold != 1 {
		return ret
	}
	//6,5等级 单独抽奖,单次概率
	six := a.GenItems()
	if six[0] == hyperSix {
		ret = append(ret, six[0])
		count--
	}
	five := a.GenItems()
	if five[0] == hyperFive {
		ret = append(ret, five[0])
		count--
	}
	four := 0
	for i := 0; i < count; i++ {
		mid := a.GenItems()
		if mid[0] == hyperSix || mid[0] == hyperFive {
			i--
			continue
		}
		if mid[0] == hyperFour {
			if four >= purpleMax {
				i--
				continue
			}
			four++
		}
		ret = append(ret, mid...)
	}
	return ret
}

func (a *AttrRandId) GenItems() []int {
	a.RLock()
	defer a.RUnlock()
	//log.Debugf("attrReward:%d,%s", a.mold, a.String())

	ret := make([]int, 0)

	if a.mold == 1 {
		//权重抽,只出一个
		rateNum := 0
		rnum := RandomN(1000)

		for id, v := range a.data { //依次 200，500，200，100
			rateNum += v //200,700,900,1000
			if rnum <= rateNum {
				ret = append(ret, id)
				return ret
			}
		}
	} else if a.mold == 2 { //每个物品计算概率
		for kid, vnum := range a.data {
			rnum := RandomN(1000)
			//fmt.Println(rnum, v.rate, v.Id)
			if rnum < vnum {
				ret = append(ret, kid)
			}
		}
	}
	//log.Debugf("attrReward ret,%+v", ret)
	return ret
}

func RandomN(n int) int {
	ret, _ := r.Int(r.Reader, big.NewInt(int64(n)))
	return int(ret.Int64())
}

func ExcelToCsv() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, _ := ioutil.ReadDir("./table/")
	for _, f := range files {
		file := readExcelFile(path + "/table/" + f.Name())
		err = parseSheetToCSV(file.Sheets[0], path+"/tablecsv/"+strings.ToLower(file.Sheets[0].Name)+".csv")
		if err != nil {
			panic(err)
		}
	}
}

func readExcelFile(path string) (f *xlsx.File) {
	f, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("excel文件读取错误")
		panic(err)
	}
	return
}

func parseSheetToCSV(sheet *xlsx.Sheet, toFile string) (err error) {
	b := &bytes.Buffer{}
	rows := sheet.MaxRow
	for i := 0; i < rows; i++ {
		cols := sheet.MaxCol
		for j := 0; j < cols; j++ {
			cell := sheet.Cell(i, j)
			val := cell.Value
			fmt.Println(val)
			b.WriteString(val)
			b.WriteString("\t")
		}
		b.WriteString("\r\n")
	}
	//写入数据到文件
	err = ioutil.WriteFile(toFile, b.Bytes(), os.ModePerm)
	return
}

type ZhiGou struct {
	Id    int
	Price int
}

var ZhiGous = []string{"id", "Price"}

type meta struct {
	Key string
	Idx int
	Typ string
}

type rowdata []interface{}

func ExcelToJson() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, _ := ioutil.ReadDir("./table/")
	for _, f := range files {
		parseFile(path + "/table/" + f.Name())
	}
}

func parseFile(file string) {

	fmt.Println("\n\n\n\n", file)

	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		panic(err.Error())
	}
	//[line][colidx][data]

	sheets := xlsx.GetSheetList()
	for _, s := range sheets {
		rows, err := xlsx.GetRows(s)
		if err != nil {
			return
		}
		if len(rows) < 5 {
			return
		}

		colNum := len(rows[1])
		fmt.Println("col num:", colNum)
		metaList := make([]*meta, 0, colNum)
		dataList := make([]rowdata, 0, len(rows)-4)

		for line, row := range rows {
			switch line {
			case 0: // sheet 名

			case 1: // col name
				//fmt.Println("meta cot:%d, rol cot:%d", len(metaList), len(row))
				// for idx, typ := range row {
				// 	metaList[idx].Typ = typ
				// }
			case 2: // data type
				for idx, colname := range row {
					fmt.Println(idx, colname, len(metaList))
					for _, v := range ZhiGous {
						if v == colname {
							metaList = append(metaList, &meta{Key: colname, Idx: idx, Typ: rows[1][idx]})
						}
					}
				}

			default: //>= 4 row data
				data := make(rowdata, 0, colNum)

				for k := 0; k < colNum; k++ {
					for _, v := range metaList {
						if v.Idx == k {
							if k < len(row) {
								data = append(data, row[k])
							}
						}
					}
				}

				dataList = append(dataList, data)
			}
		}

		//sheetName := xlsx.GetSheetName(idx)
		// to json, save
		filename := filepath.Base(file)
		suf := filepath.Ext(filename)
		jsonFile := fmt.Sprintf("%s.json", filename[:(len(filename)-len(suf))])
		err = output(jsonFile, toJson(dataList, metaList))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(toJson(dataList, metaList))
		return
	}

}

func toJson(datarows []rowdata, metalist []*meta) string {
	ret := "["

	for _, row := range datarows {
		ret += "\n\t{"
		for idx, meta := range metalist {
			ret += fmt.Sprintf("\n\t\t\"%s\":", meta.Key)
			if meta.Typ == "string" {
				if row[idx] == nil {
					ret += "\"\""
				} else {
					ret += fmt.Sprintf("\"%s\"", row[idx])
				}
			} else {
				if row[idx] == nil || row[idx] == "" {
					ret += "0"
				} else {
					ret += fmt.Sprintf("%s", row[idx])
				}
			}
			ret += ","
		}
		ret = ret[:len(ret)-1]

		ret += "\n\t},"
	}
	ret = ret[:len(ret)-1]

	ret += "\n]"
	return ret
}

func output(filename string, str string) error {
	path, _ := os.Getwd()
	f, err := os.OpenFile(path+"/tablecsv/"+filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}
