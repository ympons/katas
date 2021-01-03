package leetcode

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	units := 0
	for _, box := range boxTypes {
		for box[0] > 0 && truckSize > 0 {
			units += box[1]
			truckSize--
			box[0]--
		}
		if truckSize <= 0 {
			return units
		}
	}
	return units
}
