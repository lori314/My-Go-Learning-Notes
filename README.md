# Go 学习笔记

这是我刚开始学习 Go 时留下来的笔记和几段练习代码。内容从最基本的语法一路记到 goroutine、channel 和 select，整体更像一份随学随记的速查表，不是完整教程。

原始笔记放在：

```text
笔记.txt
```

代码练习放在：

```text
codes/
```

这次整理时主要修正了几处早期理解不准确的地方，笔记原来的顺序和写法基本保留。

## 目前记到的内容

### 基础语法

包括：

- `package`、`import`、`func main`
- `go run` 和 `go build`
- 变量、常量和零值
- `iota`
- `if`、`switch`、`for`
- 多返回值
- 类型转换

Go 的可见性规则和 Java/C++ 不太一样。包级标识符首字母大写表示 **exported**，可以被其他包访问；首字母小写表示只在当前包内可见。Go 本身没有 `public / protected / private` 这些访问修饰符。

## 数组、Slice 和 Map

笔记里分别记录了数组、slice、map 和 `range` 的基本用法。

数组是值类型，而且长度属于数组类型的一部分：

```go
var a [5]int
var b [10]int
```

这里 `a` 和 `b` 的类型不同。

Slice 更像是对底层数组一段区间的描述：

```go
s := make([]int, 5, 10)
t := arr[2:6]
```

赋值或切片后可能继续共享同一个底层数组，所以修改元素时需要注意这种共享关系。

Map 也有类似的共享语义：

```go
m := make(map[string]int)
m["apple"] = 1
delete(m, "apple")
```

## 指针、结构体和接口

笔记里还记录了：

- `&` 取地址和 `*` 解引用；
- Go 不支持指针算术；
- struct 的定义和成员访问；
- interface 的隐式实现；
- 类型断言；
- 内置 `error` 接口。

接口不要求类型显式声明“implements”。只要一个类型具有接口要求的全部方法，就可以赋给该接口。

## Goroutine

Go 的并发部分是我当时最后学到的一块。

Goroutine 是由 Go runtime 调度的轻量级并发执行单元：

```go
go f(x, y, z)
```

它不是简单地“一条 goroutine 对应一条操作系统线程”，runtime 会负责 goroutine 的调度。

## Channel

Channel 用来在 goroutine 之间传递数据，也能起到同步作用：

```go
ch := make(chan int)

ch <- 1
v := <-ch
```

也可以创建带缓冲的 channel：

```go
ch := make(chan int, 10)
```

接收时常见的写法是：

```go
v, ok := <-ch
```

只有当 channel 已关闭，并且缓冲数据也读取完以后，`ok` 才会变成 `false`。通常由发送方在确定不会继续发送时关闭 channel。

## Select

`select` 用来同时等待多个 channel 操作：

```go
select {
case x := <-ch1:
    fmt.Println(x)
case ch2 <- value:
    // send
default:
    // optional non-blocking branch
}
```

如果多个 case 同时可以执行，会从这些就绪 case 中选择一个；如果没有 case 可以执行，默认会阻塞，除非写了 `default`。

## 代码练习

`codes/` 里只有几段刚学 Go 时写的小程序：

```text
helloworld.go   Hello World
a+b.go          输入两个整数求和
cal_n.go        for 循环和函数练习
test.go         自定义 error + Newton 法求平方根
```

运行单个文件例如：

```bash
cd codes
go run helloworld.go
```

`codes/go.mod` 记录的是当时使用的 Go 版本。
