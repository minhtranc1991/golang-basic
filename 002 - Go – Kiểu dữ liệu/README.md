# Go – Kiểu dữ liệu
Trong phần trước chúng ta đã sử dụng kiểu dữ liệu string cho chuỗi Hello World. Kiểu dữ liệu phân nhóm các giá trị có liên quan với nhau, có các thao tác có thể thực hiện trên chúng và cách chúng được lưu trữ.

### Kiểu dữ liệu trong Go là tĩnh, tức là không thể thay đổi được không giống như một số ngôn ngữ như PHP, Javascript… cho phép thay đổi kiểu dữ liệu trong suốt quá trình chạy.

Trong phần này chúng ta sẽ tìm hiểu về một số kiểu dữ liệu cơ bản trong Go là kiểu số nguyên, số thực, kiểu string và kiểu boolean.

### Số nguyên

Đây là các giá trị số nguyên giống hoàn toàn như trong toán, tuy nhiên số nguyên trong máy tính có giới hạn (tức không có giá trị nào là vô cùng ∞ cả). Các kiểu số nguyên trong Go là uint8, uint16, uint32, uint64, int8, int16, int32, int64. Các con số 8, 16, 32, 64 có nghĩa là máy tính cần dùng bao nhiêu bit để biểu diễn số nguyên đó. uint tức là unsigned int – là kiểu số nguyên không âm. Bảng dưới đây cho biết giới hạn của từng loại kiểu số nguyên:
```
KIỂU	GIỚI HẠN
uint8	0                       – 255
uint16	0                       – 65535
uint32	0                       – 4294967295
uint64	0                       – 18446744073709551615
int8	-128                    – 127
int16	-32768                  – 32767
int32	-2147483648             – 2147483647
int64	-9223372036854775808    – 9223372036854775807
```
Kiểu uint8 còn có tên khác là byte, kiểu int32 có tên khác là rune. 

Đặc biệt trong Go còn có 3 kiểu số nguyên phụ thuộc hệ điều hành là uint, int và uintptr, 3 kiểu dữ liệu này có giới hạn giống như kiến trúc của hệ điều hành mà bạn đang sử dụng. Ví dụ nếu bạn đang dùng Windows 64 bit thì kiểu int sẽ có giới hạn giống như kiểu uint64. Thông thường khi sử dụng số nguyên bạn dùng int là đủ rồi. Ví dụ:

number.go
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println("1 + 1 = ", 1 + 1)
}
```
Kết quả:

Output
```
1 + 1 = 2
```
### Số thực

Đây là các giá trị số có phần thập phân, ví dụ 1.234, 123.4… Việc lưu trữ cũng như thao tác với số thực trong máy tính khá phức tạp nên chúng ta cũng không đi sâu vào làm gì. Ở đây chúng ta chỉ có một số lưu ý như sau:

Số thực không bao giờ chính xác một cách tuyệt đối, rất khó để biểu diễn chính xác một số thực. 

Ví dụ như phép trừ 1.01 – 0.99 sẽ cho ra kết quả là 0.020000000000000018 chứ không phải là 0.02 như bạn vẫn nghĩ.

Cũng giống như số nguyên, số thực trong máy tính cũng có nhiều giới hạn khác nhau.

Trong Go có 2 kiểu số thực là float32 và float64, 2 kiểu số phức là complex64 và complex128. Thông thường để biểu diễn số thực, bạn chỉ cần dùng float64 là đủ.

### String

String (chuỗi) là các kí tự được bọc trong cặp dấu nháy kép hoặc nháy đơn được dùng để biểu diễn văn bản. String nằm trong dấu nháy kép có thể sử dụng các kí tự điều khiển đặc biệt như \n là xuống dòng, \t là dấu tab.

Chúng ta có thể thực hiện một số thao tác thường dùng trên String như tính độ dài chuỗi, lấy kí tự tại vị trí nhất định, nối chuỗi. 

Ví dụ:

string.go
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
} 
```
Dấu cách cũng được tính là một kí tự, do đó chiều dài chuỗi “Hello World” là 11.

Các kí tự trong một chuỗi được đánh số thứ tự từ 0. Câu lệnh “Hello World”[1] sẽ cho ra kết quả là kí tự ở vị trí số 2, tuy nhiên nếu bạn chạy đoạn code trên thì kết quả sẽ cho ra là số 101 chứ không phải kí tự ‘e’, là bởi vì 101 là mã ASCII của kí tự ‘e’

Chúng ta có thể dùng phép toán + lên 2 chuỗi, kết quả là một chuỗi mới được nối từ 2 chuỗi đầu.

Output
```
11
101
Hello World
```
### Boolean

Các giá trị boolean là các giá trị 1 bit, có 2 giá trị là true và false được dùng để biểu diễn ý nghĩa ĐÚNG hoặc SAI. Trong Go có 3 phép toán có thể thao tác với giá trị boolean là && (phép AND), || (phép OR) và ! (phép NOT). 

Bảng dưới đây mô tả cách thực hiện phép toán với kiểu boolean:

Phép AND:
```
BIỂU THỨC	GIÁ TRỊ
true && true	true
true && false	false
false && true	false
false && false	false
```
Phép OR:
```
BIỂU THỨC	GIÁ TRỊ
true || true	true
true || false	true
false || true	true
false || false	false
```
Phép NOT:
```
BIỂU THỨC	GIÁ TRỊ
!true	false
!false	true
```
Ví dụ:

boolean.go
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println(true && false)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```
Kết quả:

Output
```
true
false
true
true
false
```