package main

import "fmt"

type Region struct {
	Id int
	Name string
	Pid int
	Children []Region
}

type RegionList []Region

func (rt *RegionList) ToTree(pid int) []Region {
	var ret []Region
	for _, item := range *rt {
		if item.Pid == pid {
			item.Children = rt.ToTree(item.Id)
			ret = append(ret, item)
		}
	}
	return ret
}

func main() {

	rt := RegionList{
		{Id: 1, Name: "aa", Pid: 0},
		{Id: 2, Name: "aa", Pid: 1},
		{Id: 3, Name: "aa", Pid: 1},
		{Id: 4, Name: "aa", Pid: 2},
		{Id: 5, Name: "aa", Pid: 4},
	}

	tree := rt.ToTree(0)

	//fmt.Println(tree)
	fmt.Printf("%+v", tree)
}
