package ksamlog

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

// 如果SetKsamLog和InitLog了就不需要以初始化的方式导入
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

var logFormat string = "txt"

// SetFormat txt、json、xml
func SetFormat(formatType string) {
	if formatType == "txt" {
		logFormat = formatType
	}
	if formatType == "json" {
		logFormat = formatType
	}
	if formatType == "xml" {
		logFormat = formatType
	}
}

const (
	Debug   = 0
	Release = 1
)

var logMode int = Debug

func SetLogMode(Mode int) {
	if Debug == Mode {
		logMode = Debug
	}
	if Release == Mode {
		logMode = Release
	} else {
		logMode = Debug
	}
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
	var writer_text string
	if logMode == Debug {

		if logFormat == "txt" {
			writer_text = fmt.Sprintf("%s\t%s\tInFile:%s,Line:%d-%s\r\n", t.Format("2006-01-02 15:04:05"), "[EROR]", infile, line, b.String())
		}
		if logFormat == "json" {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[EROR]"
			w.InFile = infile
			w.InLine = strconv.Itoa(line)
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		}
		if logFormat == "xml" {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[EROR]"
			w.InFile = infile
			w.InLine = strconv.Itoa(line)
			w.Content = b.String()
			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		}

	} else {
		if logFormat == "txt" {
			writer_text = fmt.Sprintf("%s\t%s\t%s\r\n", t.Format("2006-01-02 15:04:05"), "[EROR]", b.String())
		}
		if logFormat == "json" {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[EROR]"
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		}
		if logFormat == "xml" {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[EROR]"
			w.Content = b.String()

			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		}
	}
	write_in_file(writer_text)
	_, _ = fmt.Fprintf(os.Stdout, writer_text)
}

type wlog_debug struct {
	Date    string `json:"Date" xml:"Date"`
	Level   string `json:"Level" xml:"Level"`
	InFile  string `json:"InFile" xml:"InFile"`
	InLine  string `json:"InLine" xml:"InLine"`
	Content string `json:"Content" xml:"Content"`
}

type wlog_release struct {
	Date    string `xml:"Date" json:"Date"`
	Level   string `xml:"Level" json:"Level"`
	Content string `xml:"Content" json:"Content"`
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
	var writer_text string
	if logMode == Debug {

		if logFormat == "txt" {
			writer_text = fmt.Sprintf("%s\t%s\tInFile:%s,Line:%d-%s\r\n", t.Format("2006-01-02 15:04:05"), "[WARN]", infile, line, b.String())
		}
		if logFormat == "json" {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[WARN]"
			w.InFile = infile
			w.InLine = strconv.Itoa(line)
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		}
		if logFormat == "xml" {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[WARN]"
			w.InFile = infile
			w.InLine = strconv.Itoa(line)
			w.Content = b.String()
			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		}

	} else {
		if logFormat == "txt" {
			writer_text = fmt.Sprintf("%s\t%s\t%s\r\n", t.Format("2006-01-02 15:04:05"), "[WARN]", b.String())
		}
		if logFormat == "json" {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[WARN]"
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		}
		if logFormat == "xml" {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[WARN]"
			w.Content = b.String()

			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		}
	}
	write_in_file(writer_text)

	_, _ = fmt.Fprintf(os.Stdout, writer_text)
}

func INFO(format string, a ...interface{}) {
	var b strings.Builder

	_, err := fmt.Fprintf(&b, format, a...)
	if err != nil {
		return
	}

	t := time.Now()

	var writer_text string

	if logFormat == "txt" {
		writer_text = fmt.Sprintf("%s\t%s\t%s\r\n", t.Format("2006-01-02 15:04:05"), "[INFO]", b.String())
	}
	if logFormat == "json" {
		if logMode == Debug {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[INFO]"
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		} else {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[INFO]"
			w.Content = b.String()
			marshal, err := json.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(marshal) + "\r\n"
		}

	}
	if logFormat == "xml" {
		if logMode == Debug {
			var w wlog_debug
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[INFO]"
			w.Content = b.String()

			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		} else {
			var w wlog_release
			w.Date = t.Format("2006-01-02 15:04:05")
			w.Level = "[INFO]"
			w.Content = b.String()
			bytes, err := xml.Marshal(w)
			if err != nil {
				return
			}
			writer_text = string(bytes) + "\r\n"
		}
	}
	write_in_file(writer_text)
	_, _ = fmt.Fprintf(os.Stdout, writer_text)
}
