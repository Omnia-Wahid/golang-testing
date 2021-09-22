package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	//Init
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	//go routine
	go func() {
		//Execution
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "bubble sort took more than 50 ms")
		return
	}
	//Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])

}
func TestSortIncreasingOrder(t *testing.T) {
	//Init
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := GetElements(10)
	//Execution
	Sort(elements)

	//Validation
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := GetElements(10)
	//fmt.Println(elements)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

//func TestBubbleSortAlreadySorted(t *testing.T) {
//	//Init
//	elements := []int{5,4,3,2,1}
//
//	fmt.Println(elements)
//	//Execution
//	BubbleSort(elements)
//}
