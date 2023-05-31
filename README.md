# log
这是一个go语言日志库


使用实例如下

package main

import (
	_ "github.com/ksamwang/log" //初始化

	"github.com/ksamwang/log" //引入
)

func main() {
	ksamlog.INFO("程序开始运行......")
}
