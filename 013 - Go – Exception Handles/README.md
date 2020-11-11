# Go – Exception Handling
Trong các bài trước chúng ta đã từng thấy một kiểu biến struct có tên là error. Trong phần này chúng ta sẽ tìm hiểu kỹ hơn.

Nếu bạn đã từng lập trình Java, C#, PHP… thì có lẽ bạn đã quen với khái niệm lỗi ngoại lệ – Exception, lỗi này còn có một tên khác là Runtime Error, nếu chưa thì mình nói luôn là lỗi này là loại lỗi chỉ xuất hiện khi chúng ta chạy chương trình. Ví dụ dễ hiểu nhất khi học về lỗi ngoại lệ là chương trình máy tính bỏ túi, chương trình này rất đơn giản, chỉ là yêu cầu người dùng cung cấp 2 biến số và phép tính, chúng ta tính toán giá trị rồi hiển thị lên cho người dùng, tuy nhiên chương trình này sẽ bị lỗi ngoại lệ ở phép chia có mẫu số là 0. Đây là một lỗi ngoại lệ vì nó chỉ xuất hiện khi người dùng nhập vào số 0 – tức là lúc chương trình đang chạy.

Trong các ngôn ngữ Java, C#, PHP… thì chúng ta xử lý lỗi ngoại lệ bằng cách “bắt” lấy chúng. Còn trong Go thì có hơi khác là gán chúng vào một biến rồi xử lý. Do đó cơ chế này cho phép chúng ta theo dõi lỗi xuất hiện ở dòng code nào dễ hơn.

### Tạo biến error

Trong Go thì khi lỗi ngoại lệ xuất hiện, các hàm sẽ trả về một biến kiểu error chứa thông tin về lỗi đó. Ví dụ:

error.go
```go
package main

import (
	"errors"
	"fmt"
)

func divide(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		return -1, errors.New("Can't divide by 0")
	}
	return arg1 / arg2, nil
}

func main() {
	arg1 := 10
	arg2 := 0
	result, err := divide(arg1, arg2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```
Chúng ta dùng hàm errors.New() để tạo một biến kiểu error. Hàm này nhận vào một chuỗi để mô tả lỗi xảy ra.

```go
func divide(arg1, arg2 int) (int, error) {
    if arg2 == 0 {
       return -1, errors.New("Can't divide by 0")
    }
    return arg1 / arg2, nil
}
```
Trong ví dụ trên, chúng ta viết hàm divide() có chức năng thực hiện phép chia 2 số nguyên, chúng ta kiểm tra nếu mẫu số arg2 là 0 thì trả về -1 và một biến error. Nếu không có lỗi thì trả về kết quả phép chia, còn biến error chúng ta trả về là nil tức là rỗng.

```go
result, err := divide(arg1, arg2)
if err != nil {
    ...
} else {
    ...
}
```
Khi gọi hàm divide(), chúng ta cũng gán 2 giá trị trả về vào 2 biến result và err, sau đó chúng ta kiểm tra xem biến err có khác rỗng hay không, nếu khác rỗng thì tức là có lỗi, chúng ta in lỗi đó ra, nếu không rỗng thì chúng ta in kết quả ra.

Output
```
Can't divide by 0
```
### Tùy chỉnh thông báo lỗi

Nếu như hàm errors.New() chỉ cho phép chúng ta tạo biến error với chuỗi cố định thì chúng ta có thể tùy chỉnh để in chuỗi thông báo lỗi một cách linh hoạt hơn.

error trong Go là một interface có phương thức Error() thực hiện việc in một chuỗi ra màn hình. Chúng ta có thể code lại phương thức đó để tùy chỉnh chuỗi in ra. Ví dụ:

error2.go
```go
package main
 
import "fmt"
 
type fraction struct {
    arg1, arg2 int
}
 
func (e *fraction) Error() string {
    return fmt.Sprintf("%d can't divide by %d", e.arg1, e.arg2)
}
 
func divide(arg1, arg2 int) (int, error) {
    if arg2 == 0 {
        err := fraction{arg1, arg2}
        return -1, &err
    }
    return arg1 / arg2, nil
}
 
func main() {
    arg1 := 10
    arg2 := 0
 
    result, err := divide(arg1, arg2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
```
Ở đây chúng ta định nghĩa một struct có tên fraction để lưu trữ một phân số với tử số là arg1, mẫu số là arg2.

```go
func (e *fraction) Error() string {
    return fmt.Sprintf("%d can't divide by %d", e.arg1, e.arg2)
}
```
Chúng ta code lại phương thức Error() cho struct fraction, ở đây chúng ta dùng hàm fmt.Sprintf() để tạo một chuỗi tùy ý.

```go
func divide(arg1, arg2 int) (int, error) {
    if arg2 == 0 {
        err := fraction{arg1, arg2}
        return -1, &err
    }
    return arg1 / arg2, nil
}
```
Trong hàm divide(), thay vì chúng ta tạo một biến error từ hàm errors.New() thì bây giờ chúng ta tạo một biến kiểu fraction rồi trả về biến đó, lưu ý là chúng ta phải trả về địa chỉ đến biến struct đó.

Output
```
10 can't divide by 0
```
Mặc dù chúng ta định nghĩa kiểu fraction nhưng khi trả về thì hàm divide() sẽ trả về kiểu error. Trong trường hợp chúng ta muốn lấy kiểu trả về là kiểu gốc là fraction thì chúng ta tiến hành ép kiểu như sau:

```go
result, err := divide(10, 0)
frac, ok := err.(*fraction)
if ok {
    fmt.Println(frac.arg1)    // 10
    fmt.Println(frac.arg2)    // 0
}
```