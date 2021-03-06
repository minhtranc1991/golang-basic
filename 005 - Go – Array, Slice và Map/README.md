# Go – Array, Slice và Map
Trong phần này chúng ta sẽ tìm hiểu thêm về các kiểu dữ liệu có sẵn trong Go là array, slice và map.

### Array

Array (hay mảng) là một tập hợp các phần tử có cùng kiểu dữ liệu nằm liên tiếp nhau. Chúng ta khai báo một array trong Go như sau:
```go
var x [5]int
```
Trong đó x là tên array, có 5 phần tử có kiểu dữ liệu là int. Giả sử chúng ta có đoạn code chương trình như sau:

arrays1.go
```go
package main
 
import "fmt"
 
func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
}
```
Đoạn code trên sẽ in ra kết quả như sau:

Output
```
[0 0 0 0 100]
```
Dòng x[4] = 100 có nghĩa là gán giá trị 100 cho phần tử thứ 5 trong mảng, các phần tử không được gán giá trị sẽ có giá trị mặc định là 0. Lý do phần tử đó là phần tử thứ 5 chứ không phải thứ 4 là vì các phần tử trong array được đánh số thứ tự từ 0.

Để truy xuất một phần tử của mảng chúng ta cũng sử dụng toán tử []. Ví dụ:

arrays2.go
```go
package main
 
import "fmt"
 
func main() {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83
 
    var total float64 = 0
    for i := 0 ; i < 5 ; i++ {
         total = total + x[i]
    }
    fmt.Println(total / 5)
}
```
Đoạn code trên tính giá trị trung bình của các phần tử trong mảng. Khi chạy chương trình sẽ in kết quả ra 86.6.

Tuy nhiên chương trình này viết “chưa hay”, nếu chúng ta thay đổi số lượng phần tử từ 5 thành 6 hay 7… thì cứ mỗi lần sửa chúng ta phải thay hai biểu thức i<5 và total/5 thành i<6, total/6…v.v Để thuận tiện hơn thì chúng ta có thể sử dụng hàm len(), hàm này sẽ trả về số lượng phần tử có trong mảng:
```go
var total float64 = 0
for i := 0 ; i < len(x) ; i++ {
    total = total + x[i]
}
fmt.Println(total / len(x))
```
Tuy nhiên đoạn code trên sẽ báo lỗi:

Output
```
invalid operation: total / 5
(mismatched types float 64 and int)
```
Lỗi này có nghĩa là biểu thức total / len(x) là không hợp lệ vì total có kiểu dữ liệu float64 còn len(x) lại trả về một kiểu int. Để khắc phục thì chúng ta có thể chuyển đổi kiểu dữ liệu của len(x) sang float64 như sau:

```
fmt.Println(total / float64(len(x)))
```
Ngoài ra Go còn cho phép chúng ta dùng vòng lặp for theo cú pháp rút gọn như sau khi duyệt mảng:
```go
var total float64 = 0
 
for i, value := range x {
    total = total + value
}
fmt.Println(total / float64(len(x)))
```
Trong cú pháp trên, biến i chứa vị trí của phần tử hiện tại mà nó tham chiếu đến trong mỗi lần lặp qua mảng, value chứa dữ liệu của vị trí theo biến i, tiếp theo là từ khóa range cùng với tên mảng được sử dụng.

Đoạn code trên sẽ báo lỗi như sau:

Output
```
i declared and not used
```
Dòng báo lỗi trên có nghĩa là biến i đã được khai báo nhưng không được sử dụng. Trong các ngôn ngữ khác thì hầu như trình biên dịch chỉ cảnh báo chứ không báo lỗi khi chúng ta khai báo biến mà không dùng tới, nhưng trong Go thì khác, vì nhà phát triển muốn code trở nên tối ưu hơn nên chúng ta không được khai báo biến mà không dùng.

Để khắc phục thì chúng ta có thể thay tên biến thành dấu gạch dưới _ như sau:
```go
var total float64 = 0
for _, value := range x {
    total += value
}
fmt.Println(total / float64(len(x)))
```
Trình biên dịch sẽ hiểu rằng đây là một biến “giả” được tạo ra nhưng không được sử dụng.

Ngoài ra chúng ta còn có thể khai báo array nhanh như sau:
```go
x := [5]float64{ 98, 93, 77, 82, 83 }
```
### Slice

Slice cũng là một kiểu dữ liệu dạng tập hợp như array, các phần tử trong slice cũng được đánh chỉ số. Điểm khác biệt giữa slice và array là số phần tử trong slice có thể thay đổi được. Chúng ta khai báo một slice như sau:
```go
var x []float64
```
Dòng trên sẽ tạo một slice có 0 phần tử. Ngoài ra chúng ta có thể tạo một slice bằng hàm make() như sau:
```go
x := make([]float64, 5)
```
Dòng trên sẽ tạo một slice có 5 phần tử.

Chúng ta có thể tạo slice từ một array bằng cách dùng biểu thức [low:high] như sau:

```go
arr := [5]float64{98, 93, 77, 82, 83}
x := arr[0:5]
```
Slice x sẽ có giá trị là các phần tử của mảng arr, trong đó [low:high] tức là lấy các phần tử từ vị trí low đến vị trí high - 1. Chẳng hạn như [0:5] sẽ trả về các phần tử [98, 93, 77, 82, 83], [1: 4] trả về các phần tử [93, 77, 82].

