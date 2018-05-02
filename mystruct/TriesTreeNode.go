package mystruct

import (
	"fmt"
	"strings"
)

type TriesTreeNode struct {
	One_word string
	Tmap     map[string]TriesTreeNode
}

func TriesTreeInit(node *TriesTreeNode) {
	node.Tmap = make(map[string]TriesTreeNode)
}

func SayWorld() {
	fmt.Println("SayWorld")
}

func MakeTriesTreeNode(one_word string, Tmap map[string]TriesTreeNode) TriesTreeNode {
	return TriesTreeNode{one_word, Tmap}
}

func InsertContent(content string, root TriesTreeNode) {
	sarr := strings.Split(content, "")
	subInsertContent(sarr, 0, root)
}

func subInsertContent(sarr []string, index int, node TriesTreeNode) {
	if index >= len(sarr) {
		return
	}
	word := sarr[index]
	_, ok := node.Tmap[word]
	if ok != true {
		node.Tmap[word] = MakeTriesTreeNode(word, make(map[string]TriesTreeNode))
	}
	subInsertContent(sarr, index+1, node.Tmap[word])
}

func FindContent(content string, root TriesTreeNode) map[string]string {
	sarr := strings.Split(content, "")
	final_res := make(map[string]string)
	subFindContent(sarr, 0, root, final_res)
	return final_res
}

func subFindContent(sarr []string, index int, node TriesTreeNode, final_res map[string]string) {
	if index >= len(sarr) {
		return
	}
	word := sarr[index]
	_, ok := node.Tmap[word]
	if ok == true {
		if index == len(sarr)-1 {
			getAllRes(node.Tmap[word], strings.Join(sarr[0:index+1], ""), final_res)
		} else {
			subFindContent(sarr, index+1, node.Tmap[word], final_res)
		}
	} else {
		if index > 0 {
			getAllRes(node, strings.Join(sarr[0:index], ""), final_res)
		}
	}
}

func getAllRes(node TriesTreeNode, prefix string, final_res map[string]string) {
	for k, v := range node.Tmap {
		if len(v.Tmap) == 0 {
			tmp := prefix + k
			if len(tmp) > 0 {
				final_res[tmp] = tmp
			}
		} else {
			getAllRes(v, prefix+k, final_res)
		}
	}
}
