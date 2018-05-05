package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/fatih/structs"
	"github.com/playwolf719/test/mystruct"
	"github.com/playwolf719/test/utils"
	"log"
	"os"
	"strings"
	"time"
)

type SMController struct {
	beego.Controller
}

type SMRes struct {
	TimeDiff string
	//ResDict  map[string]string
	ResList mystruct.MidList
}

//func SMResInit(smres *SMRes, time_diff string) {
//	smres.time_diff = time_diff
//	smres.res_dict = make(map[string]string)
//}

var root = mystruct.MakeTriesTreeNode("root", make(map[string]mystruct.TriesTreeNode))
var carMap = make(map[string]int)

func (this *SMController) Get() {
	t1 := time.Now()
	query := this.GetString("query")
	initSM()
	final_res, midList := mystruct.FindContent(query, root)
	t2 := time.Now()
	diff := t2.Sub(t1)
	time_diff := float64(diff.Nanoseconds()) / 1000.0 / 1000.0
	logs.Info(final_res)
	carList := mystruct.MidList{}
	nonCarList := mystruct.MidList{}
	//logs.Info(carMap)
	for _, val := range midList {
		_, ok := carMap[val]
		//logs.Info("[dfad]%+v,%+v", val, ok)
		if ok == true {
			carList.Append(val)
		} else {
			nonCarList.Append(val)
		}
	}
	finalList := mystruct.MidList{}
	finalList.AppendList(carList)
	finalList.AppendList(nonCarList)
	smres := &SMRes{
		TimeDiff: fmt.Sprintf("%.6f ms", time_diff),
		//ResDict:  final_res,
		ResList: finalList,
		//ResList: midList,
	}
	logs.Info(smres)
	m := structs.Map(smres)

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
	if len(carMap) == 0 {
		loadCarMap()
	}
}

func loadFileToTree(rel_path string) {
	file, err := os.Open(rel_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
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

func loadCarMap() {
	file, err := os.Open("./static/others/THUOCL_car.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	for scanner.Scan() {
		// do your stuff
		tmp := scanner.Text()
		tmp1 := strings.Fields(tmp)
		carMap[tmp1[0]] = 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
