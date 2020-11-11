# golang-basic
Trong phần này chúng ta sẽ tìm hiểu về tính năng xử lý các công việc song song – Concurrency.

Những chương trình lớn đều được xây dựng nên từ những chương trình nhỏ hơn. Chẳng hạn như một webserver sẽ phải tiếp nhận và xử lý các yêu cầu từ browser rồi gửi trả nội dung đó về webserver, thì trong đó mỗi yêu cầu từ browser có thể coi như là một chương trình nhỏ.

Những công việc nhỏ như thế nên được thực hiện song song với nhau, khái niệm này được gọi là Concurrency, Go đưa ra 2 tính năng hỗ trợ concurrency rất mạnh đó là Goroutine và Channel.

### Goroutine

Goroutine là một hàm có thể chạy đồng thời với các hàm khác. Để một hàm chạy theo kiểu goroutine thì chúng ta thêm từ khóa go vào trước lời gọi hàm, ví dụ:

goroutine.go
```go
package main
 
import "fmt"
 
func f(n int) {
    for i := 0 ; i < 10 ; i++ {
        fmt.Println(n, ":", i)
    }
}
 
func main() {
    go f(0)
    var input string 
    fmt.Scanln(&input)
}
```
Trong đoạn code trên có 2 hàm goroutine, hàm đầu tiên là hàm main(), hàm thứ hai là hàm f() khi được gọi trong câu lệnh go f(0). Nếu như chúng ta gọi hàm f() một cách bình thường thì khi gọi, hàm main() sẽ phải dừng tất cả mọi thứ lại, đợi cho hàm f() thực hiện công việc của nó xong rồi trả lại quyền điều khiển cho hàm main() thì hàm main() mới tiếp tục công việc của nó.

Trong đoạn code trên chúng ta không gọi hàm f() như bình thường mà chúng ta chuyển lời gọi đó thành một goroutine, như thế sau khi gọi, hàm main() vẫn tiếp tục công việc của nó, hàm f() cũng thực hiện công việc của nó một cách song song với hàm main().

Các goroutine rất nhẹ, chúng ta có thể tạo ra cả ngàn goroutine cũng được. Ví dụ:

goroutine2.go

package main
```go
import "fmt"
 
func f(n int) {
    for i := 0 ; i < 10 ; i++ {
        fmt.Println(n, ":", i)
    }
}
 
func main() {
    for i := 0 ; i < 10 ; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
```
Đoạn code trên sẽ cho kết quả tương tự như sau:

Output
```
1 : 0
6 : 0
9 : 0
9 : 1
9 : 2
9 : 3
9 : 4
9 : 5
...
```
Nếu của bạn không giống như vậy mà nó hiển thị các số có thứ tự thì chẳng qua là do CPU chạy nhanh quá, thành ra các goroutine được gọi trước đã chạy xong rồi nên chúng không chạy đồng thời. Bạn thử chạy đoạn code trên nhiều lần sẽ thấy mỗi lần chạy kết quả sẽ khác nhau.

###Channel

Channel là tính năng cho phép 2 goroutine giao tiếp/trao đổi dữ liệu với nhau. Ví dụ:

channel.go
```go
package main
 
import (
    "fmt"
    "tme"
)
 
func pinger(c chan string) {
    for i := 0 ; ; i++ {
        c <- "ping"
    }
}
 
func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
 
func main() {
    var c chan string = make(chan string)
    go pinger(c)
    go printer(c)
    var input string
    fmt.Scanln(&input)
}
```
Đoạn code trên sẽ in dòng chữ “ping” vô số lần cho đến khi có người bấm nút Enter. Trong đó chúng ta dùng 2 hàm goroutine là pinger() và printer() và 1 channel là c. Về cơ bản thì goroutine là các luồng chương trình chạy xuyên suốt, channel có thể coi như là các “ống” truyền dữ liệu qua lại giữa các luồng chương trình đó.

![](../Untitled%20(1).png)

Một channel trong Go chỉ là một biến bình thường, chỉ khác là các goroutine có thể đọc được dữ liệu cũng như ghi dữ liệu vào đó, để khai báo một channel thì chúng ta thêm từ khóa chan vào giữa tên biến và tên kiểu dữ liệu, để khởi tạo một biến channel thì chúng ta dùng hàm make(chan <kiểu dữ liệu>).

