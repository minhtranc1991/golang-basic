# Go – List
Trong phần này chúng ta sẽ tìm hiểu về kiểu dữ liệu List trong Go.

Ngoài các kiểu dữ liệu dạng danh sách mà chúng ta đã tìm hiểu như array, slice và map, Go còn cung cấp một kiểu dữ liệu khác là List có trong gói container/list.

### List

Kiểu List Gói container/list thực chất chính là cấu trúc Danh sách liên kết đôi, nếu bạn đã từng học môn Cấu trúc dữ liệu ở Đại học thì sẽ biết cấu trúc này.

![](../capture2.png)

Danh sách liên kết là một danh sách các biến struct kiểu Node do chúng ta tự định nghĩa, các biến Node này có các trường dữ liệu thông thường (như int, float, string...) và một trường đặc biệt là một con trỏ kiểu Node, con trỏ này lưu địa chỉ của Node tiếp theo trong danh sách.

Danh sách liên kết đôi thì mỗi Node chứa 2 con trỏ, một con trỏ chỉ tới Node phía sau nó và một con trỏ chỉ tới Node phía trước nó.

Danh sách liên kết cho phép chúng ta sử dụng bộ nhớ máy tính một cách linh hoạt hơn.

Trong Go thì chúng ta có thể sử dụng struct List đã được định nghĩa sẵn như sau:

list.go
```go
package main
 
import "fmt"
import "container/list"
 
func main() {
    var x list.List
  
    x.PushBack(1)
    x.PushBack(2.5)
    x.PushBack("Hello")
 
    for e := x.Front() ; e != nil ; e = e.Next() {
        fmt.Println(e.Value)
    }
}
```
Trong đoạn code trên, chúng ta tạo một biến List là x, khi tạo thì x rỗng, không có phần tử nào cả.
```go
x.PushBack(1)
x.PushBack(2.5)
x.PushBack("Hello")
```
Chúng ta thêm một phần tử vào List bằng hàm PushBack(), chúng ta có thể đưa vào bất cứ kiểu dữ liệu gì cũng được, từ int, float cho tới string... hoặc một struct do chúng ta tự tạo cũng được.

```go
for e := x.Front() ; e != nil ; e = e.Next() 
```
Để lặp qua một list thì chúng ta có thể dùng vòng lặp for như trong ví dụ, hàm Front() sẽ trả về phần tử đầu tiên của list, hàm Next() sẽ trả về phần tử tiếp theo trong list, chúng ta kiểm tra cho đến khi không còn phần tử nào nữa (nil) thì vòng lặp dừng.

```go
fmt.Println(e.Value)
```
Để lấy giá trị mà phần tử đang lưu trữ thì chúng ta đọc trường Value là được.

Output
```
1
2.5
Hello
```

### Sort

Trong Go có gói sort chứa các hàm hỗ trợ sắp xếp các danh sách với các kiểu dữ liệu thường dùng. Ví dụ:

sorting.go
```go
package main
 
import "fmt"
import "sort"
 
func main() {
 
    x := []int{93, 48, 27, 784, 13}
    fmt.Println(x)
    sort.IntSlice.Sort(x)
    fmt.Println(x)
  
    y := []float64{3.5, 7.6, 8.93, 5.23, 7.609}
    fmt.Println(y)
    sort.Float64Slice.Sort(y)
    fmt.Println(y)
}
```
Trong đoạn code trên chúng ta có 2 slice x và y, slice x chứa các phần tử kiểu int, slice y chứa các phần tử kiểu float64.

```go
sort.IntSlice.Sort(x)
...
sort.Float64Slice.Sort(y)
```
Gói sort có 2 struct là IntSlice và Float64Slice, 2 struct này có hàm Sort()  có chức năng sắp xếp các phần tử trong một slice kiểu int hoặc kiểu float64.

Output
```
[93 48 27 784 13]
[13 27 48 93 784]
[3.5 7.6 8.93 5.23 7.609]
[3.5 5.23 7.6 7.609 8.93]
```
Trong trường hợp chúng ta muốn sắp xếp các kiểu dữ liệu “không thường dùng” như struct mà chúng ta tự định nghĩa, thì chúng ta làm như sau:

sorting2.go
```go
package main
 
import "fmt"
import "sort"
 
type Person struct {
    Name string
    Age int
}
 
type ByName []Person
 
func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
     this[i], this[j] = this[j], this[i]
}
func main() {
    kids := []Person{
        {"Jill", 9},
        {"Jack", 10},
    }
    sort.Sort(ByName(kids))
    fmt.Println(kids)
}
```
Để một slice kiểu struct mới có thể sắp xếp được bằng hàm Sort() thì slice đó phải có 3 phương thức là Len(), Less() và Swap(). Trong đó phương thức Len() trả vể số lượng phần tử của slice, Less() so sánh 2 phần tử i và j cái nào bé hơn, Swap() đổi chỗ 2 phần tử i và j cho nhau. Chẳng hạn như trong đoạn code trên chúng ta định nghĩa một struct có tên Person và một slice kiểu Person có tên ByName. Struct Person có 2 trường là Name và Age, slice ByName có 3 phương thức Len(), Less() và Swap(). Hàm Less() so sánh 2 phần tử bé hơn bằng cách so sánh trường Name của mỗi phần tử đó.

Output
```
[{Jack 10} {Jill 9}]
```
Nếu muốn sắp xếp theo trường Age thì chúng ta cũng làm tương tự với hàm Less():
```go
func (this ByName) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
```