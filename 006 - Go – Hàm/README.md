# Go – Hàm
Hàm là một khối lệnh độc lập có chức năng nhận dữ liệu, xử lý và trả về kết quả. Trong các phần trước chúng ta chỉ làm việc với một hàm duy nhất là hàm main(), trong phần này chúng ta sẽ làm việc với nhiều hàm khác.

###Khai báo hàm

Giả sử chúng ta có đoạn code tính giá trị trung bình của một slice như sau:

function.go
```go
package main
 
import "fmt"
 
func main() {
 
    xs := []float64{98, 93, 77, 82, 83}
    total := 0.0
    for _, v := range xs {
        total += v
    }
    fmt.Println(total / float64(len(xs))
}
```
Bài toán tính giá trị trung bình của một dãy số là rất phổ biến, do đó chúng ta nên định nghĩa hàm riêng để thực hiện công việc này. Ví dụ:

function2.go
```go
package main
 
import "fmt"
 
func average(input []float64) float64 {
    total := 0.0
    for _, v := range input {
        total += v
    }
    return total / float64(len(input))
}
 
func main() {
    xs := []float64{98, 93, 77, 82, 83}
    fmt.Println(average(xs))
}
```
Như trong đoạn code ví dụ trên, để khai báo một hàm thì chúng ta dùng từ khóa func, tiếp theo là tên hàm, rồi danh sách các tham số đầu vào trong cặp dấu ngoặc tròn (), rồi đến kiểu dữ liệu trả về, cuối cùng là phần thân hàm nằm trong cặp dấu ngoặc nhọn {}. Tham số đầu vào và kiểu dữ liệu trả về có thể không có cũng được.

Trong đoạn code trên chúng ta định nghĩa hàm average() có kiểu trả về là float64, hàm này nhận một tham số đầu vào là biến xs. Trong phần thân hàm chúng ta thay câu lệnh fmt.Println() thành câu lệnh return, câu lệnh return có chức năng kết thúc hàm và “trả về” một giá trị cho hàm đã gọi nó, để gọi một hàm thì chúng ta chỉ đơn giản là ghi tên hàm đó ra rồi đưa tham số vào, ở đây chúng ta gọi hàm average() trong câu lệnh fmt.Println(average(xs)) trong hàm main(). Cả 2 đoạn code trên đều cho ra kết quả giống nhau.

Có một số lưu ý như sau:

Một hàm không thể đọc một biến được định nghĩa trong hàm khác, ví dụ:
```go
func f() {
    fmt.Println(x)
}
 
func main() {
    x := 5
    f() 
}
```
Đoạn code trên sẽ báo lỗi, để hàm f() có thể đọc được biến x thì chúng ta phải truyền x vào làm tham số như sau:
```go
func f(x int) {
    fmt.Println(x)
}
 
func main() {
    x := 5
    f(x)
}
```
Hàm được gọi chồng lên nhau, ví dụ chúng ta có đoạn code sau:
```go
func main() {
    fmt.Println(f1())
}
 
func f1() int {
    return f2() 
}
 
func f2() int {
    return 1
}
```
Trong hàm main() chúng ta gọi hàm f1(), hàm f1() lại gọi hàm f2(), hàm f2() sẽ trả về 1 cho hàm f1(), hàm f1() lại trả về giá trị 1 đó cho hàm main(). Bạn có thể hình dung quá trình đó như sau:

[Code flow](../Capture.jpg)

Chúng ta có thể đặt tên cho giá trị trả về, ví dụ:
```go
func f2() (r int) {
    r = 1
    return
}
```
###Trả về nhiều giá trị

Go cho phép một hàm được trả về nhiều giá trị. Ví dụ:

```go
func f() (int, int) {
    return 5, 6
}
 
func main() {
   x, y := f()
}
```
Chúng ta khai báo danh sách các kiểu trả về trong cặp dấu ngoặc tròn (), ngăn cách nhau bởi dấu phẩy ",", phía sau câu lệnh return cũng là danh sách các giá trị trả về cách nhau bởi dấu phẩy ",", chúng ta cũng gán nhiều giá trị cho nhiều biến cách nhau bởi dấu phẩy ",".

Thông thường chúng ta thường trả về một giá trị của một kết quả tính toán nào đó, và một giá trị lỗi cho biết công việc của hàm có thành công hay không.

