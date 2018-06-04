package main

import (
	"fmt"
	//"index/suffixarray"
)

type itself []string

func (h *itself) appendToItself(test string) {
	*h = append(*h, test)
}

type IntList []int

func (p *IntList) Swap1(index1 int, index2 int) {
	fu1 := *p
	tmp := fu1[index1]
	fu1[index1] = fu1[index2]
	fu1[index2] = tmp
}

func swap(arrList *[]int, index1 int, index2 int) {
	tmp1 := *arrList
	tmp1[index1], tmp1[index2] = tmp1[index2], tmp1[index1]
}

//func (p *IntList) Append(data int) {
//	*p = append(*p, data)
//}

func main() {
	//h := itself{"1", "2"}
	//h := []string{}
	//h := make([]string, 0, 10)
	//logs.Info(cap(h))
	//h.appendToItself("3")
	//test(&h)
	arrList := []int{4, 3, 2, 7, 1, 5, 6, 0}
	fmt.Println(arrList)
	heapSort(&arrList)
	//swap(&arrList, 1s, 3)
	//subHeapSort(&arrList, 3)
	//swap(arrList, 1, 2)
	fmt.Println(arrList)
	arrList[0] = 5
	fmt.Println(arrList)
	heapSort(&arrList)
	//subHeapSort(&arrList, 0, len(arrList)-1)
	fmt.Println(arrList)
}

func subHeapSort(arrList *[]int, start int, end int) {
	tmp := *arrList
	//arrLen := len(tmp)
	for start <= end {
		parent := start
		lSon := parent*2 + 1
		rSon := parent*2 + 2
		indexMax := parent
		if lSon <= end && tmp[indexMax] < tmp[lSon] {
			indexMax = lSon
		}
		if rSon <= end && tmp[indexMax] < tmp[rSon] {
			indexMax = rSon
		}
		if parent == indexMax {
			break
		}
		swap(&tmp, parent, indexMax)
		start = indexMax
	}
}

func buildHeap(arrList *[]int) {
	arrLen := len(*arrList)
	theIndex := int(arrLen - 1/2)
	for theIndex >= 0 {
		subHeapSort(arrList, theIndex, arrLen-1)
		theIndex = theIndex - 1
	}
}

func heapSort(arrList *[]int) {
	arrLen := len(*arrList)
	buildHeap(arrList)
	for index := arrLen - 1; index >= 1; index-- {
		swap(arrList, 0, index)
		subHeapSort(arrList, 0, index-1)
	}

}

func test(h *itself) {
	h.appendToItself("44")
	fmt.Println(h)
}
