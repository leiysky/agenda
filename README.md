# Agenda
A simple meetings management system based on Golang.
## 安装
使用指令：
``` shell
$ go get -u github.com/leiysky/agenda
```
进行安装。
## 使用
```shell
agenda -h
```
使用指令获取更多帮助

第一次运行时需要管理员权限，用于创建保存数据的json文件

## 架构介绍

整个项目采用的是一个MVC的模式。

### model

model主要负责处理数据。

看到使用JSON进行存储，我第一个想到的就是面向document的NoSQL存储，于是我就动手实现了一下。

与databse交互需要一个`client`。

底层的`client`需要做的事情很简单也很复杂，它不应该不包含任何业务逻辑，但是又要提供一套简单而健壮的方案，所以我尽可能的进行了抽象。

通过`GetClient()`方法可以获取一个`*ClientType`类型的实例。

通过`ClientType`实例可以访问当前数据库的值，以下是一个例子：
```go
import "github.com/leiysky/agenda/services/store"

func main() {
  // 获取client实例
  client := store.GetClient()
  // 获取collection实例
  collection := client.DB.Collection
}
```
数据处理完毕后，需要提交时先使用`client.Commit()`尝试进行提交，如果无法提交，比如处理过程中的数据有误，则会返回一个`error`。

确认提交成功之后，使用`client.Dump()`将数据持久化。

以下是一个例子：
```go
import "github.com/leiysky/agenda/services/store"

func storeSomething() error {
  client := store.GetClient()
  if err := client.Commit(); err != nil {
    return err
  }
  return client.Dump()
}
```

### service 

service层主要负责处理业务逻辑，有了model呈现的数据之后，我们可以使用service来处理数据。

在这个架构中主要提供了两种实现方式：
* 通过`client`的指针直接对`client`进行修改，之后`Dump`
* 通过`getSomething`和`setSomething`的`DAO`接口直接进行修改

### conroller

controller层主要负责处理请求。`commands`目录下的每一个文件均为一个`controller`，负责处理请求，并且调用对应的`service`。

每个`controller`需要实现一个`command`，以下是一个demo：
```go
import (
  "github.com/leiysky/agenda/services/store"
  "github.com/spf13/cobra"
 )

type options struct {
  refs []string
}

func newSomeCommand() *cobra.Command {
  cmd := &cobra.Command{
    ...
    RunE: runSomeCommand(),
  }
  return cmd
}

func runSomeCommand() error {
  //do something
}

```
之后统一挂载在`commands.go`中的根命令下。

### 初始化
程序在开始运行时，会尝试创建`/var/agenda/`下的文件，所以第一次使用时需要以root权限来运行。
