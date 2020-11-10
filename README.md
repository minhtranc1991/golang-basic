# golang-basic
Go là một ngôn ngữ lập trình được thiết kế dựa trên tư duy lập trình hệ thống. Go được phát triển bởi Robert Griesemer, Rob Pike và Ken Thompson tại Google vào năm 2007. Điểm mạnh của Go là bộ thu gom rác và hỗ trợ lập trình đồng thời (tương tự như đa luồng – multithreading). Go là một ngôn ngữ biên dịch như C/C++, Java, Pascal… Go được giới thiệu vào năm 2009 và được sử dụng hầu hết trong các sản phẩm của Google.

Một số đặc điểm
Hỗ trợ khai báo kiểu dữ liệu động
Tốc độ biên dịch nhanh
Hỗ trợ các tác vụ đồng thời
Ngôn ngữ đơn giản, ngắn gọn
Tuy nhiên chính vì muốn ngôn ngữ này trở nên cực kỳ đơn giản mà các nhà phát triển đã loại bỏ một số tính năng (mà mình cho là hữu ích) có trong các ngôn ngữ khác như:

Không hỗ trợ thừa kế
Không hỗ trợ quá tải toán tử hoặc ghi đè phương thức
Không hỗ trợ thao tác trên con trỏ (vì lý do bảo mật)
Không hỗ trợ kiểu Generic (giống như template trong C++)
Go có trang chủ tại địa chỉ golang.org.

Cài đặt Go
Bạn download Go tại địa chỉ: https://golang.org/dl/

Hiện tại có 3 phiên bản dành cho 3 dòng hệ điều hành là Windows, Linux và MacOS và bộ mã nguồn. Ứng với mỗi dòng hệ điều hành lại có bản cho 32-bit và 64-bit. Bạn có thể chọn download file ZIP hoặc trình Installer. Bạn download phiên bản phù hợp về rồi cài đặt/giải nén vào nơi mà mình thích.

Có lưu ý là đối với trình Installer thì mặc định Go sẽ được cài vào đường dẫn C:/go và trình Installer sẽ tự động thêm đường dẫn C:/go/bin vào biến môi trường PATH. Nếu bạn chọn cách giải nén bằng file ZIP hoặc dùng Installer mà cài vào đường dẫn khác C:/go thì bạn phải thiết lập biến môi trường này bằng tay, ngoài ra bạn còn phải tạo thêm một biến môi trường khác là GOROOT và cho biến này trỏ tới thư mục mà bạn đã cài đặt Go.

Xem phiên bản Go
Sau khi đã cài đặt Go, bạn có thể mở Command Prompt lên và gõ lệnh go version để xem phiên bản Go đã cài đặt, bạn có thể phải restart lại máy nếu Command Prompt báo không tìm thấy lệnh go.

1 C:/User/PhoCode>go version
2 go version go1.7 windows/amd64

Hiện tại phiên bản mà mình sử dụng để viết series này là phiên bản 1.7.

Ngoài ra bạn có thể xem danh sách tham số khác bằng cách gõ lệnh go help.

Trong phần này chúng ta sẽ học cách sử dụng Go bằng cách viết một chương trình quen thuộc là chương trình Hello World.

Ví dụ

package main
 
import "fmt"
 
// this is a comment 
 
func main() {
    fmt.Println("Hello World")
}

Chúng ta tạo một file text có tên main.go với nội dung như trên.

Sau đó để dịch và chạy chương trình thì bạn vào Command Prompt (cmd) rồi gõ lệnh go run <tên file>. Bạn có thể phải chỉ ra cả đường dẫn đến tên file nếu thư mục hiện tại trong cmd không trùng với thư mục chứa file code. Chẳng hạn file source bạn để trong thư mục C:/Go/main.go thì thư mục hiện tại trong cmd phải là C:/Go.

C:\Go>go run main.go
Hello World

Nếu bạn làm đúng các bước trên thì sau khi gõ lệnh màn hình console sẽ in ra dòng chữ Hello World. Nếu không phải thì tức là bạn đã thực hiện sai ở bước nào đó, có thể là gõ sai code, trình biên dịch cũng sẽ thông báo lỗi ngay cho bạn. Và cũng giống như bất cứ trình biên dịch khác, chỉ cần trong code của bạn có lỗi thì cả chương trình sẽ không thể chạy được.

Giải thích
Một đoạn code chương trình Go được thực thi từ trên xuống dưới và từ trái sang phải giống như đọc một cuốn sách vậy.

package main

Dòng đầu tiên là package main, đây là câu lệnh khai báo “gói”. Tất cả mọi chương trình Go đều phải được bắt đầu bởi câu lệnh khai báo gói. Tính năng Gói có tác dụng giúp chúng ta tổ chức code và tái sử dụng code dễ dàng hơn, chúng ta sẽ tìm hiểu thêm về gói trong các bài sau.

