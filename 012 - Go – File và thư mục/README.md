# Go – File và thư mục
Trong phần này chúng ta sẽ học cách đọc và ghi file.

### Đọc file

Chúng ta sẽ viết một đoạn code mở file text có nội dung như sau để đọc:

test.txt
```
Hello World
```
Còn đây là đoạn code mở file để đọc:

open_file.go
```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    defer file.Close()
  
    stat, err := file.Stat()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
  
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
  
    str := string(bs)
    fmt.Println(str)
} 
```
Để mở file thì chúng ta dùng hàm Open() trong gói os, hàm này nhận vào đường dẫn đến file, giá trị trả về là một biến struct tên là File và một biến struct tên là error, biến File sẽ dùng cho việc đọc/ghi file, biến error cho biết hàm Open() thực thi có thành công hay không.

Lưu ý là đường dẫn file có thể là đường dẫn tuyệt đối hoặc tương đối, nếu bạn chạy chương trình bằng lệnh go run thì phải dùng đường dẫn tuyệt đối nếu không sẽ có lỗi không tìm thấy file, nếu bạn build đoạn code ra file .exe bằng lệnh go install thì có thể để đường dẫn tương đối.

```go
if err != nil {
    fmt.Println("Error: ", err)
    return
}
```
Do đó sau khi mở file chúng ta kiểm tra xem quá trình mở có bị lỗi hay không bằng cách kiểm tra xem biến error có rỗng hay không, nếu không rỗng thì tức là đã có lỗi và chúng ta cho thoát chương trình luôn bằng cách dùng lệnh return trong hàm main().

```go
defer file.Close()
```
Câu lệnh defer thực hiện việc đóng file sau khi chương trình trong hàm main() kết thúc, những câu lệnh mang tính chất dọn dẹp tài nguyên như vậy nên được để gần với câu lệnh xin cấp phát tài nguyên để dễ quản lý.

```go
stat, err := file.Stat()
if err != nil {
    fmt.Println("Error: ", err)
    return
}
```
Tiếp theo chúng ta dùng hàm Stat() của biến file để lấy về một biến struct kiểu FileInfo, biến này lưu những thông tin cơ bản của file như tên file, ngày tạo, kích thước… Hàm Stat() cũng trả về thêm biến error nên chúng ta cũng nên kiểm tra trước khi tiếp tục.

```go
bs := make([]byte, stat.Size())
_, err = file.Read(bs)
```
Tiếp theo chúng ta tạo một slice kiểu byte có kích thước lấy từ biến FileInfo để lưu từng byte dữ liệu đọc từ file, việc đọc dữ liệu từ file và ghi vào slice được thực hiện thông qua hàm Read() trong biến File, hàm này nhận tham số là slice mà nó sẽ ghi vào. Hàm Read() cũng có trả về một biến error để chúng ta kiểm tra.

```go
str := string(bs)
fmt.Println(str)
```
Nếu quá trình đọc file không có lỗi, chúng ta có thể in nội dung của slice ra màn hình.

Output
```
Hello World
```
Đoạn code trên có thể viết ngắn lại bằng cách dùng một package khác là io/ioutil như sau:

open_file2.go
```go
package main
 
import (
    "fmt"
    "io/ioutil"
)
 
func main() {
    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    str := string(bs)
    fmt.Println(str)
}
```
Hàm ReadFile() trong gói io/ioutil sẽ thực hiện cả việc mở file, đọc nội dung file rồi ghi vào một slice luôn và trả về cho chúng ta.

### Tạo file

Để tạo file thì chúng ta dùng hàm Create() trong gói os như sau:

create_file.go
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	_, _ = file.WriteString("test")
}
```
Hàm Create() nhận tên file và tạo ra file đó, sau đó chúng ta có thể dùng hàm WriteString() để ghi những chuỗi text vào file đó.

### Xem nội dung thư mục

Chúng ta có thể mở một thư mục bằng hàm Open() trong gói os bằng cách đưa vào đường dẫn thư mục thay vì đường dẫn file, sau đó dùng hàm Readdir() để đọc nội dung thư mục. Ví dụ:

open_folder.go
```go
package main 
 
import (
    "fmt"
    "os"
)
 
func main() {
    dir, err := os.Open(".")
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    defer dir.Close()
    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}
```
Chúng ta dùng dấu chấm “.” để viết ngắn cho đường dẫn thư mục hiện tại.

```go
fileInfos, err := dir.Readdir(-1)
```
Hàm Readdir() nhận vào số lượng thư mục và file mà chúng ta muốn đọc, nếu đưa -1 vào thì đọc tất cả, hàm này sẽ trả về một slice chứa các biến kiểu FileInfo. 

```go
fmt.Println(fi.Name())
```
Chúng ta có thể lấy tên file hoặc thư mục thông qua hàm Name() trong biến FileInfo.

Gói io trong Go chỉ chứa một số ít các hàm, còn lại phần lớn là các interface, trong đó có 2 interface chính là Reader và Writer. Reader chứa các hàm hỗ trợ đọc dữ liệu (nhập), Writer chứa các hàm hỗ trợ ghi dữ liệu (xuất). Hầu hết các hàm trong Go nhận tham số là các biến Reader hoặc Writer này.

Ví dụ gói io có hàm Copy() có chức năng sao chép dữ liệu từ một Reader sang một Writer:

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```
Để đọc hoặc ghi dữ liệu vào một slice []byte hoặc một string thì chúng ta có thể dùng struct Buffer trong gói bytes:

```go
var buf bytes.Buffer
buf.Write([]byte("test"))
```
Biến Buffer có thể không cần phải khởi tạo trước, Buffer hỗ trợ cả Reader và Writer. Từ biến Buffer chúng ta có thể chuyển thành một slice []byte bằng cách dùng hàm Bytes(). Nếu chúng ta chỉ có nhu cầu đọc dữ liệu từ string thì có một hàm khác tiện hơn là strings.NewReader().