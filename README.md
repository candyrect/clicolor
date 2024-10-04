# CliColor

## 简介

目前市面上大多数的Golang彩色打印库都无法在Goland（等Jerbarins系IDE）中的`go.run.processes.with.pty`模式下正常显示。

CliColor库解决了上述问题，并且提供了方便易用的链式操作设置输出样式

## 安装

您可以使用go modules进行安装

```bash
go get github.com/candyrect/clicolor
```

## 使用

打印亮绿色文本

其中Underline意味着使用下划线样式，Bold意味使用粗体样式

```go
LightGreen.Underline().Bold().Println("Hello, World!")
```
