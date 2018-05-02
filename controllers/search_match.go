package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fatih/structs"
	"github.com/playwolf719/test/mystruct"
	"github.com/playwolf719/test/utils"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type SMController struct {
	beego.Controller
}

type SMRes struct {
	Time_diff string
	Res_dict  map[string]string
}

//func SMResInit(smres *SMRes, time_diff string) {
//	smres.time_diff = time_diff
//	smres.res_dict = make(map[string]string)
//}

var root = mystruct.MakeTriesTreeNode("root", make(map[string]mystruct.TriesTreeNode))

func (this *SMController) Get() {
	t1 := time.Now()
	query := this.GetString("query")
	initSM()
	final_res := mystruct.FindContent(query, root)
	//json_res
	t2 := time.Now()

	fmt.Println(t1)
	fmt.Println(t2)
	diff := t2.Sub(t1)

	//smres := new(SMRes)
	//smres.time_diff = strconv.Itoa(int(diff.Nanoseconds()))
	//smres.res_dict = final_res

	smres := &SMRes{
		Time_diff: strconv.Itoa(int(diff.Nanoseconds())/1000 )+"ms",
		Res_dict:  final_res,
	}

	//tmp := utils.Struct2Map(smres)
	//log.Println(tmp)
	//s := &Server{
	//	Name:    "gopher",
	//	ID:      123456,
	//	Enabled: true,
	//}

	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	log.Println(smres)
	m := structs.Map(smres)

	log.Println(m)
	this.Data["json"] = &m
	this.ServeJSON()
}

func initSM() {
	if len(root.Tmap) == 0 {
		tmp := utils.GetFileMap("./static/others/")
		for _, val := range tmp {
			if strings.Contains(val, "txt") {
				loadFileToTree(val)
			}
		}
		log.Println("初始化搜索树")
	}
}

func loadFileToTree(rel_path string) {
	log.Println(rel_path)
	file, err := os.Open(rel_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	tmp := scanner.Text()
	//	tmp1 := strings.Split(tmp, " ")
	//	mystruct.InsertContent(tmp1[0], root)
	//}

	//buf := make([]byte, 0, 64*1024)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	for scanner.Scan() {
		// do your stuff
		tmp := scanner.Text()
		tmp1 := strings.Fields(tmp)
		mystruct.InsertContent(tmp1[0], root)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
