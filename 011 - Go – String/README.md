# Go – String
Go có gói strings hỗ trợ các thao tác với chuỗi rất mạnh.

###Các hàm thao tác với chuỗi

Gói strings cung cấp một số hàm thao tác với string thường dùng như sau:

string_functions.go
```go
package main
 
import (
    "fmt"
    s "strings"
)
 
func main() {
    fmt.Println("Contains:  ", s.Contains("test", "es"))
    fmt.Println("Count:     ", s.Count("test", "t"))
    fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
    fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
    fmt.Println("Index:     ", s.Index("test", "e"))
    fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))
    fmt.Println("Repeat:    ", s.Repeat("a", 5))
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))
    fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))
    fmt.Println("ToLower:   ", s.ToLower("TEST"))
    fmt.Println("ToUpper:   ", s.ToUpper("test"))  
 
    fmt.Println("Len: ", len("hello"))
    fmt.Println("Char: ", "hello"[1])
}
```
Trong đó:

Hàm Contains(a, b) cho biết chuỗi b có nằm trong chuỗi a hay không.
Hàm Count(a, b) đếm số lần xuất hiện của chuỗi b trong chuỗi a
Hàm HasPrefix(a, b) cho biết chuỗi b có bắt đầu từ vị trí đầu tiên trong chuỗi a hay không
Hàm HasSuffix(a, b) cho biết chuỗi b có bắt đầu từ vị trí cuối cùng trong chuỗi a hay không
Hàm Index(a, b) cho biết vị trí đầu tiên mà chuỗi b xuất hiện trong chuỗi a từ trái qua phải
Hàm Join([]a, b) nối các phần tử trong slice a lại thành một chuỗi mới, các phần tử ngăn cách nhau bởi chuỗi b
Hàm Repeat(a, b) tạo ra một chuỗi bằng cách lặp lại b lần chuỗi a.
Hàm Replace(a, b, c, d) thay các kí tự b trong chuỗi a thành kí tự c với số lần là d, nếu d < 0 thì thay tất cả.
Hàm Split(a, b) tách chuỗi a thành các chuỗi nhỏ hơn dựa theo kí tự phân tách là b
Hàm ToLower(a) chuyển chuỗi a thành viết hoa
Hàm ToUpper(a) chuyển chuỗi a thành viết thường
Ngoài ra chúng ta còn có 2 hàm không nằm trong gói strings nhưng cũng làm việc với chuỗi là hàm len() và phép toán [], hàm len() lấy độ dài của chuỗi, phép toán [i] lấy kí tự tại vị trí i.
Output
```
Contains:  true
Count:     2
HasPrefix: true
HasSuffix: true
Index:     1
Join:      a-b
Repeat:    aaaaa
Replace:   f00
Replace:   f0o
Split:     [a b c d e]
ToLower:   test
ToUpper:   TEST
Len:       5
Char:      101
```
Go còn cho phép chúng ta chuyển một string thành một slice như sau:
```go
arr := []byte("test")
str := string([]byte{"t", "e", "s", "t"})
```
###Định dạng chuỗi

Chúng ta có thể định dạng dữ liệu xuất ra màn hình. Ví dụ:

string_format.go
```go
package main
 
import "fmt"
import "os"
 
type point struct {
    x, y int
}
 
func main() {
    p := point{1, 2}
 
    fmt.Printf("%v\n", p)
    fmt.Printf("%v+\n", p)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%T\n", p)
 
    fmt.Printf("%t\n", true)
 
    fmt.Printf("%d\n", 123)
    fmt.Printf("%b\n", 14)
    fmt.Printf("%c\n", 33)
    fmt.Printf("%x\n", 456)
 
    fmt.Printf("%f\n", 78.9)
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
 
    fmt.Printf("%s\n", "\"string\"")
    fmt.Printf("%q\n", "\"string\"")
 
    fmt.Printf("%x\n", "hex this")
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    fmt.Printf("|%-6s|%-6s|", "foo", "b")
  
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
}
```
Hàm Printf(<định dạng>, a) cho phép chúng ta in dữ liệu theo các dạng nhất định, tham số đầu tiên của hàm này là dạng chuỗi sẽ được in ra, các chuỗi này có kí tự đầu tiên là %, các kí tự tiếp theo mô tả kiểu in của dữ liệu, tham số thứ 2 là dữ liệu, trong ví dụ trên thì chúng ta có định nghĩa một struct tên là point có 2 trường kiểu int là x, y, tiếp theo chúng ta in một số dữ liệu ra màn hình bằng hàm Printf():

%v: in các giá trị của một biến struct
%+v: in các giá trị kèm với tên trường của biến struct
%#v: giống %+v kèm theo tên kiểu dữ liệu của struct và tên hàm đã gọi nó
%T: in tên struct và tên hàm đã gọi nó
%t: in giá trị boolean
%d: in số nguyên (hệ 10)
%b: in số nguyên dưới dạng số nhị phân (hệ 2)
%c: in kí tự dựa theo mã ASCII
%x: in số nguyên dưới dạng số thập lục phân (hệ 16) hoặc chuyển một chuỗi thành số thập lục phân
%f: in số thập phân
%e và %E: in số thập phân dưới dạng số mũ
%s: in một chuỗi
%q: in một chuỗi có 2 cặp dấu nháy kép “”
%6d: in một số nguyên, nếu số đó không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
%6.2f: in số thập phân, làm tròn đến 2 chữ số thập phân, sau đó nếu phần thập phân và phần nguyên cùng với dấu chấm không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
%-6.2f: tương tự với %6.2f nhưng các khoảng trống được thêm vào bên phải
%6s: in một chuỗi, nếu chuỗi không đủ 6 kí tự thì thêm các khoảng trống vào bên trái cho đủ
%-6s: tương tự %6s nhưng thêm các khoảng trống vào bên phải
Ngoài ra chúng ta có hàm Sprintf() chỉ trả về một chuỗi chứ không in ra màn hình.

Output
```
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0xc04203c1d0
| 12| 345|| 1.20| 3.45|
|1.20 |3.45 || foo| b|
|foo |b |a string
```