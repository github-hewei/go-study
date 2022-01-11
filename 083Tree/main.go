package main

import (
	"encoding/json"
	"fmt"
)

// Region 地区数据结构体
type Region struct {
	Id         int             `json:"id"`
	Pid        int             `json:"pid"`
	Name       string          `json:"name"`
	Children   []Region        `json:"children"`
	PointerMap map[int]*Region `json:"-"`
}

// RegionList 地区数据列表
type RegionList []Region

// GetTree 获取树型结构数据
func (r *RegionList) GetTree() []Region {
	listMap := make(map[int]*Region)

	for _, item := range *r {
		tempVar := item
		listMap[item.Id] = &tempVar
	}

	newMap := make(map[int]*Region)
	for _, item := range *r {
		if _, ok := listMap[item.Pid]; ok {
			if listMap[item.Pid].PointerMap == nil {
				listMap[item.Pid].PointerMap = make(map[int]*Region)
			}
			listMap[item.Pid].PointerMap[item.Id] = listMap[item.Id]
		} else if item.Pid == 0 {
			newMap[item.Pid] = listMap[item.Id]
		}
	}

	return r.buildTree(newMap)
}

// buildTree 构建树型结构数据
func (r *RegionList) buildTree(listMap map[int]*Region) []Region {
	list := make([]Region, 0)
	for _, item := range listMap {
		if item.PointerMap != nil {
			item.Children = r.buildTree(item.PointerMap)
		} else {
			item.Children = make([]Region, 0)
		}
		item.PointerMap = nil
		list = append(list, *item)
	}
	return list
}

func main() {

	list := RegionList{
		{Id: 1, Pid: 0, Name: "中国"},
		{Id: 2, Pid: 1, Name: "北京"},
		{Id: 3, Pid: 1, Name: "河北省"},
		{Id: 4, Pid: 2, Name: "北京市"},
		{Id: 5, Pid: 4, Name: "朝阳区"},
		{Id: 6, Pid: 3, Name: "保定市"},
	}

	tree := list.GetTree()

	result, err := json.Marshal(tree)
	if err == nil {
		fmt.Printf("%+v\n", string(result))
	}
}
