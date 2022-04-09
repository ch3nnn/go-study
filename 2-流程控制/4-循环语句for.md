### 循环语句for

Golang for支持三种循环方式，包括类似 while 的语法。

#### for循环是一个循环控制结构，可以执行指定次数的循环。

语法
Go语言的For循环有3中形式，只有其中的一种使用分号。
```go
/* 第一种 */
for init; condition; post{
	
}

/* 第二种 */
for condition{

}

/* 第三种 */
for {

}
```
* init： 一般为赋值表达式，给控制变量赋初值； 
* condition： 关系表达式或逻辑表达式，循环控制条件； 
* post： 一般为赋值表达式，给控制变量增量或减量。

for语句执行过程如下：

  ①先对表达式 init 赋初值；

  ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition； 
  否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

实例

```go
package main

import "fmt"

func main() {
  str := "abcd"

  for i, n := 0, len(str); i < n; i++ {
    fmt.Println(str[i])
  }

  n := len(str)

  for n > 0 {
    fmt.Println(str[n])
    n--
  }

  for {
    fmt.Println(str)
  }

}

```
#### 循环嵌套
在 for 循环中嵌套一个或多个 for 循环

语法

以下为 Go 语言嵌套循环的格式：

```go
for [condition |  ( init; condition; increment ) | Range] {
    for [condition |  ( init; condition; increment ) | Range]{
        statement(s)
	}
    statement(s)
}
```
实例：

以下实例使用循环嵌套来输出 2 到 100 间的素数：
```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i)
      }
   }  
}

```
#### 无限循环

如过循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：
```go
package main

import "fmt"

func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
    }
}
```