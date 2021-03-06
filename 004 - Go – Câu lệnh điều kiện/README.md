# Go – Câu lệnh điều kiện
Trong phần này chúng ta sẽ tìm hiểu về các lệnh điều khiển chương trình là for, if và switch.

### Lệnh for

Giả sử chúng ta cần in các con số từ 1 đến 10 ra màn hình, chúng ta có thể ghi 10 câu lệnh fmt.Println() như sau:
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println(1)
    fmt.Println(2)
    fmt.Println(3)
    fmt.Println(4)
    fmt.Println(5)
    fmt.Println(6)
    fmt.Println(7)
    fmt.Println(8)
    fmt.Println(9)
    fmt.Println(10)
}
```
Lệnh for cho phép lặp đi lặp lại các câu lệnh nhiều lần. Ví dụ:

for.go
```go
package main
 
import "fmt"
 
func main() {
    var i int = 1
    for i <= 10 {
        fmt.Println(i)
        i  = i + 1
    }
}
```
Đoạn code trên sẽ in các chữ số từ 1 đến 10 ra màn hình, thay vì phải dùng 10 câu lệnh fmt.Println() thì bây giờ chúng ta chỉ cần dùng câu lệnh for là đủ.

Trong đoạn code trên, chúng ta khai báo biến i và gán giá trị là 1. Sau đó chúng ta sử dụng lệnh for để chạy câu lệnh fmt.Println() 10 lần bằng cách ghi từ khóa for, theo sau là một biểu thức điều kiện i <= 10, rồi tới khối lệnh nằm trong cặp dấu ngoặc nhọn {}. 

Khi câu lệnh for bắt đầu chạy, đầu tiên câu lệnh này kiểm tra xem biến i có giá trị bé hơn 10 hay không, nếu đúng thì thực hiện những câu lệnh bên dưới, rồi quay ngược lại kiểm tra i cho đến khi nào i không còn bé hơn 10 thì dừng lại. Do đó chúng ta đặt lệnh i = i + 1 bên trong vòng lặp, cứ mỗi lần lặp giá trị của biến i sẽ được tăng lên 1 cho đến khi i = 10 thì vòng lặp for sẽ thoát.

Output
```
1
2
3
4
5
6
7
8
9
10
```
Chúng ta có thể đưa câu lệnh khai báo biến và lệnh tăng giá trị của biến ngay trên một dòng như sau:

for2.go
```go
func main() {
    for i := 1; i <= 10 ; i++ {
       fmt.Println(i)
    }
}
```
Hâu hết các ngôn ngữ khác có rất nhiều lệnh lặp như while, do, until, foreach.... nhưng Go chỉ hỗ trợ một lệnh lặp duy nhất là lệnh for.

### Lệnh if

Bây giờ chúng ta thử in các con số từ 1 đến 10 và cho biết số đó là chẵn hay lẻ. Để làm điều này thì chúng ta sẽ cần dùng đến câu lệnh if. Ví dụ:

if.go
```go
package main
 
import "fmt"
 
func main() {
    for i := 1; i <= 10 ; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "chan")
        } else {
            fmt.Println(i, "le")
        }
    }
}
```
Ở đây chúng ta sử dụng câu lệnh if để kiểm tra xem biến i mang giá trị là số chẵn hay lẻ. Chúng ta ghi câu lệnh if, theo sau là một biểu thức điều kiện (tức là kết quả phải là true hoặc false), rồi đến khối lệnh nằm trong cặp dấu ngoặc nhọn {}, sau đó chúng ta có thể có thêm các câu lệnh else if hoặc else.

Nếu biểu thức điều kiện phía sau if là true thì thực hiện câu khối lệnh phía sau nó, ngược lại thì tiếp tục kiểm tra các biểu thức điều kiện tiếp theo nếu có.

Ở đây biểu thức điều kiện là câu lệnh i % 2 == 0, tức là chúng ta kiểm tra xem i có chia hết cho 2 hay không (hay i chia cho 2 không dư), nếu đúng thì i là số chẵn.

Output
```
1 le
2 chan
3 le
4 chan
5 le
6 chan
7 le
8 chan
9 le
10 chan
```
### Lệnh switch

Giả sử chúng ta muốn in các chữ số bằng chữ, chúng ta có thể viết đoạn code như sau:
```go
if i == 0 {
    fmt.Println("Khong")
} else if i == 1 {
    fmt.Println("Mot")
} else if i == 2 {
    fmt.Println("Hai")
} else if i == 3 {
    fmt.Println("Ba")
} else if i == 4 {
    fmt.Println("Bon")
} else if i == 5 {
    fmt.Println("Nam")
}
```
Thay vì dùng câu lệnh if như trên, chúng ta có thể dùng câu lệnh switch như sau:
```go
switch i {
case 0:  fmt.Println("Khong")
case 1:  fmt.Println("Mot")
case 2:  fmt.Println("Hai")
case 3:  fmt.Println("Ba")
case 4:  fmt.Println("Bon")
case 5:  fmt.Println("Nam")
default: fmt.Println("Khong biet")
}
```
Chúng ta dùng từ khóa switch, theo sau là một biểu thức điều kiện, rồi tới một danh sách các từ khóa case, ứng với mỗi từ khóa case là một giá trị nào đó, rồi tới dấu 2 chấm : và các lệnh sẽ được thực hiện.

Ý nghĩa của lệnh switch là, nếu biểu thức điều kiện ở switch cho kết quả trùng với giá trị ở từ khóa case nào thì thực hiện các câu lệnh sau từ khóa case đó. Ngoài ra ở đây chúng ta còn có từ khóa default, có tác dụng thực hiện các câu lệnh nếu giá trị ở switch không trùng với bất kì từ khóa case nào.