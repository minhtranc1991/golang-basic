# Go – Biến
Trong phần này chúng ta sẽ tìm hiểu về biến trong Go.

# Biến

Biến là nơi lưu trữ dữ liệu, một biến gồm có 2 phần là tên biến và kiểu dữ liệu. 

Ví dụ:

variables.go
```go
package main
 
import "fmt"
 
func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
```
Đoạn code chương trình trên in ra dòng chữ Hello World như trong các bài trước, nhưng thay vì chúng ta đưa chuỗi Hello World trực tiếp vào trong hàm Println() thì ở đây chúng ta gán vào một biến có tên là x.

Để khai báo một biến trong Go thì chúng ta dùng từ khóa var, theo sau là tên biến, rồi đến kiểu dữ liệu, cuối cùng chúng ta có thể gán giá trị cho biến hoặc gán sau cũng được. 

Ví dụ:

variables_2.go
```go
package main
 
import "fmt"
 
func main() {
    var x string
    x  = "Hello World"
    fmt.Println(x)
}
```
Biến trong Go nói riêng hay lập trình nói chung thì cũng tương đương như biến mà chúng ta được học trong toán, tuy nhiên có một điểm hơi khác đó là biến trong lập trình thì có thể thay đổi giá trị được, 

ví dụ:

variables_3.go
```go
package main
 
import "fmt"
 
func main() {
    var x string
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)
}
```
Cũng vì thế cách phát biểu về một giá trị của biến cũng khác đi. Chẳng hạn như khi gặp dòng x="Hello World", bạn có thể nói “x bằng Hello World”, tuy nhiên nói đúng hơn là “x nhận giá trị Hello World”, hoặc “x được gán giá trị Hello World”.

Ngoài ra thường trong lập trình chúng ta hay khai báo một biến rồi sau đó gán giá trị cho nó luôn nên Go cho phép chúng ta khai báo và gán giá trị nhanh như sau:

```
x := "Hello World"
```
Ở đây chúng ta không dùng từ khóa var, không khai báo kiểu dữ liệu, thay vào đó chúng ta ghi tên biến rồi dùng toán tử := theo sau là giá trị để khai báo nhanh một biến và gán giá trị ngay tại chỗ. Trình biên dịch Go sẽ tự động nhận diển kiểu dữ liệu dựa vào giá trị mà bạn gán cho biến. Chẳng hạn như ở đây Go thấy giá trị là “Hello World”, tức là một string nên sẽ tự động cho biến x kiểu dữ liệu string. Tóm lại 2 dòng sau đây có chức năng y hệt nhau:
```
var x string = "Hello World"
 
x := "Hello World"
```
### Đặt tên biến

Tên biến có thể có một hoặc nhiều kí tự, có thể chứa chữ cái, dấu gạch dưới _ và chữ số. Kí tự đầu tiên phải là chữ cái hoặc dấu gạch dưới.

Go không quan tâm bạn đặt tên biến như thế nào, nên khi đặt thì chúng ta nên đặt sao cho dễ nhớ và dễ hiểu. 

Ví dụ:
```
title := "Pho Code"
 
fmt.Println("Website name: ", title)
```
### Phạm vi hoạt động của biến

Chúng ta viết lại chương trình Hello World ở trên như sau:

scope.go
```go
package main
 
import "fmt"
 
var x string = "Hello World"
 
func main() {
    fmt.Println(x)
}
```
Ở đây chúng ta đưa dòng khai báo biến x ra ngoài hàm main(). Làm như thế biến x sẽ có thể truy cập bởi bất kì hàm nào. Ví dụ:
```go
var x string = "Hello World"
 
func main() {
    fmt.Println(x)
}
 
func f() {
    fmt.Println(x)
}
```
Trong đoạn code trên hàm f() có thể đọc được giá trị của biến x. Giả sử chúng ta đưa biến x vào lại bên trong hàm main như sau:
```go
func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
 
func f() {
    fmt.Println(x)
}
```
Đoạn code trên khi biên dịch sẽ báo lỗi như sau:

Output
```
main.go:11: undefined: x
```
Dòng báo lỗi trên có nghĩa là biến x không tồn tại. Bởi vì nó chỉ được khai báo ở trong hàm main() nên chỉ có thể đọc được ở trong hàm main(). Chính xác hơn là một biến chỉ có thể được đọc ở trong cặp dấu ngoặc nhọn gần nhất {} theo tài liệu của Go.

### Hằng số

Hằng số đơn giản chỉ là các biến không thể thay đổi dược giá trị. Cách khai báo và gán giá trị cho hằng số cũng giống như với biến, chỉ khác một chỗ là thay từ khóa var bằng từ khóa const. 

Ví dụ:

constants.go
```go
package main
 
import "fmt"
 
func main() {
    const x string = "Hello World"
    fmt.Println(x)
}
```
Chúng ta không thể thay đổi giá trị của hằng số.
```go
const x string = "Hello World"
 
x = "Goodbye World"
```
Đoạn code trên sẽ báo lỗi như sau:

Output
```
main.go:7: cannot assign to x
```
Hằng số thường được dùng để lưu các giá trị dùng nhiều lần mà không cần phải khai báo lại. Trong Go có rất nhiều hằng số được tạo sẵn, ví dụ như hằng Pi trong gói math.

### Khai báo nhiều biến

Ngoài cách khai báo từng biến trên một dòng, bạn có thể khai báo nhiều biến một lúc như sau:
```go
var (
    a = 5
    b = 10
    c = 15
)
```
Chúng ta dùng từ khóa var (hoặc const), theo sau là cặp dấu ngoặc tròn (), rồi tới danh sách các biến và giá trị của chúng.