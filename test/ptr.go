package test

import "fmt"

type EquipBar struct {
	Data [4]*Equip
}

type Equip struct {
	Id   int
	name string
}

type FEquipBar struct {
	Data []*Equip
}

func TestPtr1() {
	pe := &EquipBar{}
	r := &FEquipBar{
		Data: pe.Data[:],
	}

	if r.Data == nil {
		fmt.Println(11)
	}
}
