/**
* @Auther:gy
* @Date:2021/4/28 9:15
 */

package main

import (
	"fmt"
	"gitee.com/yanggit123/tool"
	"html/template"
	"os"
	"strings"
	"sync"
)

type TableInfo struct {
	StructName      string //结构体
	StructLowerName string
	AppName         string
	Dir             string
	Want            string
}

func parse(info TableInfo) {
	w := []string{"c", "s", "m"}
	if info.Want != "" {
		w = strings.Split(info.Want, ",")
	}
	tpl := []string{"./template/controller.tpl", "./template/service.tpl", "./template/curd.tpl"}
	if info.Dir != "" {
		if strings.Contains(info.Dir, ",") {
			tpl = strings.Split(info.Dir, ",")
		} else {
			if info.Dir[len(info.Dir)-1] == '/' {
				tpl = []string{info.Dir + "controller.tpl", info.Dir + "service.tpl", info.Dir + "curd.tpl"}
			} else {
				tpl = []string{info.Dir + "/controller.tpl", info.Dir + "/service.tpl", info.Dir + "/curd.tpl"}
			}
		}
	}
	if len(tpl) != 3 {
		fmt.Println("请传入正确的路径 example:")
		fmt.Println(" -dir ./template")
		fmt.Println(" -dir ./template/c.tpl,./template/s.tpl,./template/m.tpl")
		return
	}
	var wg sync.WaitGroup
	for _, v := range w {
		if v == "c" {
			wg.Add(1)
			go parseController(tpl[0], info, &wg)
		} else if v == "s" {
			wg.Add(1)
			go parseService(tpl[1], info, &wg)
		} else if v == "m" {
			wg.Add(1)
			go parseModel(tpl[2], info, &wg)
		}
	}
	wg.Wait()
}
func parseModel(path string, info TableInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	os.MkdirAll("./models", os.ModePerm)
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	files, err := os.OpenFile("./models/"+strings.ToLower(info.StructName)+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmpl.Execute(files, info)

}
func parseService(path string, info TableInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	os.MkdirAll("./service", os.ModePerm)
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	files, err := os.OpenFile("./service/"+strings.ToLower(info.StructName)+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmpl.Execute(files, info)

}
func parseController(path string, info TableInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	os.MkdirAll("./endpoint/api/"+info.StructLowerName, os.ModePerm)
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	files, err := os.OpenFile("./endpoint/api/"+info.StructLowerName+"/"+strings.ToLower(info.StructName)+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmpl.Execute(files, info)

}
func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" {
		fmt.Println("Options:")
		fmt.Println("	-a address 指定数据库ip:port")
		fmt.Println("	-u user 指定数据库账号")
		fmt.Println("	-p password 指定数据库密码")
		fmt.Println("	-d database 指定数据库名称")
		fmt.Println("	-P prefix 指定数据库表名前缀")
		fmt.Println("	-s struct 指定生成结构体多个以逗号隔开 a,b")
		fmt.Println("	-n name 指定项目名称 wenshanzhou")
		fmt.Println("	-want c,s,m 指定想要输出的层(默认3层都会输出),以英文逗号隔开")
		fmt.Println("	-dir ./template 指定模板文件路径")
		fmt.Println("	     ./template/controller.tpl,./template/service.tpl,./template/models.tpl 指定模板文件")
		fmt.Println("示例 ./main -a 127.0.0.1:3306 -u root -p 123456 -P t_")
		return
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}

	}()
	conf := tool.MysqlConf{}
	tables := []string{}
	appName := ""
	want := ""
	dir := ""
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-a":
			conf.Address = os.Args[i+1]
			i++
		case "-u":
			conf.Username = os.Args[i+1]
			i++
		case "-p":
			conf.Password = os.Args[i+1]
			i++
		case "-P":
			conf.Prefix = os.Args[i+1]
			i++
		case "-d":
			conf.DbName = os.Args[i+1]
			i++
		case "-s":
			tables = strings.Split(os.Args[i+1], ",")
			i++
		case "-n":
			appName = os.Args[i+1]
			i++
		case "-want":
			want = os.Args[i+1]
			i++
		case "-dir":
			dir = os.Args[i+1]
			i++
		}
	}
	if appName == "" {
		fmt.Println("项目名称为必填项")
		return
	}
	if len(tables) == 0 {
		conf.MaxOpenConns = 64
		conf.MaxIdleConns = 16
		conf.ConnMaxLifetime = 3600000
		db, err := tool.EnableMysql(conf)
		if err != nil {
			fmt.Println("数据库连接出错 err:" + err.Error())
			return
		}
		rows, err := db.Raw("SELECT `TABLE_NAME` AS 'table_name' FROM "+
			"information_schema.tables WHERE `TABLE_SCHEMA` = ?", conf.DbName).Rows()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		for rows.Next() {
			var table string
			rows.Scan(&table)
			table = transferBigCamelCase(strings.Replace(table, conf.Prefix, "", 1))
			tables = append(tables, table)
		}

		for _, v := range tables {
			parse(TableInfo{
				StructName:      v,
				StructLowerName: strings.ToLower(v[:1]) + v[1:],
				AppName:         appName,
				Want:            want,
				Dir:             dir,
			})
		}
	} else {
		for _, v := range tables {
			parse(TableInfo{
				StructName:      v,
				StructLowerName: strings.ToLower(v[:1]) + v[1:],
				AppName:         appName,
				Want:            want,
				Dir:             dir,
			})
		}
	}
}
func transferBigCamelCase(tableName string) string {
	str := []byte{}
	for i := 0; i < len(tableName); i++ {
		if i == 0 {
			if 'a' <= tableName[0] && tableName[0] <= 'z' {
				str = append(str, tableName[0]-32)
			}
		} else if tableName[i] == '_' {
			if 'a' <= tableName[i+1] && tableName[i+1] <= 'z' {
				str = append(str, tableName[i+1]-32)
			}
			i++
		} else {
			str = append(str, tableName[i])
		}
	}
	return string(str)
}
