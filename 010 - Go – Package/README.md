# Go – Package
Một trong những tính năng quan trọng của một ngôn ngữ lập trình là tính năng cho phép tái sử dụng code, trong các phần trước chúng ta đã biết một tính năng cho phép tái sử dụng code đó là hàm (func hay function). Package (hay gói) là một tính năng cho phép tái sử dụng code ở phạm vi rộng hơn, hầu hết các ngôn ngữ cấp cao như Java, C#… đều hỗ trợ package, và cả Go cũng vậy.

Trong các phần trước hầu như đoạn code nào chúng ta cũng có dòng này:

```go
import "fmt"
```
Dòng đó có nghĩa là chúng ta yêu cầu được sử dụng package fmt, package fmt chứa các hàm làm công việc liên quan đến định dạng dữ liệu ra.  Việc sử dụng package có các lợi ích sau:

Giảm thiểu rủi rõ trùng lắp tên hàm. Chẳng hạn như trong gói fmt có hàm Println(), chúng ta có thể định nghĩa một gói khác cũng có hàm Println() nhưng 2 hàm này khác nhau vì chúng nằm ở 2 gói khác nhau.
Dễ dàng tổ chức code hơn, do đó giúp chúng ta tìm các hàm cần dùng dễ dàng hơn.

Tốc độ biên dịch nhanh, bởi vì trình biên dịch không biên dịch lại code trong các package.

###Tạo package

Chúng ta sẽ viết package math chứa hàm Average() tính giá trị trung bình của một dãy số.

Các package được viết ra sẽ được đặt trong thư mục được định nghĩa trong biến môi trường GOPATH, biến này trỏ tới thư mục mà trình biên dịch Go sẽ tìm code trong đó, nếu bạn chưa tạo thì tạo một cái (tìm cách tạo trên Google), chẳng hạn như trong máy mình biến này trỏ tới C:/Project/Go, trong thư mục này sẽ chứa 3 thư mục nữa là src, bin và pkg, thư mục src là thư mục chứa code, thưc mục bin là thư mục chứa các file .exe, thư mục pkg chứa các file thư viện liên kết tĩnh (đọc thêm Phân biệt thư viện liên kết động và thư viện liên kết tĩnh). Khi viết code cho package thì chúng ta sẽ đặt code đó trong thư mục GOPATH/src.

Bây giờ chúng ta tạo một thư mục là myMath trong thư mục GOPATH/src, trong thư mục này tạo tiếp một thư mục có tên math, trong thư mục myMath/math chúng ta tạo file có tên math.go như sau:

GOPATH/src/myMath/math/math.go
```go
package math
 
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x :=  range xs {
        total += x
    }
    return total / float64(len(xs))
}
```
Vây là xong, chúng ta đã tạo gói myMath/math có hàm Average() tính giá trị trung bình của một dãy số.

Bây giờ chúng ta có thể import gói đó để sử dụng, ví dụ:

main.go
```go
package main
 
import "fmt"
import "myMath/math"
 
func main() {
    xs := []float64{1, 2, 3, 4}
    avg := math.Average(xs)
    fmt.Println(avg)
}
```
Chạy đoạn code trên chúng ta sẽ ra được kết quả là 2.5. Bạn có thể chạy lệnh go install trong thư mục GOPATH/src/myMath/math để tạo file thư viện nếu muốn.

Output
```
2.5
```
Có một số lưu ý như sau:

Go cũng có một package có tên là math, tuy chúng ta cũng đặt tên package của mình là math nhưng 2 package này khác nhau vì đường dẫn package của chúng ta là myMath/math.
Khi dùng lệnh import thì chúng ta phải ghi rõ ràng đường dẫn ra, như myMath/math, nhưng trong file math.go thì dòng khai báo package chỉ dùng tên ngắn thôi, ví dụ package math.
Khi gọi hàm thì chúng ta cũng chỉ dùng tên ngắn, ví dụ math.Average(...). Nếu giả sử bạn dùng cả 2 gói math của Go và myMath/math thì bạn có thể đặt tên giả cho package để phân biệt chúng, ví dụ:
```go
import m "myMath/math"
import "math"
```
Trong đó m là tên giả của package myMath/math.

Các hàm được định nghĩa trong một package nếu có tên có chữ cái đầu tiên viết hoa thì mới có thể gọi được từ package khác, nếu chữ cái đầu tiên viết thường thì không gọi được.
Tên package phải trùng với tên thư mục của file chứa nó, ví dụ file math.go phải được đặt trong thư mục math.