###Tùy biến số lượng tham số

Tham số cuối cùng của một hàm có thể được khai báo theo dạng đặc biệt như sau:

variadic_function.go
```go
package main
 
import "fmt"
 
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    } 
    return total
}
 
func main() {
    fmt.Println(add(1, 2, 3))
}
```
Chúng ta thêm 3 dấu chấm "..." vào trước tên kiểu dữ liệu của tham số cuối cùng, Go sẽ hiểu rằng tham số này có thể có 0 hoặc nhiều giá trị được truyền vào, khi gọi hàm chúng ta có thể truyền vào 0 hoặc nhiều giá trị, ngăn cách nhau bởi dấu phẩy, đặc tính này cho phép hàm nhận tham số một cách linh hoạt hơn.

###Hàm Closure

Chúng ta có thể định nghĩa một hàm bên trong một hàm khác. Ví dụ:

closure_function.go
```go
package main
 
import "fmt"
 
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1, 1))
}
```
Trong đoạn code trên chúng ta khai báo biến add có kiểu func(int, int) int. Các hàm được định nghĩa kiểu này có thể đọc được các biến nằm cùng hàm với nó, ví dụ:

closure_function2.go
```go
package main
 
import "fmt"
 
func main() {
    x := 0
    increment := func() int {
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
}
```
Trong đoạn code trên, increment thực hiện tăng biến x lên một đơn vị, mặc dù x được khai báo ngoài hàm increment nhưng hàm này vẫn có thể đọc được.

###Đệ quy

Một hàm có thể gọi chính nó, ví dụ:

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x - 1)
}
```
Khi một hàm gọi chính nó thì đây là kỹ thuật lập trình đệ quy.

Hàm Closure và đệ quy là các kỹ thuật lập trình cao cấp tạo nên mô hình lập trình chức năng. Hầu hết người mới học sẽ thấy hơi khó hiểu cách hoạt động của chúng so với việc dùng cách câu lệnh bình thường như for, if….

###Defer, Panic và Recover

Lệnh defer có tác dụng chạy một lệnh khác sau khi một hàm đã kết thúc. Ví dụ:

defer.go
```go
package main
 
import "fmt"
 
func first() {
    fmt.Println("1st")
}
 
func second() {
    fmt.Println("2nd")
}
 
func main() {
    defer second()
    first()
}
```
Trong ví dụ trên, hàm first() sẽ thực hiện đầu tiên, sau đó đến hàm second() mặc dù hàm second() đứng trước, lý do là bởi vì chúng ta thêm từ khóa defer vào trước second(), do đó hàm second() sẽ được thực hiện cuối cùng khi tất cả các công việc khác đã hoàn tất.

Thường chúng ta dùng defer cho các công việc dọn dẹp tài nguyên.

Hàm panic dùng để phát sinh lỗi, ví dụ:

panic.go
```go
package main
 
func main() {
    panic("Co loi xay ra") 
}
```
Đoạn code trên sẽ cho kết quả như sau:

Output
```
panic: Co loi xay ra
 
goroutine 1 [running]:
panic(0x456e20, 0xc0420040b0)
    C:/Go/src/runtime/panic.go:500 +0x1af main.main()
    C:/main.go:4 +0x74
exit status 2
```
Chẳng hạn như khi chúng ta viết chương trình tính phép chia mà người dùng nhập vào mẫu số là 0 thì chúng ta có thể dùng hàm panic để báo lỗi. Tuy nhiên hàm panic() ngoài việc thông báo lỗi sẽ dừng chương trình luôn, do vậy Go đưa ra hàm recover(), hàm recover() có tác dụng khôi phục chương trình và trả về tham số đã được truyền vào trong hàm panic(). Ví dụ:

recover.go
```go
package main
 
import "fmt"
 
func main() {
    panic("Co loi xay ra")
    str := recover()
    fmt.Println(str)
}
```
Tuy nhiên đoạn code trên sẽ không chạy vì hàm panic() đã dừng hoàn toàn chương trình trước khi chúng ta gọi hàm recover(). Do đó để dùng hàm recover() chúng ta phải bọc trong lệnh defer như sau:

recover2.go
```go
package main
 
import "fmt"
 
func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("Co loi xay ra")
}
```
Kết quả:

Output
```
Co loi xay ra
```