Ngoài ra các vị trí low, high cũng có thể bỏ đi, chẳng hạn như arr[:] sẽ lấy toàn bộ array, arr[0:] sẽ lấy các phần tử từ vị trí 0 đến vị trí cuối cùng, arr[:5] sẽ lấy các phần tử từ 0 đến 4.

Ngoài ra trong còn Go có 2 hàm hỗ trợ làm việc với slice là append() và copy(). Ví dụ:

append_slice.go
```go
package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}
```
Hàm append() trong dòng slice2 := append(slice1, 4, 5) có tác dụng tạo một slice mới có các phần tử giống như slice1 và thêm vào 2 phần tử mang giá trị 4 và 5, tức slice2 sẽ có các phần tử là [1, 2, 3, 4, 5].

Ví dụ đối với hàm copy():

copy_slice.go
```go
package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}
```
Dòng copy(slice2, slice1) sẽ sao chép các phần tử trong slice1 vào slice2, tuy nhiên khi tạo slice2 chỉ có 2 phần tử nên hàm này chỉ sao chép 2 phần tử đầu tiên của slice1 vào slice2, do đó slice2 sẽ có các phần tử là [1, 2].

### Map

Map cũng là kiểu dữ liệu dạng tập hợp nhưng các phần tử trong map không có thứ tự, tức là chúng ta không thể truy xuất các phần tử theo chỉ số như slice với array. Thay vào đó, các phần tử trong map là các cặp khóa-giá trị, trong các ngôn ngữ khác thì chúng còn có cái tên như mảng liên kết, bảng băm, từ điển… Việc truy xuất các phần tử trong map được thực hiện thông qua khóa.

Trong Go chúng ta khai báo một map như sau:
```go
var x map[string]int
```
Chúng ta dùng từ khóa map, sau đó là kiểu dữ liệu của khóa trong cặp dấu ngoặc vuông [], rồi đến kiểu dữ liệu của giá trị. Trong dòng code trên mỗi phần tử trong map x có khóa kiểu string mang giá trị kiểu int.

Hoặc chúng ta có thể tạo một map bằng cách dùng hàm make:
```go
x := make(map[string]int)
```
Việc gán giá trị và truy xuất giá trị trong map cũng giống với array và slice, chỉ khác là thay vì dùng số thì bây giờ chúng ta dùng khóa. Ví dụ:

maps.go
```go
package main
 
import "fmt"
 
func main() {
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])
}
```
Đoạn code trên sẽ in số 10 ra màn hình.

Chúng ta có thể xóa một phần tử trong map bằng hàm delete():

```go
delete(x, "key")
```
Tham số cho hàm delete() gồm có tên map và khóa cần xóa.

Thực chất khi truy xuất một phần tử của map chúng ta nhận được 2 giá trị trả về chứ không chỉ có giá trị của khóa, giá trị thứ 2 là một giá trị kiểu boolean cho biết việc truy xuất có thành công hay không. Nếu chúng ta truy xuất một khóa có tồn tại trong map thì giá trị boolean trả về true, ngược lại trả về false, ví dụ:

false_key.go
```go
package main
 
import "fmt"
 
func main() {
    x := make(map[string]int)
    x["key"] = 10
  
    value, ok := x["key"]
    fmt.Println(value, ok)
   
    value2, ok2 := x["key2"]
    fmt.Println(value2, ok2)
}
```
Trong đoạn code trên, map x không có khóa key2, khi truy xuất giá trị sẽ trả về 0, ngoài ra giá trị boolean sẽ là false. Còn khóa key có tồn tại nên sẽ trả về số 10 và giá trị boolean của nó là true. Ngoài ra ở đây chúng ta còn biết được là Go cho phép chúng ta gán nhiều giá trị vào nhiều biến trên một dòng thông qua dấu phẩy ",".

Output
```
10 true
0 false
```
Giống như array, chúng ta có thể vừa khai báo vừa gán giá trị cho map như sau:
```go
elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be" : "Beryllium",
    "B" : "Boron",
    "C" : "Carbon",
    "N" : "Nitrogen",
    "O" : "Oxygen",
    "F" : "Fluorine",
    "Ne" : "Neon",
}
```
Lưu ý là sau phần tử cuối cùng vẫn có dấu phẩy “,” vì Go muốn giúp chúng ta khi cần comment một dòng thì không phải mất công xóa dấu phẩy ở dòng trước:

```go
"F" : "Fluorine",
//"Ne" : "Neon",
```
Chúng ta còn có thể lồng các map vào nhau, tức là mỗi phần tử của map sẽ là một map khác, ví dụ:

```go
elements := make(map[string]map[string]string) {
        "H" : map[string]string {
        "name" : "Hydrogen",
    },
        "He" : map[string]string {
        "name" : "Helium",
    },
        "Li" : map[string]string {
        "name" : "Lithium",
    },
        "Be" : map[string]string {
        "name" : "Beryllium",
    },
        "B" : map[string]string {
         "name" : "Boron",
    },
        "C" : map[string]string {
        "name" : "Carbon",
    },
        "N" : map[string]string {
        "name" : "Nitrogen",
    },
        "O" : map[string]string {
        "name" : "Oxygen",
    },
        "F" : map[string]string {
        "name" : "Fluorine",
    },
        "Ne" : map[string]string {
        "name" : "Neon",
    }, 
} 
```
Trong đoạn code trên, elements là một map, mỗi phần tử của elements là một map khác.