最近由于用go做字符串处理，用到了go的strings库，借此对go strings库做个总结，将go strings中所有函数的功能做一个简单的说明，当然，这是一个重复造轮子的过程，因为go语言标准库已经有中文版了。

🔗链接: [Golang标准库文档](https://studygolang.com/pkgdoc)

## strings
package strings
import "strings"

strings包实现了用于操作字符的简单函数。


## 函数列表
**func EqualFold**

```go
func EqualFold(s, t string) bool
```
判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。
Example

```go
fmt.Println(strings.EqualFold("Go", "go"))
```

Output:

```go
true
```

**func HasPrefix**

```go
func HasPrefix(s, prefix string) bool
```

判断s是否有前缀字符串prefix。

**func HasSuffix**

```go
func HasSuffix(s, suffix string) bool
```

判断s是否有后缀字符串suffix。

**func Contains**

```go
func Contains(s, substr string) bool
```

判断字符串s是否包含子串substr。

Example

```go
fmt.Println(strings.Contains("seafood", "foo"))
fmt.Println(strings.Contains("seafood", "bar"))
fmt.Println(strings.Contains("seafood", ""))
fmt.Println(strings.Contains("", ""))
```
Output:

```go
true
false
true
true
```

**func ContainsRune**

```go
func ContainsRune(s string, r rune) bool
```

判断字符串s是否包含utf-8码值r。

**func ContainsAny**

```go
func ContainsAny(s, chars string) bool
```

判断字符串s是否包含字符串chars中的任一字符。
Example

```go
fmt.Println(strings.ContainsAny("team", "i"))
fmt.Println(strings.ContainsAny("failure", "u & i"))
fmt.Println(strings.ContainsAny("foo", ""))
fmt.Println(strings.ContainsAny("", ""))
```

Output:

```go
false
true
false
false
```

**func Count**

```go
func Count(s, sep string) int
```

返回字符串s中有几个不重复的sep子串。

Example

```go
fmt.Println(strings.Count("cheese", "e"))
fmt.Println(strings.Count("five", "")) // before & after each rune
```

Output:

```go
3
5
```
**func Index**

```go
func Index(s, sep string) int
```

子串sep在字符串s中第一次出现的位置，不存在则返回-1。


**func IndexByte**

```go
func IndexByte(s string, c byte) int
```

字符c在s中第一次出现的位置，不存在则返回-1。

**func IndexRune**

```go
func IndexRune(s string, r rune) int
```

unicode码值r在s中第一次出现的位置，不存在则返回-1。
Example

```go
fmt.Println(strings.IndexRune("chicken", 'k'))
fmt.Println(strings.IndexRune("chicken", 'd'))
```

Output:

```go
4
-1
```

**func IndexAny**

```go
func IndexAny(s, chars string) int
```

字符串chars中的任一utf-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1。

Example

```go
fmt.Println(strings.IndexAny("chicken", "aeiouy"))
fmt.Println(strings.IndexAny("crwth", "aeiouy"))
```

Output:

```go
2
-1
```

**func IndexFunc**

```go
func IndexFunc(s string, f func(rune) bool) int
```

s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1。

Example

```go
f := func(c rune) bool {
    return unicode.Is(unicode.Han, c)
}
fmt.Println(strings.IndexFunc("Hello, 世界", f))
fmt.Println(strings.IndexFunc("Hello, world", f))
```

Output:

```go
7
-1
```

**func LastIndex**

```go
func LastIndex(s, sep string) int
```

子串sep在字符串s中最后一次出现的位置，不存在则返回-1。

Example

```go
fmt.Println(strings.Index("go gopher", "go"))
fmt.Println(strings.LastIndex("go gopher", "go"))
fmt.Println(strings.LastIndex("go gopher", "rodent"))
```

Output:

```go
0
3
-1
```

**func LastIndexAny**

```go
func LastIndexAny(s, chars string) int
```

字符串chars中的任一utf-8码值在s中最后一次出现的位置，如不存在或者chars为空字符串则返回-1。

**func LastIndexFunc**

```go
func LastIndexFunc(s string, f func(rune) bool) int
```

s中最后一个满足函数f的unicode码值的位置i，不存在则返回-1。

**func Title**

```go
func Title(s string) string
```

返回s中每个单词的首字母都改为标题格式的字符串拷贝。

BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。

Example

```go
fmt.Println(strings.Title("her royal highness"))
```
Output:

```go
Her Royal Highness
```

**func ToLower**

```go
func ToLower(s string) string
```

返回将所有字母都转为对应的小写版本的拷贝。
Example

```go
fmt.Println(strings.ToLower("Gopher"))
```

Output:

```go
gopher
```

**func ToLowerSpecial**

```go
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
```

使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝。

**func ToUpper**

```go
func ToUpper(s string) string
```

返回将所有字母都转为对应的大写版本的拷贝。
Example

```go
fmt.Println(strings.ToUpper("Gopher"))
```

Output:

```go
GOPHER
```

**func ToUpperSpecial**

```go
func ToUpperSpecial(_case unicode.SpecialCase, s string) string
```

使用_case规定的字符映射，返回将所有字母都转为对应的大写版本的拷贝。

**func ToTitle**

```go
func ToTitle(s string) string
```

返回将所有字母都转为对应的标题版本的拷贝。

Example
func ToTitleSpecial
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
使用_case规定的字符映射，返回将所有字母都转为对应的标题版本的拷贝。

func Repeat
func Repeat(s string, count int) string
返回count个s串联的字符串。

Example

```go
fmt.Println(strings.ToTitle("loud noises"))
fmt.Println(strings.ToTitle("хлеб"))
```

Output:

```go
LOUD NOISES
ХЛЕБ
```

**func Replace**

```go
func Replace(s, old, new string, n int) string
```

返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。

Example

```go
fmt.Println("ba" + strings.Repeat("na", 2))
```

Output:

```go
banana
```

**func Map**

```go
func Map(mapping func(rune) rune, s string) string
```

将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换。（返回值中对应位置将没有码值）

Example

```go
rot13 := func(r rune) rune {
    switch {
    case r >= 'A' && r <= 'Z':
        return 'A' + (r-'A'+13)%26
    case r >= 'a' && r <= 'z':
        return 'a' + (r-'a'+13)%26
    }
    return r
}
fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
```

Output:

```go
'Gjnf oevyyvt naq gur fyvgul tbcure...
```

**func Trim**

```go
func Trim(s string, cutset string) string
```

返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。

Example

```go
fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
```

Output:

```go
["Achtung! Achtung"]
```

**func TrimSpace**

```go
func TrimSpace(s string) string
```

返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。

Example

```go
fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
```

Output:

```go
a lone gopher
```

**func TrimFunc**

```go
func TrimFunc(s string, f func(rune) bool) string
```

返回将s前后端所有满足f的unicode码值都去掉的字符串。

**func TrimLeft**

```go
func TrimLeft(s string, cutset string) string
```

返回将s前端所有cutset包含的utf-8码值都去掉的字符串。

**func TrimLeftFunc**

```go
func TrimLeftFunc(s string, f func(rune) bool) string
```

返回将s前端所有满足f的unicode码值都去掉的字符串。

**func TrimPrefix**

```go
func TrimPrefix(s, prefix string) string
```

返回去除s可能的前缀prefix的字符串。

Example

```go
var s = "Goodbye,, world!"
s = strings.TrimPrefix(s, "Goodbye,")
s = strings.TrimPrefix(s, "Howdy,")
fmt.Print("Hello" + s)
```

Output:

```go
Hello, world!
```

**func TrimRight**

```go
func TrimRight(s string, cutset string) string
```

返回将s后端所有cutset包含的utf-8码值都去掉的字符串。

**func TrimRightFunc**

```go
func TrimRightFunc(s string, f func(rune) bool) string
```

返回将s后端所有满足f的unicode码值都去掉的字符串。

**func TrimSuffix**

```go
func TrimSuffix(s, suffix string) string
```

返回去除s可能的后缀suffix的字符串。

Example

```go
var s = "Hello, goodbye, etc!"
s = strings.TrimSuffix(s, "goodbye, etc!")
s = strings.TrimSuffix(s, "planet")
fmt.Print(s, "world!")
```

Output:

```go
Hello, world!
```

**func Fields**

```go
func Fields(s string) []string
```

返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。

Example

```go
fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
```

Output:

```go
Fields are: ["foo" "bar" "baz"]
```

**func FieldsFunc**

```go
func FieldsFunc(s string, f func(rune) bool) []string
```

类似Fields，但使用函数f来确定分割符（满足f的unicode码值）。如果字符串全部是分隔符或者是空字符串的话，会返回空切片。

Example

```go
f := func(c rune) bool {
    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}
fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
```

Output:

```go
Fields are: ["foo1" "bar2" "baz3"]
```

**func Split**

```go
func Split(s, sep string) []string
```

用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。

Example

```go
fmt.Printf("%q\n", strings.Split("a,b,c", ","))
fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
fmt.Printf("%q\n", strings.Split(" xyz ", ""))
fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
```

Output:

```go
["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```

**func SplitN**

```go
func SplitN(s, sep string, n int) []string
```

用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：

```bash
n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
n == 0: 返回nil
n < 0 : 返回所有的子字符串组成的切片
```

Example

```go
fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
z := strings.SplitN("a,b,c", ",", 0)
fmt.Printf("%q (nil = %v)\n", z, z == nil)
```

Output:

```go
["a" "b,c"]
[] (nil = true)
```

**func SplitAfter**

```go
func SplitAfter(s, sep string) []string
```

用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。

Example

```go
fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
```

Output:

```go
["a," "b," "c"]
```

**func SplitAfterN**

```go
func SplitAfterN(s, sep string, n int) []string
```

用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：

```bash
n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
n == 0: 返回nil
n < 0 : 返回所有的子字符串组成的切
```

Example

```go
fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
```

Output:

```go
["a," "b,c"]
```

**func Join**

```go
func Join(a []string, sep string) string
```

将一系列字符串连接为一个字符串，之间用sep来分隔。

Example

```go
s := []string{"foo", "bar", "baz"}
fmt.Println(strings.Join(s, ", "))
```

Output:

```go
foo, bar, baz
```

**type Reader**

```go
type Reader struct {
    // 内含隐藏或非导出字段
}
```

Reader类型通过从一个字符串读取数据，实现了io.Reader、io.Seeker、io.ReaderAt、io.WriterTo、io.ByteScanner、io.RuneScanner接口。

**func NewReader**

```go
func NewReader(s string) *Reader
```

NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString，但是更有效率，且为只读的。

**func (*Reader) Len**

```go
func (r *Reader) Len() int
```

Len返回r包含的字符串还没有被读取的部分。

**func (*Reader) Read**

```go
func (r *Reader) Read(b []byte) (n int, err error)
```

**func (*Reader) ReadByte**

```go
func (r *Reader) ReadByte() (b byte, err error)
```

**func (*Reader) UnreadByte**

```go
func (r *Reader) UnreadByte() error
```

**func (*Reader) ReadRune**

```go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```

**func (*Reader) UnreadRune**

```go
func (r *Reader) UnreadRune() error
```

**func (*Reader) Seek**

```go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

Seek实现了io.Seeker接口。

**func (*Reader) ReadAt**

```go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```

**func (*Reader) WriteTo**

```go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo实现了io.WriterTo接口。

**type Replacer**

```go
type Replacer struct {
    // 内含隐藏或非导出字段
}
```

Replacer类型进行一系列字符串的替换。

**func NewReplacer**

```go
func NewReplacer(oldnew ...string) *Replacer
```

使用提供的多组old、new字符串对创建并返回一个*Replacer。替换是依次进行的，匹配时不会重叠。

Example

```go
r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
fmt.Println(r.Replace("This is <b>HTML</b>!"))
```

Output:

```go
This is &lt;b&gt;HTML&lt;/b&gt;!
```

**func (*Replacer) Replace**

```go
func (r *Replacer) Replace(s string) string
```

Replace返回s的所有替换进行完后的拷贝。

**func (*Replacer) WriteString**

```go
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```

WriteString向w中写入s的所有替换进行完后的拷贝。






































































