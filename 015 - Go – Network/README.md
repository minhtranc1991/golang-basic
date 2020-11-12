# Go – Network
Trong phần này chúng ta sẽ tìm hiểu cách lập trình mạng trong Go.

Hiện tại thì có rất nhiều thư viện do các coder bên ngoài viết hỗ trợ lập trình mạng rất tốt. Tuy nhiên chúng ta sẽ sử dụng thư viện có sẵn của Go là gói net vì thư viện mặc định này cũng khá mạnh mẽ rồi.

### TCP Server

Đầu tiên chúng ta sẽ viết một ứng dụng client-server sử dụng giao thức TCP đơn giản. Chúng ta sẽ viết 2 file là server.go và client.go

server.go
```go
package main
 
import (
    "encoding/gob"
    "fmt"
    "net"
)
 
func server() {
    ln, err := net.Listen("tcp", "127.0.0.1:12345")
    if err != nil {
    fmt.Println(err)
        return
    }
    defer ln.Close()
    fmt.Println("Server starts listening on port 12345")
    for {
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        } 
        go handleServerConnection(c)
    }
}
 
func handleServerConnection(c net.Conn) {
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
       fmt.Println(err)
    } else {
       fmt.Println("Received : ", msg)
    }
    c.Close()
}
 
func main() {
    server()
}
```
File server.go sẽ chạy phần server, ở đây chúng ta sử dụng gói net để thực hiện việc lắng nghe các yêu cầu từ client.

```go
ln, err := net.Listen("tcp", "127.0.0.1:12345")
if err != nil {
fmt.Println(err)
    return
}
defer ln.Close()
```
Để bắt đầu lắng nghe thì chúng ta gọi hàm net.Listen(), hàm này nhận tham số là tên giao thức và địa chỉ cổng lắng nghe, ở đây chúng ta dùng giao thức tcp, cổng lắng nghe là 12345, tất nhiên bạn có thể sử dụng các cổng khác tùy ý, hàm này trả về một biến kiểu net.Listener và một biến error. Câu lệnh defer cuối cùng sẽ đảm bảo khi kết thúc chương trình thì cổng cũng sẽ được trả lại cho hệ điều hành.
```go
for {
    c, err := ln.Accept()
    if err != nil {
        fmt.Println(err)
        continue
    } 
    go handleServerConnection(c)
}
```
Sau khi lắng nghe thì chúng ta đi vào một vòng lặp vô hạn, mỗi lần lặp chúng ta gọi hàm ln.Accept(), hàm này sẽ kiểm tra xem có ứng dụng client nào kết nối đến server hay không, nếu không có thì đợi đến khi có và chấp nhận kết nối đó rồi trả về một biến kiểu net.Conn và một biến error. Biến net.Conn lưu các thông tin về client đã kết nối tới server. Nếu quá trình chấp nhận kết nối không có lỗi thì chúng ta xử lý kết nối đó trong hàm handleServerConnection() và hàm này được thực hiện trong một routine riêng biệt với routine của hàm main(), nếu bạn chưa biết routine là gì thì xem bài Concurrency, nếu có lỗi thì chúng ta bỏ qua.
```go
func handleServerConnection(c net.Conn) {
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    ...
}
```
Trong hàm handleServerConnection(), chúng ta đọc dữ liệu truyền từ client đến, ở đây client sẽ gửi những chuỗi tin nhắn đã được mã hóa bằng gói gob (phần code client ở dưới), do đó chúng ta giải mã chuỗi đó bằng hàm gob.NewDecoder() rồi dùng hàm Decode() truyền nội dung đã được giải mã đó vào biến msg để in ra.

client.go
```go
package main
 
import (
    "fmt"
    "net"
    "encoding/gob"
)
 
func client() {
    c, err := net.Dial("tcp", "127.0.0.1:12345")
    if err != nil {
        fmt.Println(err)
        return
    }
  
    msg := "Hello server"
    fmt.Println("Sending : ", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    c.Close()
}
 
func main() {
    client()
}
```
Trong file client.go chúng ta thực hiện kết nối và gửi dữ liệu đến server.

```go
c, err := net.Dial("tcp", "127.0.0.1:12345")
```
Để kết nối đến một server, chúng ta sử dụng hàm net.Dial(), hàm này nhận vào 2 tham số là tên giao thức và địa chỉ IP cùng số cổng. Giá trị trả về là một biến net.Conn và một biến error.
```go
msg := "Hello server"
...
err := gob.NewEncoder(c).Encode(msg)
```
Để mã hóa và gửi dữ liệu lên server thì ở đây chúng ta dùng hàm gob.NewEncoder().Encode()

Output
```
go run server.go
Server starts listening on port 12345
Received : Hello server
```
Output
```
go run client.go
Sending : Hello server
```

### HTTP Server

Chúng ta còn có thể viết cả một webserver với gói net/http. Ví dụ:

webserver.go
```go
package main
 
import (
    "net/http"
    "io"
)
 
func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    content := `<!DOCTYPE html>
                <html>
                    <head>
                        <title>Sample Go Web Server</title>
                    </head>
                    <body>
                        <h1>It Worked!</h1>
                    </body>
                </html>`
    io.WriteString(
        res,
        content,
    )
}
 
func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe("127.0.0.1:12345", nil)
}
```
Trong đoạn code trên, chúng ta viết một webserver.

```go
http.HandleFunc("/", hello)
http.ListenAndServe("127.0.0.1:12345", nil)
```
Hàm http.HandleFunc() cho biết khi nhập đường dẫn URL đến địa chỉ của server thì trả về nội dung trong hàm hello().

Hàm http.ListenAndServe() cho server chạy trên cổng 12345.

```go
func hello(res http.ResponseWriter, req *http.Request) {
    ...
}
```
Bất cứ hàm nào làm công việc trả về nội dung HTML cho client đều phải có 2 tham số đầu vào là một biến http.ResponseWriter và một biến *http.Request. Trong đó biến http.ResponseWriter sẽ được dùng để gửi trả dữ liệu về, *http.Request là biến chứa thông tin về client.

```go
res.Header().Set(
    "Content-Type",
    "text/html",
)
```
Đầu tiên chúng ta cho biết dữ liệu trả về sẽ là nội dung HTML, bằng cách thiết lập trong hàm res.Header().Set()

```go
content := `<!DOCTYPE html>
           ...
            </html>`
io.WriteString(
    res,
    content,
)
```
Tiếp theo chúng ta ghi nội dung HTML trong biến content (chuỗi được bọc trong cặp dầu huyền `` có thể ghi trên nhiều dòng), cuối cùng để chuyển dữ liệu đó về client thì chúng ta có thể dùng hàm io.WriteString(), hàm này nhận vào đối tượng sẽ được ghi dữ liệu (là res) và dữ liệu sẽ được ghi (là content).

![](../Capture%20(1).jpg)