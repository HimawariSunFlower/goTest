package test

import (
	r "crypto/rand"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math/big"
	"strconv"
	"strings"
	"sync"
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

	row1:=map[string]string{
		"A1": "id",
		"B1": "普通",
		"C1":"优秀",
		"D1":"精良",
		"E1":"史诗",
		"F1":"传说",
		"G1":"神话",
	}
	for k,v:=range row1{
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
			//todo日志记录
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
