### Map

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

#### map定义

Go语言中 map的定义语法如下

```go
map[KeyType]ValueType
```

其中

```text
KeyType:表示键的类型。

ValueType:表示键对应的值的类型。
```

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

#### map定义

map中的数据都是成对出现的，map的基本使用示例代码如下：

```go
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["小名"] = 18
	scoreMap["李四"] = 28

	fmt.Println(scoreMap)
	fmt.Println(scoreMap["李四"])
	fmt.Printf("type of a:%T\n", scoreMap)

}
```

输出

```text
    map[小明:100 张三:90]
    100
    type of a:map[string]int
```

map也支持在声明的时候填充元素，例如：

```go
package main

import "fmt"

func main() {
	userInfo := map[string]string{
		"username": "chenTong",
		"password": "123456",
	}

	fmt.Println(userInfo)
}

```

#### 判断某个键是否存在

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

举个例子：

```go
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["小名"] = 18
	scoreMap["李四"] = 28

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	value, ok := scoreMap["小名"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}

}
```

#### map的遍历

Go语言中使用for range遍历map。

```go
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["小名"] = 18
	scoreMap["李四"] = 28
	scoreMap["张三"] = 38

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

}
```

但我们只想遍历key的时候，可以按下面的写法：

```go
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["小名"] = 18
	scoreMap["李四"] = 28
	scoreMap["张三"] = 38

	for key := range scoreMap {
		fmt.Println(key)
	}

}
```

注意： 遍历map时的元素顺序与添加键值对的顺序无关。

#### 使用delete()函数删除键值对

使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：

```go
delete(map, key)
```

其中

```text
map: 表示要删除键值对的map
key: 表示要删除键值对的键
```

示例如下:

```go
package main

func main() {
	scoreMap := make(map[string]int)
	scoreMap["小名"] = 18
	scoreMap["李四"] = 28
	scoreMap["张三"] = 38

	// 删除小名
	delete(scoreMap, "小名")

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

}
```

#### 按照指定顺序遍历map

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)          // 生成0~99的随机整数
		scoreMap[key] = value
	}

	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

```

#### 元素为map类型的切片
下面的代码演示了切片中的元素为map类型时的操作：

```go
package main

func main() {
	
	mapSlice := make([]map[string]string, 3)
	
	for index, value := range mapSlice{
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")

	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

}
```
输出
```text
index:0 value:map[]
index:1 value:map[]
index:2 value:map[]
after init
index:0 value:map[address:红旗大街 name:王五 password:123456]
index:1 value:map[]
index:2 value:map[]
```
#### 值为切片类型的map
```go
package main

import "fmt"

func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 3)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
	
}
```
输出
```text
map[]
after init
map[中国:[北京 上海]]

```



