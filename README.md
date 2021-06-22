# 高性能的简繁体转换

> 因为 fork 无法被 go 项目导入，且改动未必是原项目期望的方向，顾克隆修改了一份出来。

## 基本使用方式
```go
dicter := DefaultDict()
dicter.Read("什麼")
dicter.ReadReverse("什么")
```
## 支持自定义加载词库
```go
DefaultDict(SetPath("/users/go/src/xxx.txt"))
```

## 性能测试
```go
goos: darwin
goarch: amd64
pkg: github.com/go-creed/sat
BenchmarkNewDict
BenchmarkNewDict-12    	14721091	        71.2 ns/op
PASS
```

