package main

import "ksamlog"

func main() {
	loginit := ksamlog.SetKsamLog("test")
	_ = loginit.InitLog()

	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")

}
