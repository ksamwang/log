package ksamlog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type filectl struct {
	FileName string
	Lock     sync.RWMutex
}

var g_file filectl

type KsamLog struct {
	FileName string
}

// SetKsamLog 首先使用这个 如果使用初始化的方式导入，需要给init赋予文件名
func SetKsamLog(FileName string) *KsamLog {
	return &KsamLog{FileName: FileName}
}

// InitLog 然后使用这个
func (p *KsamLog) InitLog() error {
	g_file.FileName = p.FileName
	_ = os.Mkdir(fmt.Sprintf("%s/logs", AbsDirectory()), 0664)
	return nil
}

func init() {
	run_name := os.Args[0]
	name_len := len(run_name) - 1

	dont := 0

	for ; ; name_len-- {
		if run_name[name_len] == '.' {
			dont = name_len
		}
		if run_name[name_len] == '\\' || run_name[name_len] == '/' {
			run_name = run_name[name_len+1 : dont]
			break
		}
		if name_len == 0 {
			break
		}
	}
	g_file.FileName = run_name
	_ = os.Mkdir(fmt.Sprintf("%s/logs", AbsDirectory()), 0664)
}

func AbsDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

// func get_date() string {
//	t := time.Now()
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	shanghaiTime := t.In(location)
//	return shanghaiTime.Format("2006-01-02")
// }

func get_date() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func write_in_file(text string) {
	defer g_file.Lock.Unlock()
	g_file.Lock.Lock()
	logpath := fmt.Sprintf("%s/logs/%s-%s.log", AbsDirectory(), g_file.FileName, get_date())

	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		_, _ = os.Create(logpath)
	}

	flag := os.O_RDWR | os.O_APPEND
	perm := os.FileMode(0644)
	file, err := os.OpenFile(logpath, flag, perm)
	if err != nil {
		return
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	_, _ = write.WriteString(text)
	_ = write.Flush()
}

func Error(format string, a ...interface{}) {
	var b strings.Builder
	_, err := fmt.Fprintf(&b, format, a...)
	if err != nil {
		return
	}

	t := time.Now()

	_, infile, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	writer_text := fmt.Sprintf("%s\t%s\tInFile:%s,Line:%d-%s\r\n", t.Format("2006-01-02 15:04:05"), "[EROR]", infile, line, b.String())
	write_in_file(writer_text)
	red := "\033[31m"
	reset := "\033[0m"
	_, _ = fmt.Fprintln(os.Stdout, red+writer_text+reset)
}

func WARN(format string, a ...interface{}) {
	var b strings.Builder

	_, err := fmt.Fprintf(&b, format, a...)
	if err != nil {
		return
	}

	t := time.Now()

	_, infile, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	writer_text := fmt.Sprintf("%s\t%s\tInFile:%s,Line:%d-%s\r\n", t.Format("2006-01-02 15:04:05"), "[WARN]", infile, line, b.String())
	write_in_file(writer_text)
	red := "\033[33m"
	reset := "\033[0m"
	_, _ = fmt.Fprintln(os.Stdout, red+writer_text+reset)
}

func INFO(format string, a ...interface{}) {
	var b strings.Builder

	_, err := fmt.Fprintf(&b, format, a...)
	if err != nil {
		return
	}

	t := time.Now()

	writer_text := fmt.Sprintf("%s\t%s\t%s\r\n", t.Format("2006-01-02 15:04:05"), "[INFO]", b.String())
	write_in_file(writer_text)

	_, _ = fmt.Fprintln(os.Stdout, writer_text)
}
