package mystruct

func swap(arrList *NodeList, index1 int, index2 int) {
	tmp1 := *arrList
	tmp1[index1], tmp1[index2] = tmp1[index2], tmp1[index1]
}

func subHeapSort(arrList *NodeList, start int, end int) {
	tmp := *arrList
	//arrLen := len(tmp)
	for start <= end {
		parent := start
		lSon := parent*2 + 1
		rSon := parent*2 + 2
		indexMax := parent
		if lSon <= end && tmp[indexMax].Score > tmp[lSon].Score {
			indexMax = lSon
		}
		if rSon <= end && tmp[indexMax].Score > tmp[rSon].Score {
			indexMax = rSon
		}
		if parent == indexMax {
			break
		}
		swap(&tmp, parent, indexMax)
		start = indexMax
	}
}

func buildHeap(arrList *NodeList) {
	arrLen := len(*arrList)
	theIndex := int(arrLen - 1/2)
	for theIndex >= 0 {
		subHeapSort(arrList, theIndex, arrLen-1)
		theIndex = theIndex - 1
	}
}

func MyHeapSort(arrList *NodeList) {
	arrLen := len(*arrList)
	buildHeap(arrList)
	for index := arrLen - 1; index >= 1; index-- {
		swap(arrList, 0, index)
		subHeapSort(arrList, 0, index-1)
	}

}
