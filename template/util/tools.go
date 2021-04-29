package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

//检查某字符是否存在文件里
func CheckFileContainsChar(filename, s string) bool {
	data := ReadFile(filename)
	if len(data) > 0 {
		return strings.LastIndex(data, s) > 0
	}
	return false
}

//读取文件内容
func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

//写文件
func WriteFile(filename string, data string) (count int, err error) {
	var f *os.File
	if IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

//追加写文件
func WriteFileAppend(filename string, data string) (count int, err error) {
	var f *os.File
	if IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

//创建文件
func CreateFile(path string) bool {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

//创建目录
func CreateDir(path string) bool {
	if IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//生成目录,不存在则创建,存在则加/
func GenerateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("directory is null")
	}
	last := path[len(path)-1:]
	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if !IsDir(path) {
		if CreateDir(path) {
			return path, nil
		}
		return "", errors.New(path + "Failed to create or insufficient permissions")
	}
	return path, nil
}

//判断文件 或 目录是否存在
func IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断给定文件名是否是一个目录
// 如果文件名存在并且为目录则返回 true。如果 filename 是一个相对路径，则按照当前工作目录检查其相对路径。
func IsDir(filename string) bool {
	return isFileOrDir(filename, true)
}

// 判断给定文件名是否为一个正常的文件
// 如果文件存在且为正常的文件则返回 true
func IsFile(filename string) bool {
	return isFileOrDir(filename, false)
}

// 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
func isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

//将字符串转换成驼峰格式
// Capitalize 字符首字母大写
func Capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val) // 后文有介绍
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
					vv[i] -= 32 // string的码表相差32位
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

//将分隔_拆掉,全大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//转json
func ToJson(s interface{}) string {
	js, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(js)
}

//创建目录
func CreateDir1(path string) string {
	if IsDirOrFileExist(path) == false {
		b := CreateDir(path)
		if !b {
			log.Fatalf("Directory created failed>>%s\n", path)
			return ""
		}
		fmt.Printf("Directory created success:%s\n", path)
	}
	return path
}

//写文件
func WriteFile1(path, data string) (err error) {
	if _, err := WriteFile(path, data); err == nil {
		fmt.Printf("Ganerate success: %s\n", path)
		return nil
	} else {
		return errors.New(fmt.Sprintf("Create file failed>>%s", path))
	}
}

//追加写文件
func WriteAppendFile1(path, data string) (err error) {
	if _, err := WriteFileAppend(path, data); err == nil {
		fmt.Printf("Generate success:%s\n", path)
		return nil
	} else {
		return err
	}
}

//GetRootDir 获取执行路径
func GetExeRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

//获取根目录
func GetRootPath(path string) string {
	if path[len(path)-1:] == string(os.PathSeparator) {
		path = path[:len(path)-1]
	}
	return SubStr(path, 0, strings.LastIndex(path, string(os.PathSeparator)))
}

//截取字符串
func SubStr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ErrMsg(msg string, err error) interface{} {
	m := make(map[string]interface{}, 2)
	m["msg"] = msg
	m["err"] = err
	return m
}

//FMT 格式代码
func Gofmt(path string) bool {
	if IsDirOrFileExist(path) {
		if !ExecCommand("goimports", "-l", "-w", path) {
			if !ExecCommand("gofmt", "-l", "-w", path) {
				return ExecCommand("go", "fmt", path)
			}
		}
		return true
	}
	return false
}

//清屏
func Clean() {
	switch GetOs() {
	case Darwin, Linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case Window:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//获取操作系统
func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return Darwin
	case "windows":
		return Window
	case "linux":
		return Linux
	default:
		return Unknown
	}
}

func ExecCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}

//拼接特殊字符串
func FormatField(field string, formats []string) string {
	if len(formats) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return "`" + strings.TrimRight(buf.String(), " ") + "`"
}

//添加注释 //
func AddToComment(s string, suff string) string {
	if strings.EqualFold(s, "") {
		return ""
	}
	return "// " + s + suff
}

//判断是否包存在某字符
func InArrayString(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

//检查字符串,去掉特殊字符
func CheckCharDoSpecial(s string, char byte, regs string) string {
	reg := regexp.MustCompile(regs)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr)-1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}
func CheckCharDoSpecialArr(s string, char byte, reg string) []string {
	s = CheckCharDoSpecial(s, char, reg)
	return strings.Split(s, string(char))
}

// 添加``符号
func AddQuote(str string) string {
	return "`" + str + "`"
}

// 去掉 `符号
func CleanQuote(str string) string {
	return strings.Replace(str, "`", "", -1)
}
func GetMysqlDir(Path string) string {
	return CreateDir1(Path)
}