Chúng ta dùng dấu <- để đưa dữ liệu vào channel, dấu -> để lấy dữ liệu từ channel. Chẳng hạn c <- "ping" nghĩa là đưa chuỗi “ping” vào channel c, msg := <- c nghĩa là lấy dữ liệu trong channel c gán vào biến msg. 

Việc sử dụng channel cho phép đồng bộ hóa dữ liệu giữa các goroutine bởi vì khi một goroutine truyền dữ liệu vào channel goroutine đó sẽ dừng chương trình của nó và đợi đến khi có một goroutine khác lấy dữ liệu ra khỏi channel rồi thì nó mới tiếp tục. Ví dụ:

channel2.go
```go
package main
 
import (
    "fmt"
    "time"
)
 
func pinger(c chan string) {
    for i := 0 ; ; i++ {
        c <- "ping"
    }
}
 
func ponger(c chan string) {
    for i := 0 ; ; i++ {
        c <- "pong"
    }
}
 
func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
 
func main() {
    var c chan string = make(chan string)
 
    go pinger(c)
    go ponger(c)
    go printer(c)
 
    var input string
    fmt.Scanln(&input)
}
```
Đoạn code trên sẽ in các chuỗi “ping” và “pong” liên tiếp nhau.

Output
```
ping
pong
ping
pong
ping
pong
...
```
###Điều hướng channel

Chúng ta có thể quy định channel chỉ được phép đọc hoặc chỉ được phép ghi dữ liệu vào đó. Ví dụ:

```go
func pinger(c chan<- string)
```
Dòng code trên quy định channel c chỉ được truyền dữ liệu vào.

```go
func printer(c <-chan string)
```
Dòng code trên quy định channel c chỉ được phép đọc dữ liệu.

Nếu không quy định hướng đi của channel thì mặc định channel sẽ được phép vừa đọc vừa ghi.

###Lệnh Select

Lệnh select trong Go có chức năng giống hệt như lệnh switch, chỉ khác là select thì được dùng với biến channel. Ví dụ:

select.go
```go
package main
 
import (
	"fmt"
	"time"
)
 
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
 
    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()
    go func() {
        for {
            c2 <- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()
    go func() {
        for {
            select {
            case msg1 := <- c1:
                fmt.Println(msg1)
            case msg2 := <- c2:
                fmt.Println(msg2)
            }
        }
    }()
     
    var input string
    fmt.Scanln(&input)
}
```
Đoạn code trên cứ sau 2 giây sẽ in chuỗi “from 1” và cứ sau 3 giây thì in chuỗi “from 2”. Lệnh select sẽ chọn những channel có dữ liệu để xử lý, nếu không có channel nào có dữ liệu thì chương trình sẽ tạm dừng để “đợi” cho đến khi có một channel có dữ liệu.

Cũng giống như lệnh switch, lệnh channel cũng có trường hợp default, ví dụ:

```go
select {
    case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
    case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
    default:
        fmt.Println("nothing ready")
}
```
Trong trường hợp này nếu không có channel nào có dữ liệu thì câu lệnh sau default sẽ được chạy.

###Buffered Channel

Như đã nói ở trên, các goroutine khi truyền dữ liệu vào channel thì phải có một goroutine khác lấy dữ liệu ra hoặc ngược lại, nếu không các goroutine sẽ đi vào trạng thái “chờ”.

Tuy nhiên chúng ta có thể cho phép goroutine không chờ nữa bằng cách dùng các buffered channel. Buffered Channel tức là channel đó giới hạn dữ liệu vào, ví dụ:

buffered_channel.go
```go
package main
 
import "fmt"
 
func main() {
    msg := make(chan string, 2)
 
    msg <- "buffered"
    msg <- "channel"
 
    fmt.Println(msg)
    fmt.Println(msg)
}
```
Trong đoạn code trên chúng ta tạo ra một channel có tên msg, channel này chỉ được phép nhận vào 2 chuỗi. Và vì đây là một buffered channel nên chúng ta có thể truyền dữ liệu vào mà không cần đợi một goroutine khác lấy dữ liệu ra.

Output
```
buffered
channel
```