package main

import (
	_ "github.com/ksamwang/log"

	"github.com/ksamwang/log"
)

func main() {
	ksamlog.SetFormat("xml")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.SetLogMode(ksamlog.Release)
	ksamlog.SetFormat("json")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")

	ksamlog.SetFormat("txt")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.Error("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.WARN("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
	ksamlog.INFO("第%d行运行时出现的错误，不必要立即进行修复%s", 1, "请联系网络")
}
