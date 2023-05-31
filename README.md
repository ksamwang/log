# log
这是一个go语言日志库
```cmd
go get -u github.com/ksamwang/log
```
使用实例如下


```go
package main

import (
	_ "github.com/ksamwang/log" //初始化

	"github.com/ksamwang/log" //引入
)

func main() {
	ksamlog.SetFormat("xml") // 设置日志格式，支持txt.json.xml
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	
	ksamlog.SetLogMode(ksamlog.Release) //设置日志模式，默认debug
	ksamlog.SetFormat("json") // 设置日志格式
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")

	ksamlog.SetFormat("txt") // 设置日志格式
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
}
```
