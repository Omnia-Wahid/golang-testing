package services

import "github.com/omniaw/golang-testing/src/api/utils/sort"

const (
	privateConst = "private"
	PublicConst  = "public"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
