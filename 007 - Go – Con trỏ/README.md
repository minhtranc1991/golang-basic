# Go – Con trỏ
Trong phần này chúng ta sẽ tìm hiểu một khái niệm quan trọng đó là con trỏ.

Khi chúng ta gọi một hàm và truyền tham số vào đó, giá trị của tham số đó sẽ được sao chép vào trong hàm đó, ví dụ:

pointer.go
```go
package main
 
import "fmt"
 
func zero(x int) {
    x = 0
}
 
func main() {
    x := 5
    zero(x)
    fmt.Println(x)
}
```
Đoạn code trên sẽ cho kết quả là 5.

Trong đoạn code trên, hàm zero() có tác dụng gán cho biến x giá trị là 0. Trong hàm main() chúng ta khai báo một biến x có giá trị là 5 rồi truyền vào trong hàm zero(), sau đó chúng ta in giá trị của biến x ra màn hình, kết quả vẫn bằng 5. Lý do là vì giá trị của biến x trong hàm main() được sao chép vào tham số x của riêng hàm zero() chứ hàm zero() không nhận một biến x nào cả.

Tuy nhiên nếu chúng ta muốn hàm zero() thao tác trực tiếp luôn với biến x của hàm main() thì chúng ta phải dùng đến con trỏ. Ví dụ:

pointer2.go
```go
package main
 
import "fmt"
 
func zero(xPtr *int) {
    *xPtr = 0
}
 
func main() {
    x := 5
    zero(&x)
    fmt.Println(x)  
}
```
Con trỏ là một loại biến đặc biệt, được dùng để lưu trữ địa chỉ của biến khác trong bộ nhớ RAM chứ không lưu trữ một giá trị cụ thể nào. Để khai báo một biến con trỏ thì chúng ta thêm dấu sao "*" vào trước tên kiểu dữ liệu. Khi chúng ta in giá trị của một biến con trỏ ra màn hình thì giá trị đó sẽ là một số ở hệ hexa (hệ 16), đó là địa chỉ bộ nhớ mà con trỏ này đang “trỏ” tới.

Khi gán giá trị cho biến con trỏ, chúng ta cũng phải đưa vào đó một địa chỉ bộ nhớ nào đó chứ không đưa một giá trị vào, để lấy địa chỉ bộ nhớ của một biến thì chúng ta dùng dấu "&" trước tên biến.

Ngoài chức năng khai báo biến con trỏ, dấu "*" còn có tác dụng lấy giá trị của địa chỉ bộ nhớ mà con trỏ đang tham chiếu tới, ngược lại chúng ta cũng có thể gán giá trị cho địa chỉ đó thông qua dấu "*".

Ví dụ:

pointer3.go
```go
package main
 
import "fmt"
 
func main() {
    var x *int 
    var y int
    y = 0
    x = &y
  
    fmt.Println(x)
    fmt.Println(&y)
    fmt.Println(*x)
    fmt.Println(y)
  
    *x = 1
  
    fmt.Println(*x)
    fmt.Println(y)
}
```
Trong đoạn code trên, chúng ta :

Khai báo x là biến con trỏ kiểu int, y là một biến int bình thường.

Gán giá trị cho y là 0

Cho x trỏ tới địa chỉ bộ nhớ của y

Lúc này x sẽ mang giá trị là địa chỉ bộ nhớ của y, chúng ta có thể dùng dấu & để lấy địa chỉ bộ nhớ của y, hoặc dùng dấu * để lấy giá trị tại địa chỉ của y (hay giá trị cửa chính biến y)
Gán giá trị cho y (hay giá trị tại địa chỉ bộ nhớ mà x đang tham chiếu tới) là 1 bằng cách dùng dấu *
Output
```
0xc0420044f8
0xc0420044f8
0
0
1
1
```
Phần Output ở máy bạn có thể khác vì địa chỉ bộ nhớ mỗi lần hệ điều hành cung cấp cho biến là khác nhau.

### New

Ngoài cách khai báo biến con trỏ bằng dấu “*”, Go còn cho phép chúng ta khai báo biến con trỏ bằng cách dùng hàm new(). Ví dụ:

new.go
```go
package main
 
import "fmt"
 
func one(xPtr *int) {
    *xPtr = 1
}
 
func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr)
}
```
Tham số mà hàm new() nhận là tên của một kiểu dữ liệu.

Nếu bạn đã từng làm việc với C/C++ thì bạn biết là sau khi khai báo một biến con trỏ, trước khi kết thúc chương trình chúng ta phải dùng lệnh delete để trả biến đó về cho hệ điều hành, nếu không hệ điều hành sẽ “nghĩ” rằng biến đó vẫn còn đang được sử dụng (nhưng thực chất là không) nên sẽ không cấp cho các ứng dụng khác, dẫn đến lãng phí bộ nhớ. Trong Go thì chúng ta không phải làm việc này, Go sẽ tự động dọn dẹp mọi thứ trước khi chương trình kết thúc.