Khi biên dịch code Go thì có 2 loại là biên dịch thành chương trình chạy trực tiếp (executable) và biên dịch thành thư viện (library). Chương trình chạy trực tiếp là các file khả thi (có đuôi .exe trong Windows) có thể chạy một cách trực tiếp trong terminal (Command Prompt trong Windows). Còn thư viện là tập hợp code được gom lại với nhau và có thể được sử dụng trong các chương trình khác, chúng ta sẽ tìm hiểu về thư viện sau. Hiện tại chỉ bạn chỉ cần biết là cần phải có câu lệnh khai báo package trong code của mình.

Sau dòng khai báo package là một dòng trống, giống như các ngôn ngữ khác, trình biên dịch không quan tâm các các khoảng trống này, chúng ta chỉ dùng chúng để đọc code cho dễ hơn thôi.

import "fmt"

Từ khóa import có nghĩa là chúng ta yêu cầu được sử dụng code từ các gói khác trong chương trình của chúng ta. Ở đây là gói fmt (viết tắt của format), gói này chủ yếu chứa code thực hiện việc định dạng dữ liệu ra/vào.

Khi dùng từ khóa import thì tên gói được đặt trong cặp dấu nháy kép. Những kí tự được bọc trong cặp dấu nháy kép đều được gọi chung là chuỗi (hoặc string trong tiếng Anh), chúng ta sẽ tìm hiểu về string sau.

// this is a comment

Ký tự // cho biết những kí tự đứng sau nó là câu bình luận (comment). Các câu bình luận sẽ không được biên dịch. Trong Go có 2 loại comment, // là loại comment trên một dòng, tất cả các kí tự phía sau // sẽ không được biên dịch, và /* */ là loại comment có thể sử dụng trên nhiều dòng, tất cả các kí tự nằm trong cặp dấu /* */ sẽ không được biên dịch.

func main() {
    fmt.Println("Hello World");
}

Tiếp theo là phần khai báo hàm. Hàm là các thành phần xây dựng nên một chương trình. Hàm nhận dữ liệu vào, xử lý dữ liệu và xuất dữ liệu ra. Tất cả các hàm trong Go đều được định nghĩa bởi từ khóa func, theo sau là tên hàm (ở đây chúng ta định nghĩa hàm có tên main), tiếp theo là cặp dấu (), bên trong cặp dấu này chúng ta có thể khai báo một danh sách các tham số, tiếp theo là kiểu dữ liệu trả về (ở đây chúng ta không khai báo), rồi đến phần thân hàm nằm trong cặp dấu ngoặc nhọn {}, thân hàm chứa các câu lệnh, ở đây chúng ta chỉ có duy nhất một câu lệnh. Chúng ta sẽ tìm hiểu thêm về hàm sau.

Ngoài ra cái tên main là một cái tên đặc biệt, hàm main tự động được hệ điều hành “gọi” đến đầu tiên khi chạy từ file khả thi.


fmt.Println("Hello World");

Bên trong hàm main chúng ta chỉ có một câu lệnh. Câu lệnh này có 3 phần. Câu lệnh này gọi hàm Println() trong gói fmt. Hàm này nhận vào tham số là một string có tên “Hello World”. Hàm Println (viết tắt của print line) thực hiện in chuỗi mà nó nhận được ra màn hình.

Bạn có thể tìm hiểu thêm về hàm Println của gói fmt qua lệnh godoc:


C:\User\PhoCode>godoc fmt Println
use 'godoc cmd/fmt' for documentation on the fmt command
 
func Println(a ...interface{}) (n int, err error)
    Println formats using the default forrmats for its operands and writes to 
    standard output. Spaces are always added between operands and a newline is appended. 
    It returns the number of bytes written and any write error encountered.
Tài liệu của Go được viết rất kỹ nhưng nếu bạn chưa học lập trình bao giờ thì đọc sẽ hơi thấy khó hiểu, ngoài ra tài liệu chủ yếu tiếng Anh là chính.

Trong phần trước chúng ta đã sử dụng kiểu dữ liệu string cho chuỗi Hello World. Kiểu dữ liệu phân nhóm các giá trị có liên quan với nhau, có các thao tác có thể thực hiện trên chúng và cách chúng được lưu trữ.

#Kiểu dữ liệu trong Go là tĩnh, tức là không thể thay đổi được không giống như một số ngôn ngữ như PHP, Javascript… cho phép thay đổi kiểu dữ liệu trong suốt quá trình chạy.

Trong phần này chúng ta sẽ tìm hiểu về một số kiểu dữ liệu cơ bản trong Go là kiểu số nguyên, số thực, kiểu string và kiểu boolean.

