package main

import (
				"text/scanner"
				"strings"
				"fmt"
				"os"
				"flag"
				"io"
				"io/ioutil"
				"bufio"
				//"time"
)

func main() {
				//test()
				main1()
}

func test() {

				const src = `
	// This is scanned code.
	if a > 10 {
		someParsable = text
	}`
				var s scanner.Scanner
				s.Filename = "example"
				s.Init(strings.NewReader(src))
				fmt.Println(s)
				var tok rune
				for tok != scanner.EOF {
								tok = s.Scan()
								fmt.Println("At position", s.Pos(), ":", s.TokenText())
				}

}

func read1(path string) string {
				fi, err := os.Open(path)
				if err != nil {
								panic(err)
				}
				defer fi.Close()

				chunks := make([]byte, 1024, 1024)
				buf := make([]byte, 1024)
				for {
								n, err := fi.Read(buf)
								if err != nil && err != io.EOF {
												panic(err)
								}
								if 0 == n {
												break
								}
								chunks = append(chunks, buf[:n]...)
								// fmt.Println(string(buf[:n]))
				}
				return string(chunks)
}

func read2(path string) string {
				fi, err := os.Open(path)
				if err != nil {
								panic(err)
				}
				defer fi.Close()
				r := bufio.NewReader(fi)

				chunks := make([]byte, 1024, 1024)

				buf := make([]byte, 1024)
				for {
								n, err := r.Read(buf)
								if err != nil && err != io.EOF {
												panic(err)
								}
								if 0 == n {
												break
								}
								chunks = append(chunks, buf[:n]...)
								// fmt.Println(string(buf[:n]))
				}
				return string(chunks)
}

func read3(path string) string {
				fi, err := os.Open(path)
				if err != nil {
								panic(err)
				}
				defer fi.Close()
				fd, err := ioutil.ReadAll(fi)
				// fmt.Println(string(fd))
				return string(fd)
}

func main1() {
				var map_names = make(map[string] string)
				map_names["fileld1"] = "机构名称"
				map_names["fileld2"] = "病人档案ID"
				map_names["fileld3"] = "检测时间"
				map_names["fileld4"] = "样本来源"
				map_names["fileld5"] = "婴儿信息"
				map_names["fileld6"] = "母亲信息"
				map_names["fileld7"] = "姓名"
				map_names["fileld8"] = "性别"
				map_names["fileld9"] = "年龄(岁)"
				map_names["fileld10"] = "身高(cm)"
				map_names["fileld11"] = "体重(KG)"
				map_names["fileld12"] = "头围(cm)"
				map_names["fileld13"] = "是否早产"
				map_names["fileld14"] = "BMI"
				map_names["fileld15"] = "分娩方式"

				map_names["fileld16"] = "情绪"
				map_names["fileld17"] = "Apgar评分"
				map_names["fileld18"] = "胎次"
				map_names["fileld19"] = "住院号"
				map_names["fileld20"] = "出生日期"
				map_names["fileld21"] = "门诊号"

				map_names["fileld22"] = "床号:"
				map_names["fileld23"] = "送检科室"
				map_names["fileld24"] = "送检医生"
				map_names["fileld25"] = "手机号"
				map_names["FAT"] = "FAT"
				map_names["SNF"] = "SNF"
				map_names["Density"] = "Density"

				map_names["Protein"] = "Protein"
				map_names["Lactose"] = "Lactose"
				map_names["Minerals"] = "Minerals"
				map_names["Freezing"] = "Freezing"
				map_names["Energy"] = "Energy"
				map_names["waterContent"] = "waterContent"

				map_names["Carbohydrate"] = "Carbohydrate:"
				map_names["grayscale"] = "grayscale"
				map_names["result"] = "诊断结果"

				map_names["fileld38"] = "报告者"
				map_names["fileld39"] = "管理员审核者"
				map_names["datetime"] = "日期"

				flag.Parse()
				//file := flag.Arg(0)
				filepath := "/Users/doit/tmp/20160716143816001.txt"
				content, err := ioutil.ReadFile(filepath)
				if err != nil {
								fmt.Printf("%s\n", err)
								panic(err)
				}
				lines := strings.Split(string(content), "\n")
				j := 0;
				if err == nil {
								for _, line := range lines {
												if len(line) > 0 {
																j ++
																line = strings.TrimSpace(line)
																switch  {
																				case strings.Contains(line, "病人档案ID"):
																				pid, datetime := extract1(line)
																				fmt.Println("pid : ", pid , "datetime : ", datetime)
																}
												}

								}
				}

}

func extract1(line string)(pid, datetime string)  {
								tmp := strings.Split(line, " ")
								if len(tmp) > 0 {
												//fmt.Printf("%v  %d \n", tmp, len(tmp))
												i := 0
												date := make([]string, 2)
												for _, str := range tmp {
																if len(str) > 0 {
																				i ++
																				switch i {
																				case 1:
																								fmt.Println("病人档案ID")
																				case 2:
																								pid = str
																				case 3:
																								date_arr := strings.Split(str, ":")
																								date = append(date, date_arr[1])
																								fmt.Println(date_arr[0])

																				case 4:

																								date = append(date, str)

																								datetime = strings.Join(date, " ")
																				}


																}


												}
								}
								return
				
}

func extract3(line string)(source, sample string)  {
				tmp := strings.Split(line, " ")
				if len(tmp) > 0 {
								//fmt.Printf("%v  %d \n", tmp, len(tmp))
								i := 0
								date := make([]string, 2)
								for _, str := range tmp {
												if len(str) > 0 {
																i ++
																switch i {
																case 1:
																				fmt.Println("病人档案ID")
																case 2:
																				pid = str
																case 3:
																				date_arr := strings.Split(str, ":")
																				date = append(date, date_arr[1])
																				fmt.Println(date_arr[0])

																case 4:

																				date = append(date, str)

																				datetime = strings.Join(date, " ")
																}


												}


								}
				}
				return

}
