# Strconv

> strconv包实现了基本数据类型和其字符串表示的相互转换。

更多函数请查看
* [官方文档](https://golang.org/pkg/strconv/)
* [中文文档](https://studygolang.com/pkgdoc)

### string与int类型转换

**Atoi()**
> Atoi()函数用于将字符串类型的整数转换为int类型，等价于ParseInt (s, 10, 0)，转换为 int 类型。
```go
func Atoi(s string) (i int, err error)
```

```go
v := "10"
if s, err := strconv.Atoi(v); err == nil {
	fmt.Printf("%T, %v", s, s)
}

//Output:
//
//int, 10
```

**Itoa()**
> Itoa()函数用于将int类型数据转换为对应的字符串表示，等价于FormatInt (int64(i), 10)。
```go
func Itoa(i int) string
```

例子
```go
i := 10
s := strconv.Itoa(i)
fmt.Printf("%T, %v\n", s, s)

//Output:
//
//string, 10
```

### Parse 系列函数

>Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()

**ParseBool**

```go
func ParseBool(str string) (value bool, err error)
```
>ParseBool 返回字符串表示的布尔值。它接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False。任何其他值都会返回错误。

例子
```go

v := "true"
if s, err := strconv.ParseBool(v); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}

//Output:
//
//bool, true
	
```

**ParseInt**

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

>ParseInt 以给定的基数（0、2 到 36）和位大小（0 到 64）解释字符串 s，并返回相应的值 i。
字符串可以以前导符号开头：“+”或“-”。

>如果 base 参数为 0，则真正的基数由符号后面的字符串前缀（如果存在）暗示：2 表示“0b”，8 表示“0”或“0o”，16 表示“0x”，否则为 10。此外，仅对于以 0 为基数的参数，允许使用 Go 语法为整数文字定义的下划线字符。

>bitSize 参数指定结果必须适合的整数类型。位大小 0、8、16、32 和 64 对应于 int、int8、int16、int32 和 int64。如果 bitSize 小于 0 或大于 64，则返回错误。

>ParseInt 返回的错误具有具体类型 * NumError并包括 err.Num = s。如果 s 为空或包含无效数字，则 err.Err = ErrSyntax ，返回值为 0；如果 s 对应的值不能用给定大小的有符号整数表示，则 err.Err = ErrRange并且返回值是适当的 bitSize 和符号的最大幅度整数

例子

```go
v32 := "-354634382"
if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}

v64 := "-3546343826724305832"
if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}

//Output:
//
//int64, -354634382
//int64, -3546343826724305832
	
```

**ParseUnit**

```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```
>ParseUint 类似于ParseInt但用于无符号数。
不允许使用符号前缀

例子
```go
v := "42"
if s, err := strconv.ParseUint(v, 10, 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseUint(v, 10, 64); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}

//Output:
//
//uint64, 42
//uint64, 42
```

**ParseFloat**
```go
func ParseFloat(s string, bitSize int) (f float64, err error)
```
>ParseFloat 将字符串 s 转换为精度由 bitSize 指定的浮点数：对于 float32 为 32，对于 float64 为 64。当 bitSize=32 时，结果仍然是 float64 类型，但它可以转换为 float32 而不会改变它的值。

>ParseFloat 接受十进制和十六进制浮点数语法。如果 s 格式正确且接近有效浮点数，则 ParseFloat 返回使用 IEEE754 无偏舍入舍入的最接近的浮点数。 （仅当十六进制表示中的位数多于尾数中的位数时，才对十六进制浮点值进行四舍五入。）

>ParseFloat 返回的错误具有具体类型 * NumError并包括 err.Num = s。

>如果 s 的语法格式不正确，则 ParseFloat 返回 err.Err = ErrSyntax 。
如果 s 在语法上格式正确，但距离给定大小的最大浮点数超过 1/2 ULP，则 ParseFloat 返回 f = ±Inf, err.Err = ErrRange 。

>ParseFloat 将字符串“NaN”和（可能有符号的）字符串“Inf”和“Infinity”识别为它们各自的特殊浮点值。匹配时忽略大小写

例子
```go
v := "3.1415926535"
if s, err := strconv.ParseFloat(v, 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat(v, 64); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("NaN", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
// ParseFloat is case insensitive
if s, err := strconv.ParseFloat("nan", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("inf", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("+Inf", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("-Inf", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("-0", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}
if s, err := strconv.ParseFloat("+0", 32); err == nil {
	fmt.Printf("%T, %v\n", s, s)
}

//Output:
//
//float64, 3.1415927410125732
//float64, 3.1415926535
//float64, NaN
//float64, NaN
//float64, +Inf
//float64, +Inf
//float64, -Inf
//float64, -0
//float64, 0
	
```

### Format系列函数

> Format系列函数实现了将给定类型数据格式化为string类型数据的功能。

**FormatBool**
```go
func FormatBool(b bool) string
```
>FormatBool 根据 b 的值返回“true”或“false”。

例子
```go
v := true
s := strconv.FormatBool(v)
fmt.Printf("%T, %v\n", s, s)

//Output:
//
//string, true
```

**FormatInt**
```go
func FormatInt(i int64, base int) string
```
>FormatInt 返回给定基数中 i 的字符串表示形式，即 2 <= base <= 36。对于 >= 10 的数字值，结果使用小写字母 'a' 到 'z'。

例子
```go
v := int64(-42)

s10 := strconv.FormatInt(v, 10)
fmt.Printf("%T, %v\n", s10, s10)

s16 := strconv.FormatInt(v, 16)
fmt.Printf("%T, %v\n", s16, s16)

//Output:
//
//string, -42
//string, -2a
```

**FormatUint**

```go
func FormatUint(i uint64, base int) string
```

> FormatUint 返回给定基数中 i 的字符串表示形式，即 2 <= base <= 36。对于 >= 10 的数字值，结果使用小写字母 'a' 到 'z'

例子

```go

v := uint64(42)

s10 := strconv.FormatUint(v, 10)
fmt.Printf("%T, %v\n", s10, s10)

s16 := strconv.FormatUint(v, 16)
fmt.Printf("%T, %v\n", s16, s16)

//Output:
//
//string, 42
//string, 2a
	
```

**FormatFloat**

```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

>FormatFloat 根据格式 fmt 和精度 prec 将浮点数 f 转换为字符串。
假设原始值是从 bitSize 位的浮点值（32 表示 float32，64 表示 float64）获得的，它会对结果进行四舍五入。
格式 fmt 是 'b'（-ddddp±ddd，二进制指数）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）之一), 'f' (-ddd.dddd, 无指数), 'g' ('e' 表示大指数，'f' 否则), 'G' ('E' 表示大指数，'f' 否则), ' x'（-0xd.ddddp±ddd，十六进制分数和二进制指数），或 'X'（-0Xd.ddddP±ddd，十六进制分数和二进制指数）。
精度 prec 控制由“e”、“E”、“f”、“g”、“G”、“x”和“X”格式打印的位数（不包括指数）。对于'e'、'E'、'f'、'x'和'X'，它是小数点后的位数。对于“g”和“G”，它是有效数字的最大数量（删除尾随零）。
特殊精度 -1 使用所需的最少位数，以便ParseFloat将准确返回 f

例子
```go

v := 3.1415926535

s32 := strconv.FormatFloat(v, 'E', -1, 32)
fmt.Printf("%T, %v\n", s32, s32)

s64 := strconv.FormatFloat(v, 'E', -1, 64)
fmt.Printf("%T, %v\n", s64, s64)

//Output:
//
//string, 3.1415927E+00
//string, 3.1415926535E+00
```

### 其他

**isPrint**

```go
func IsPrint(r rune) bool
```
>返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。

**CanBackquote**

```go
func CanBackquote(s string) bool
```
>返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。


#### 除上文列出的函数外，strconv包中还有Append系列、Quote系列等函数。具体用法可查看[官方文档](https://golang.org/pkg/strconv/)。

