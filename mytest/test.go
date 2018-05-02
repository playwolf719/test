package main

import (
	"bufio"
	"fmt"
	"github.com/playwolf719/test/mystruct"
	"log"
	"os"
	"strings"
)

var root = mystruct.MakeTriesTreeNode("root", make(map[string]mystruct.TriesTreeNode))

func main() {
	file, err := os.Open("./THUOCL_chengyu.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		tmp1 := strings.Split(tmp, " ")
		mystruct.InsertContent(tmp1[0], root)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mystruct.FindContent("你", root))
}
