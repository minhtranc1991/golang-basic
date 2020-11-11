# Go – Struct và Interface
Trong phần này chúng ta sẽ tìm hiểu về khái niệm struct và interface trong Go. Đây là các khái niệm trong lập trình hướng đối tượng (OOP), nếu bạn chưa từng làm việc với OOP thì bạn nên tham khảo thêm vì lý thuyết OOP trên mạng.

Giả sử chúng ta có đoạn code tính diện tích hình tròn và diện tích hình chữ nhật từ các điểm trên mặt phẳng như sau:

area.go
```go
package main
 
import (
    "fmt"
    "math"
)
 
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
 
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x2, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}
 
func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r;
}
 
func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5
 
    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
    fmt.Println(circleArea(cx, cy, cr))
}
```
Thoạt nhìn thì bài toán này có vẻ đơn giản, tuy nhiên khi bạn muốn tính diện tích của khoảng vài chục hình vuông/hình tròn, lúc này bản thân việc đặt tên biến đã muốn mệt chứ chưa nói đến việc tính toán, dĩ nhiên bạn có thể dùng array, slice… nhưng giả sử bạn muốn hình dung trong đầu hình nào ở vị trí số mấy trong mảng là cũng khá khó khăn rồi. Do đó bạn cần dùng kiểu struct để có thể quản lý tốt hơn.

### Struct

Một struct là một kiểu dữ liệu đặc biệt, kiểu này chứa biến thuộc các kiểu dữ liệu khác, các biến ở đây thường được gọi là các trường hoặc các thuộc tính… Ví dụ chúng ta định nghĩa struct có tên Circle (hình tròn) và Rectangle (hình chữ nhật) như sau:

struct.go
```go
type Circle struct {
    x float64
    y float64
    z float64
}
 
type Rectangle struct {
    x1 float64
    y1 float64
    x2 float64
    y2 float64
}
```
Để định nghĩa một struct thì chúng ta dùng từ khóa type, từ khóa này báo cho Go biết là chúng ta đang định nghĩa một kiểu dữ liệu mới, theo sau là tên kiểu dữ liệu do chúng ta tự đặt, tiếp theo là từ khóa struct để báo cho Go biết là chúng ta đang định nghĩa một struct, cuối cùng là danh sách các trường của struct này. Mỗi trường chúng ta khai báo gồm tên trường và tên kiểu dữ liệu. Ngoài ra chúng ta có thể khai báo ngắn gọn lại như sau:

```go
type Circle struct {
    x, y, r float64 
}
 
type Rectangle struct {
    x1, y1, x2, y2 float64
}
```
### Khai báo biến kiểu struct

Chúng ta khai báo biến kiểu struct giống như khai báo một biến bình thường:

```go
var c Circle
var r Rectangle
```
Lúc này biến c có kiểu dữ liệu là Circle, các trường của biến c sẽ có kiểu dữ liệu mặc định là 0 với int, 0.0 với float, “” với string, nil với con trỏ, biến r cũng tương tự như vậy.  Hoặc khai báo bằng cách dùng hàm new():

```go
c := new(Circle)
r := new(Rectangle)
```
Nếu muốn khởi tạo và gán giá trị cho các trường luôn thì chúng ta làm như sau:

```go
c := Circle{x: 0, y : 0, r : 5}
r := Rectangle{x1 : 0, y1 : 10, x2 : 0, y2 : 10}
```
Hoặc chúng ta không cần ghi tên trường ra nhưng phải truyền kiểu dữ liệu theo đúng thứ tự đã định nghĩa:

```go
c := Circle{0, 0, 5}
r := Rectangle{0, 0, 10, 10}
```
### Trường

Để thao tác các trường thì chúng ta dùng dấu chấm “.” như sau:
```go
fmt.Println(c.x, c.y, c.z)
c.x = 10
c.y = 5
```
Định nghĩa struct trong một hàm cũng giống như các kiểu dữ liệu khác:

```go
func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}
 
func rectangleArea(r Rectangle) float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
```
Khi gọi hàm chúng ta cũng chỉ truyền tên biến struct vào là được:

