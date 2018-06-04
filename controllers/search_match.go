package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/fatih/structs"
	"github.com/playwolf719/test/mystruct"
	"github.com/playwolf719/test/utils"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SMController struct {
	beego.Controller
}

type SMRes struct {
	TimeDiff     string
	TimeDiffFast string
	//ResDict  map[string]string
	ResList  mystruct.MidList
	NodeList mystruct.NodeList
}

//func SMResInit(smres *SMRes, time_diff string) {
//	smres.time_diff = time_diff
//	smres.res_dict = make(map[string]string)
//}

var root = mystruct.MakeTriesTreeNode("root", make(map[string]*mystruct.TriesTreeNode), &mystruct.NodeList{})
var carMap = make(map[string]int)

func (this *SMController) Get() {
	query := this.GetString("query")
	initSM()
	t1 := time.Now()
	final_res, midList := mystruct.FindContent(query, root)
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
	sort.Strings(carList)
	sort.Strings(nonCarList)
	finalList := mystruct.MidList{}
	finalList.AppendList(carList)
	finalList.AppendList(nonCarList)
	t2 := time.Now()
	diff := t2.Sub(t1)
	time_diff := float64(diff.Nanoseconds()) / 1000.0 / 1000.0
	logs.Info("[FindContent] %+v %+v]", final_res, time_diff)

	t3 := time.Now()
	nodeList := mystruct.FindContentFast(query, root)
	t4 := time.Now()
	diffFast := t4.Sub(t3)
	time_diff_fast := float64(diffFast.Nanoseconds()) / 1000.0 / 1000.0
	logs.Info("[FindContentFast] %+v %+v]", nodeList, time_diff_fast)
	smres := &SMRes{
		TimeDiff: fmt.Sprintf("%.6f ms", time_diff),
		//ResDict:  final_res,
		ResList:      finalList,
		TimeDiffFast: fmt.Sprintf("%.6f ms", time_diff_fast),
		NodeList:     nodeList,
		//ResList: midList,
	}
	logs.Info(smres)
	m := structs.Map(smres)

	this.Data["json"] = &m
	this.ServeJSON()
}

func initSM() {
	if len(root.Tmap) == 0 {
		tmp := utils.GetFileMap("./static/others1/")
		for _, val := range tmp {
			if strings.Contains(val, "txt") {
				loadFileToTree(val)
			}
		}
		logs.Info("初始化搜索树")
	}
	if len(carMap) == 0 {
		loadCarMap()
	}
}

func loadFileToTree(rel_path string) {
	file, err := os.Open(rel_path)
	if err != nil {
		logs.Error(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	for scanner.Scan() {
		// do your stuff
		tmp := scanner.Text()
		tmp1 := strings.Fields(tmp)
		if tmp1Len := len(tmp1); tmp1Len != 2 {
			//logs.Error("[len wrong]%+v", tmp)
			continue
		}
		if score, err := strconv.Atoi(tmp1[1]); err == nil {
			mystruct.InsertContent(tmp1[0], score, root)
		} else {
			//logs.Error("[not looks like a number]%+v", tmp)
		}
	}
	if err := scanner.Err(); err != nil {
		logs.Error(err)
	}
}

func loadCarMap() {
	file, err := os.Open("./static/others/THUOCL_car.txt")
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
	}
}