Số nguyên

Đây là các giá trị số nguyên giống hoàn toàn như trong toán, tuy nhiên số nguyên trong máy tính có giới hạn (tức không có giá trị nào là vô cùng ∞ cả). Các kiểu số nguyên trong Go là uint8, uint16, uint32, uint64, int8, int16, int32, int64. Các con số 8, 16, 32, 64 có nghĩa là máy tính cần dùng bao nhiêu bit để biểu diễn số nguyên đó. uint tức là unsigned int – là kiểu số nguyên không âm. Bảng dưới đây cho biết giới hạn của từng loại kiểu số nguyên:

KIỂU	GIỚI HẠN
uint8	0 – 255
uint16	0 – 65535
uint32	0 – 4294967295
uint64	0 – 18446744073709551615
int8	-128 – 127
int16	-32768 – 32767
int32	-2147483648 – 2147483647
int64	-9223372036854775808 – 9223372036854775807
Kiểu uint8 còn có tên khác là byte, kiểu int32 có tên khác là rune. 

Đặc biệt trong Go còn có 3 kiểu số nguyên phụ thuộc hệ điều hành là uint, int và uintptr, 3 kiểu dữ liệu này có giới hạn giống như kiến trúc của hệ điều hành mà bạn đang sử dụng. Ví dụ nếu bạn đang dùng Windows 64 bit thì kiểu int sẽ có giới hạn giống như kiểu uint64. Thông thường khi sử dụng số nguyên bạn dùng int là đủ rồi. Ví dụ:

number.go

package main
 
import "fmt"
 
func main() {
    fmt.Println("1 + 1 = ", 1 + 1)
}

Kết quả:

Output
1 + 1 = 2

Số thực

Đây là các giá trị số có phần thập phân, ví dụ 1.234, 123.4… Việc lưu trữ cũng như thao tác với số thực trong máy tính khá phức tạp nên chúng ta cũng không đi sâu vào làm gì. Ở đây chúng ta chỉ có một số lưu ý như sau:

Số thực không bao giờ chính xác một cách tuyệt đối, rất khó để biểu diễn chính xác một số thực. Ví dụ như phép trừ 1.01 – 0.99 sẽ cho ra kết quả là 0.020000000000000018 chứ không phải là 0.02 như bạn vẫn nghĩ.
Cũng giống như số nguyên, số thực trong máy tính cũng có nhiều giới hạn khác nhau.
Trong Go có 2 kiểu số thực là float32 và float64, 2 kiểu số phức là complex64 và complex128. Thông thường để biểu diễn số thực, bạn chỉ cần dùng float64 là đủ.

String
String (chuỗi) là các kí tự được bọc trong cặp dấu nháy kép hoặc nháy đơn được dùng để biểu diễn văn bản. String nằm trong dấu nháy kép có thể sử dụng các kí tự điều khiển đặc biệt như \n là xuống dòng, \t là dấu tab.

Chúng ta có thể thực hiện một số thao tác thường dùng trên String như tính độ dài chuỗi, lấy kí tự tại vị trí nhất định, nối chuỗi. Ví dụ:

string.go

package main
 
import "fmt"
 
func main() {
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
} 

Dấu cách cũng được tính là một kí tự, do đó chiều dài chuỗi “Hello World” là 11.

Các kí tự trong một chuỗi được đánh số thứ tự từ 0. Câu lệnh “Hello World”[1] sẽ cho ra kết quả là kí tự ở vị trí số 2, tuy nhiên nếu bạn chạy đoạn code trên thì kết quả sẽ cho ra là số 101 chứ không phải kí tự ‘e’, là bởi vì 101 là mã ASCII của kí tự ‘e’

Chúng ta có thể dùng phép toán + lên 2 chuỗi, kết quả là một chuỗi mới được nối từ 2 chuỗi đầu.

Output
11
101
Hello World

Boolean

Các giá trị boolean là các giá trị 1 bit, có 2 giá trị là true và false được dùng để biểu diễn ý nghĩa ĐÚNG hoặc SAI. Trong Go có 3 phép toán có thể thao tác với giá trị boolean là && (phép AND), || (phép OR) và ! (phép NOT). 

Bảng dưới đây mô tả cách thực hiện phép toán với kiểu boolean:

Phép AND:

BIỂU THỨC	GIÁ TRỊ
true && true	true
true && false	false
false && true	false
false && false	false

Phép OR:

BIỂU THỨC	GIÁ TRỊ
true || true	true
true || false	true
false || true	true
false || false	false

Phép NOT:

BIỂU THỨC	GIÁ TRỊ
!true	false
!false	true

Ví dụ:

boolean.go

package main
 
import "fmt"
 
func main() {
    fmt.Println(true && false)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}

Kết quả:

Output

true
false
true
true
false