```go
c := Circle{0, 0, 5}
r := Rectangle{0, 0, 10, 10}
fmt.Println(circleArea(c))
fmt.Println(rectangleArea(r))
```
Cũng giống như các kiểu dữ liệu khác, khi truyền một struct vào hàm, thực chất Go sẽ sao chép biến đó vào trong tham số của hàm chứ không trực tiếp thao tác với hàm, do đó nếu muốn hàm thực hiện các thao tác trên chính struct được truyền vào thì chúng ta phải truyền con trỏ (hoặc địa chỉ bộ nhớ của biến) vào hàm bằng cách dùng phép toán &:

```go
func circleArea(c *Circle) float64 {
    return math.Pi * c.r*c.r
}
 
func main() {
    c := Circle{0, 0, 5}
    fmt.Println(circleArea(&c))
}
```
### Phương thức
Phương thức là các hàm của riêng một struct, khi dùng thì chỉ có các biến có kiểu struct đó mới gọi được hàm. Ví dụ:

```go
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
 
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
```
Để định nghĩa một phương thức thì ở giữa từ khóa func và tên hàm chúng ta khai báo struct sở hữu phương thức đó theo cú pháp giống như khai báo tham số của hàm. Để gọi một phương thức của một struct thì chúng ta cũng dùng dấu chấm “.” như sau:
```go
fmt.Println(c.area())
fmt.Println(r.area())
```
Một phương thức của một struct có thể đọc giá trị của các trường trong struct đó, do đó dùng struct sẽ làm cho việc code trở nên dễ dàng hơn rất nhiều, chúng ta không còn phải truyền các biến không cần thiết vào hàm nữa, và hay hơn là không phải truyền con trỏ. Ngoài ra phương thức của một struct chỉ có struct đó mới dùng được nên việc đặt tên cũng dễ dàng hơn nhiều, thay vì đặt tên circleArea() và rectangleArea(), chúng ta chỉ cần đặt là area() là đủ.

### Interface

Trong các ví dụ trên, chúng ta đã định nghĩa 2 struct là Rectangle (hình chữ nhật) và Circle (hình tròn), cả 2 struct này đều một phương thức tính diện tích có tên giống nhau là area(). Chúng ta có thể “gộp chung” 2 phương thức đó vào một kiểu dữ liệu khác có tên là Interface:


type Shape interface {
    area() float64
}

Trong ví dụ trên chúng ta định nghĩa một Interface có tên Shape, interface này có một phương thức là area(). Để định nghĩa một interface thì cũng giống như định nghĩa một struct, chúng ta dùng từ khóa type, tiếp đến là tên interface rồi đến từ khóa interface, sau đó là danh sách các phương thức trong cặp dấu ngoặc nhọn {}.

Interface thực ra cũng không hẳn là một kiểu dữ liệu như struct vì interface chỉ chứa các phương thức chứ không chứa các trường, interface cũng không có phần định nghĩa phương thức ở ngoài như các struct, chúng chỉ chứa tên phương thức là hết. Vậy thì việc sử dụng interface có gì hay? Câu trả lời là chúng ta có thể dùng interface để thực hiện tính toán trên nhiều kiểu struct khác nhau mà không quan tâm các struct đó là gì. Ví dụ:

interface.go
```go
package main 
 
import (
    "fmt"
    "math"
)
 
type Shape interface {
    area() float64
}
 
type Circle struct {
    x, y, r float64
}
 
type Rectangle struct {
    x1, y1, x2, y2 float64
}
 
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
 
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
 
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
 
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
    area += s.area()
    }
    return area
}
 
func main() {
    c := Circle{0, 0, 5}
    r := Rectangle{0, 0, 10, 10}
  
    fmt.Println(totalArea(&c, &r))
}
```
Trong ví dụ trên, chúng ta định nghĩa hàm totalArea() có chức năng tính tổng diện tích của bất cứ hình nào, hàm này nhận vào tham số là kiểu Shape, nhưng chúng ta có thể truyền vào kiểu Circle hoặc kiểu Rectangle đều được, nếu chúng ta truyền vào kiểu Circle, khi gọi phương thức area() thì Go sẽ gọi phương thức area() của struct Circle, và ngược lại khi truyền vào Rectangle thì gọi phương thức area() của Rectangle.

Output
```
178.53981633974485
```