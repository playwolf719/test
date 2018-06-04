package mystruct

import (
	"fmt"
	"strings"
)

type WordNode struct {
	Word  string
	Score int
}

type TriesTreeNode struct {
	One_word string
	Tmap     map[string]*TriesTreeNode
	WordList *NodeList
}

type NodeList []WordNode

func (p *NodeList) Append(node WordNode) {
	*p = append(*p, node)
}

func TriesTreeInit(node *TriesTreeNode) {
	node.Tmap = make(map[string]*TriesTreeNode)
}

func SayWorld() {
	fmt.Println("SayWorld")
}

func MakeTriesTreeNode(one_word string, Tmap map[string]*TriesTreeNode, WordList *NodeList) TriesTreeNode {
	return TriesTreeNode{one_word, Tmap, WordList}
}

func InsertContent(content string, strScore int, root TriesTreeNode) {
	sarr := strings.Split(content, "")
	subInsertContent(sarr, 0, root, strScore)
}

func subInsertContent(sarr []string, index int, node TriesTreeNode, strScore int) {
	if index >= len(sarr) {
		return
	}
	word := sarr[index]
	_, ok := node.Tmap[word]
	if ok != true {
		nodeList := NodeList{WordNode{strings.Join(sarr, ""), strScore}}
		tmp := MakeTriesTreeNode(word, make(map[string]*TriesTreeNode), &nodeList)
		node.Tmap[word] = &tmp
	} else {
		insertNode(node.Tmap[word], sarr, strScore)
	}
	MyHeapSort(node.Tmap[word].WordList)
	subInsertContent(sarr, index+1, *node.Tmap[word], strScore)
}

func insertNode(node *TriesTreeNode, sarr []string, strScore int) {
	if listLen := len(*node.WordList); listLen > 0 {
		if listLen >= 20 {
			tmpList := *node.WordList
			tmpList[listLen-1] = WordNode{strings.Join(sarr, ""), strScore}
		} else {
			node.WordList.Append(WordNode{strings.Join(sarr, ""), strScore})
		}
	}
}

type MidList []string

func (p *MidList) Append(data string) {
	*p = append(*p, data)
}
func (p *MidList) AppendList(data []string) {
	*p = append(*p, data...)
}
func FindContent(content string, root TriesTreeNode) (final_res map[string]string, final_list MidList) {
	sarr := strings.Split(content, "")
	final_res = make(map[string]string)
	final_list = MidList{}
	subFindContent(sarr, 0, root, final_res, &final_list)
	//logs.Info("[list]%+v", final_list)
	return final_res, final_list
}

func FindContentFast(content string, root TriesTreeNode) (nodeList NodeList) {
	sarr := strings.Split(content, "")
	index := 0
	theNode := root
	nodeList = NodeList{}
	for index < len(sarr) {
		theStr := sarr[index]
		theNode1, ok := theNode.Tmap[theStr]
		if index == len(sarr)-1 && ok {
			nodeList = *theNode1.WordList
			break
		} else if ok {
			index = index + 1
			theNode = *theNode1
			continue
		} else {
			break
		}

	}
	return nodeList
}

func subFindContent(sarr []string, index int, node TriesTreeNode, final_res map[string]string, final_list *MidList) {
	if index >= len(sarr) {
		return
	}
	word := sarr[index]
	_, ok := node.Tmap[word]
	if ok == true {
		if index == len(sarr)-1 {
			getAllRes(*node.Tmap[word], strings.Join(sarr[0:index+1], ""), final_res, final_list)
		} else {
			subFindContent(sarr, index+1, *node.Tmap[word], final_res, final_list)
		}
	} else {
		//if index > 0 {
		//	getAllRes(node, strings.Join(sarr[0:index], ""), final_res)
		//}
	}
}

func getAllRes(node TriesTreeNode, prefix string, final_res map[string]string, final_list *MidList) {
	if len(node.Tmap) == 0 {
		final_res[prefix] = prefix
		//final_list = append(final_list, prefix)
		final_list.Append(prefix)
	} else {
		for k, v := range node.Tmap {
			if len(v.Tmap) == 0 {
				tmp := prefix + k
				if len(tmp) > 0 {
					//final_res[tmp] = tmp
					final_list.Append(tmp)
				}
			} else {
				getAllRes(*v, prefix+k, final_res, final_list)
			}
		}
	}
}
