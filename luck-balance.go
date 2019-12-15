package main

import "sort"

type ImportantContests []int32

func (a ImportantContests) Len() int           { return len(a) }
func (a ImportantContests) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ImportantContests) Less(i, j int) bool { return a[i] < a[j] }

func luckBalance(k int32, contests [][]int32) int32 {
	var luck int32
	var numImpContests int32
	importantContests := make(ImportantContests, 0)

	for _, value := range contests {
		if value[1] == 0 {
			luck += value[0]
		} else {
			numImpContests++
			importantContests = append(importantContests, value[0])
		}
	}

	// k = number of allowed important losses. k may be negative
	if k > 0 {
		// sort array in reverse order
		sort.Sort(sort.Reverse(importantContests))
		for index, value := range importantContests {
			if int32(index+1) <= k {
				luck += value
			} else {
				luck -= value
			}
		}
	} else {
		for _, value := range importantContests {
			luck -= value
		}
	}
	return luck
}

func main() {}
