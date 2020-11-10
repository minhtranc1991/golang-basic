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

-----------------------------------------------------------------------------------------------------------------------

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

----------------------------------------------------------------------------------------------------------------------------------------------------------

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

Trong phần này chúng ta sẽ tìm hiểu về biến trong Go.

---------------------------------------------------------------------------------------
#Biến

Biến là nơi lưu trữ dữ liệu, một biến gồm có 2 phần là tên biến và kiểu dữ liệu. Ví dụ:

variables.go

package main
 
import "fmt"
 
func main() {
    var x string = "Hello World"
    fmt.Println(x)
}

Đoạn code chương trình trên in ra dòng chữ Hello World như trong các bài trước, nhưng thay vì chúng ta đưa chuỗi Hello World trực tiếp vào trong hàm Println() thì ở đây chúng ta gán vào một biến có tên là x.

Để khai báo một biến trong Go thì chúng ta dùng từ khóa var, theo sau là tên biến, rồi đến kiểu dữ liệu, cuối cùng chúng ta có thể gán giá trị cho biến hoặc gán sau cũng được. Ví dụ:

variables_2.go

package main
 
import "fmt"
 
func main() {
    var x string
    x  = "Hello World"
    fmt.Println(x);
}

Biến trong Go nói riêng hay lập trình nói chung thì cũng tương đương như biến mà chúng ta được học trong toán, tuy nhiên có một điểm hơi khác đó là biến trong lập trình thì có thể thay đổi giá trị được, ví dụ:

variables_3.go

package main
 
import "fmt"
 
func main() {
    var x string
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)
}

Cũng vì thế cách phát biểu về một giá trị của biến cũng khác đi. Chẳng hạn như khi gặp dòng x="Hello World", bạn có thể nói “x bằng Hello World”, tuy nhiên nói đúng hơn là “x nhận giá trị Hello World”, hoặc “x được gán giá trị Hello World”.

Ngoài ra thường trong lập trình chúng ta hay khai báo một biến rồi sau đó gán giá trị cho nó luôn nên Go cho phép chúng ta khai báo và gán giá trị nhanh như sau:


x := "Hello World"

Ở đây chúng ta không dùng từ khóa var, không khai báo kiểu dữ liệu, thay vào đó chúng ta ghi tên biến rồi dùng toán tử := theo sau là giá trị để khai báo nhanh một biến và gán giá trị ngay tại chỗ. Trình biên dịch Go sẽ tự động nhận diển kiểu dữ liệu dựa vào giá trị mà bạn gán cho biến. Chẳng hạn như ở đây Go thấy giá trị là “Hello World”, tức là một string nên sẽ tự động cho biến x kiểu dữ liệu string. Tóm lại 2 dòng sau đây có chức năng y hệt nhau:


var x string = "Hello World"
 
x := "Hello World"

Đặt tên biến
Tên biến có thể có một hoặc nhiều kí tự, có thể chứa chữ cái, dấu gạch dưới _ và chữ số. Kí tự đầu tiên phải là chữ cái hoặc dấu gạch dưới.

Go không quan tâm bạn đặt tên biến như thế nào, nên khi đặt thì chúng ta nên đặt sao cho dễ nhớ và dễ hiểu. Ví dụ:


title := "Pho Code"
 
fmt.Println("Website name: ", title)

Phạm vi hoạt động của biến
Chúng ta viết lại chương trình Hello World ở trên như sau:

scope.go

package main
 
import "fmt"
 
var x string = "Hello World"
 
func main() {
    fmt.Println(x)
}

Ở đây chúng ta đưa dòng khai báo biến x ra ngoài hàm main(). Làm như thế biến x sẽ có thể truy cập bởi bất kì hàm nào. Ví dụ:


var x string = "Hello World"
 
func main() {
    fmt.Println(x)
}
 
func f() {
    fmt.Println(x)
}

Trong đoạn code trên hàm f() có thể đọc được giá trị của biến x. Giả sử chúng ta đưa biến x vào lại bên trong hàm main như sau:

func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
 
func f() {
    fmt.Println(x)
}

Đoạn code trên khi biên dịch sẽ báo lỗi như sau:

Output

main.go:11: undefined: x

Dòng báo lỗi trên có nghĩa là biến x không tồn tại. Bởi vì nó chỉ được khai báo ở trong hàm main() nên chỉ có thể đọc được ở trong hàm main(). Chính xác hơn là một biến chỉ có thể được đọc ở trong cặp dấu ngoặc nhọn gần nhất {} theo tài liệu của Go.

Hằng số

Hằng số đơn giản chỉ là các biến không thể thay đổi dược giá trị. Cách khai báo và gán giá trị cho hằng số cũng giống như với biến, chỉ khác một chỗ là thay từ khóa var bằng từ khóa const. Ví dụ:

constants.go

package main
 
import "fmt"
 
func main() {
    const x string = "Hello World"
    fmt.Println(x)
}

Chúng ta không thể thay đổi giá trị của hằng số.


const x string = "Hello World"
x = "Goodbye World"

Đoạn code trên sẽ báo lỗi như sau:

Output

main.go:7: cannot assign to x

Hằng số thường được dùng để lưu các giá trị dùng nhiều lần mà không cần phải khai báo lại. Trong Go có rất nhiều hằng số được tạo sẵn, ví dụ như hằng Pi trong gói math.

Khai báo nhiều biến
Ngoài cách khai báo từng biến trên một dòng, bạn có thể khai báo nhiều biến một lúc như sau:

var (
    a = 5
    b = 10
    c = 15
)

Chúng ta dùng từ khóa var (hoặc const), theo sau là cặp dấu ngoặc tròn (), rồi tới danh sách các biến và giá trị của chúng.

----------------------------------------------------------------------------------------------

Trong phần này chúng ta sẽ tìm hiểu về các lệnh điều khiển chương trình là for, if và switch.

Lệnh for
Giả sử chúng ta cần in các con số từ 1 đến 10 ra màn hình, chúng ta có thể ghi 10 câu lệnh fmt.Println() như sau:

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

Lệnh for cho phép lặp đi lặp lại các câu lệnh nhiều lần. Ví dụ:

for.go

package main
 
import "fmt"
 
func main() {
    var i int = 1
    for i <= 10 {
        fmt.Println(i)
        i  = i + 1
    }
}

Đoạn code trên sẽ in các chữ số từ 1 đến 10 ra màn hình, thay vì phải dùng 10 câu lệnh fmt.Println() thì bây giờ chúng ta chỉ cần dùng câu lệnh for là đủ.

Trong đoạn code trên, chúng ta khai báo biến i và gán giá trị là 1. Sau đó chúng ta sử dụng lệnh for để chạy câu lệnh fmt.Println() 10 lần bằng cách ghi từ khóa for, theo sau là một biểu thức điều kiện i <= 10, rồi tới khối lệnh nằm trong cặp dấu ngoặc nhọn {}. 

Khi câu lệnh for bắt đầu chạy, đầu tiên câu lệnh này kiểm tra xem biến i có giá trị bé hơn 10 hay không, nếu đúng thì thực hiện những câu lệnh bên dưới, rồi quay ngược lại kiểm tra i cho đến khi nào i không còn bé hơn 10 thì dừng lại. Do đó chúng ta đặt lệnh i = i + 1 bên trong vòng lặp, cứ mỗi lần lặp giá trị của biến i sẽ được tăng lên 1 cho đến khi i = 10 thì vòng lặp for sẽ thoát.

Output
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

Chúng ta có thể đưa câu lệnh khai báo biến và lệnh tăng giá trị của biến ngay trên một dòng như sau:

for2.go

func main() {
    for i := 1; i <= 10 ; i++ {
       fmt.Println(i)
    }
}

Hâu hết các ngôn ngữ khác có rất nhiều lệnh lặp như while, do, until, foreach.... nhưng Go chỉ hỗ trợ một lệnh lặp duy nhất là lệnh for.

Lệnh if
Bây giờ chúng ta thử in các con số từ 1 đến 10 và cho biết số đó là chẵn hay lẻ. Để làm điều này thì chúng ta sẽ cần dùng đến câu lệnh if. Ví dụ:

if.go

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

Ở đây chúng ta sử dụng câu lệnh if để kiểm tra xem biến i mang giá trị là số chẵn hay lẻ. Chúng ta ghi câu lệnh if, theo sau là một biểu thức điều kiện (tức là kết quả phải là true hoặc false), rồi đến khối lệnh nằm trong cặp dấu ngoặc nhọn {}, sau đó chúng ta có thể có thêm các câu lệnh else if hoặc else.

Nếu biểu thức điều kiện phía sau if là true thì thực hiện câu khối lệnh phía sau nó, ngược lại thì tiếp tục kiểm tra các biểu thức điều kiện tiếp theo nếu có.

Ở đây biểu thức điều kiện là câu lệnh i % 2 == 0, tức là chúng ta kiểm tra xem i có chia hết cho 2 hay không (hay i chia cho 2 không dư), nếu đúng thì i là số chẵn.

Output
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

Lệnh switch
Giả sử chúng ta muốn in các chữ số bằng chữ, chúng ta có thể viết đoạn code như sau:

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
Thay vì dùng câu lệnh if như trên, chúng ta có thể dùng câu lệnh switch như sau:


switch i {
case 0:  fmt.Println("Khong")
case 1:  fmt.Println("Mot")
case 2:  fmt.Println("Hai")
case 3:  fmt.Println("Ba")
case 4:  fmt.Println("Bon")
case 5:  fmt.Println("Nam")
default: fmt.Println("Khong biet")
}

Chúng ta dùng từ khóa switch, theo sau là một biểu thức điều kiện, rồi tới một danh sách các từ khóa case, ứng với mỗi từ khóa case là một giá trị nào đó, rồi tới dấu 2 chấm : và các lệnh sẽ được thực hiện.

Ý nghĩa của lệnh switch là, nếu biểu thức điều kiện ở switch cho kết quả trùng với giá trị ở từ khóa case nào thì thực hiện các câu lệnh sau từ khóa case đó. Ngoài ra ở đây chúng ta còn có từ khóa default, có tác dụng thực hiện các câu lệnh nếu giá trị ở switch không trùng với bất kì từ khóa case nào.

------------------------------------------------------------------------------------------------------------------------

Trong phần này chúng ta sẽ tìm hiểu thêm về các kiểu dữ liệu có sẵn trong Go là array, slice và map.

Array

Array (hay mảng) là một tập hợp các phần tử có cùng kiểu dữ liệu nằm liên tiếp nhau. Chúng ta khai báo một array trong Go như sau:

var x [5]int

Trong đó x là tên array, có 5 phần tử có kiểu dữ liệu là int. Giả sử chúng ta có đoạn code chương trình như sau:

arrays1.go

package main
 
import "fmt"
 
func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
}

Đoạn code trên sẽ in ra kết quả như sau:

Output

[0 0 0 0 100]

Dòng x[4] = 100 có nghĩa là gán giá trị 100 cho phần tử thứ 5 trong mảng, các phần tử không được gán giá trị sẽ có giá trị mặc định là 0. Lý do phần tử đó là phần tử thứ 5 chứ không phải thứ 4 là vì các phần tử trong array được đánh số thứ tự từ 0.

Để truy xuất một phần tử của mảng chúng ta cũng sử dụng toán tử []. Ví dụ:

arrays2.go

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

Đoạn code trên tính giá trị trung bình của các phần tử trong mảng. Khi chạy chương trình sẽ in kết quả ra 86.6.

Tuy nhiên chương trình này viết “chưa hay”, nếu chúng ta thay đổi số lượng phần tử từ 5 thành 6 hay 7… thì cứ mỗi lần sửa chúng ta phải thay hai biểu thức i<5 và total/5 thành i<6, total/6…v.v Để thuận tiện hơn thì chúng ta có thể sử dụng hàm len(), hàm này sẽ trả về số lượng phần tử có trong mảng:

var total float64 = 0
for i := 0 ; i < len(x) ; i++ {
    total = total + x[i]
}

fmt.Println(total / len(x))
Tuy nhiên đoạn code trên sẽ báo lỗi:

Output

invalid operation: total / 5
(mismatched types float 64 and int)

Lỗi này có nghĩa là biểu thức total / len(x) là không hợp lệ vì total có kiểu dữ liệu float64 còn len(x) lại trả về một kiểu int. Để khắc phục thì chúng ta có thể chuyển đổi kiểu dữ liệu của len(x) sang float64 như sau:

fmt.Println(total / float64(len(x)))

Ngoài ra Go còn cho phép chúng ta dùng vòng lặp for theo cú pháp rút gọn như sau khi duyệt mảng:

var total float64 = 0
 
for i, value := range x {
    total = total + value
}
fmt.Println(total / float64(len(x)))

Trong cú pháp trên, biến i chứa vị trí của phần tử hiện tại mà nó tham chiếu đến trong mỗi lần lặp qua mảng, value chứa dữ liệu của vị trí theo biến i, tiếp theo là từ khóa range cùng với tên mảng được sử dụng.

Đoạn code trên sẽ báo lỗi như sau:

Output

i declared and not used

Dòng báo lỗi trên có nghĩa là biến i đã được khai báo nhưng không được sử dụng. Trong các ngôn ngữ khác thì hầu như trình biên dịch chỉ cảnh báo chứ không báo lỗi khi chúng ta khai báo biến mà không dùng tới, nhưng trong Go thì khác, vì nhà phát triển muốn code trở nên tối ưu hơn nên chúng ta không được khai báo biến mà không dùng.

Để khắc phục thì chúng ta có thể thay tên biến thành dấu gạch dưới _ như sau:

var total float64 = 0
for _, value := range x {
    total += value
}
fmt.Println(total / float64(len(x)))

Trình biên dịch sẽ hiểu rằng đây là một biến “giả” được tạo ra nhưng không được sử dụng.

Ngoài ra chúng ta còn có thể khai báo array nhanh như sau:

x := [5]float64{ 98, 93, 77, 82, 83 }

Slice

Slice cũng là một kiểu dữ liệu dạng tập hợp như array, các phần tử trong slice cũng được đánh chỉ số. Điểm khác biệt giữa slice và array là số phần tử trong slice có thể thay đổi được. Chúng ta khai báo một slice như sau:

var x []float64

Dòng trên sẽ tạo một slice có 0 phần tử. Ngoài ra chúng ta có thể tạo một slice bằng hàm make() như sau:

x := make([]float64, 5)

Dòng trên sẽ tạo một slice có 5 phần tử.

Chúng ta có thể tạo slice từ một array bằng cách dùng biểu thức [low:high] như sau:

arr := [5]float64{98, 93, 77, 82, 83}
x := arr[0:5]

Slice x sẽ có giá trị là các phần tử của mảng arr, trong đó [low:high] tức là lấy các phần tử từ vị trí low đến vị trí high - 1. Chẳng hạn như [0:5] sẽ trả về các phần tử [98, 93, 77, 82, 83], [1: 4] trả về các phần tử [93, 77, 82].

Ngoài ra các vị trí low, high cũng có thể bỏ đi, chẳng hạn như arr[:] sẽ lấy toàn bộ array, arr[0:] sẽ lấy các phần tử từ vị trí 0 đến vị trí cuối cùng, arr[:5] sẽ lấy các phần tử từ 0 đến 4.

Ngoài ra trong còn Go có 2 hàm hỗ trợ làm việc với slice là append() và copy(). Ví dụ:

append_slice.go

package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}

Hàm append() trong dòng slice2 := append(slice1, 4, 5) có tác dụng tạo một slice mới có các phần tử giống như slice1 và thêm vào 2 phần tử mang giá trị 4 và 5, tức slice2 sẽ có các phần tử là [1, 2, 3, 4, 5].

Ví dụ đối với hàm copy():

copy_slice.go

package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}

Dòng copy(slice2, slice1) sẽ sao chép các phần tử trong slice1 vào slice2, tuy nhiên khi tạo slice2 chỉ có 2 phần tử nên hàm này chỉ sao chép 2 phần tử đầu tiên của slice1 vào slice2, do đó slice2 sẽ có các phần tử là [1, 2].

Map

Map cũng là kiểu dữ liệu dạng tập hợp nhưng các phần tử trong map không có thứ tự, tức là chúng ta không thể truy xuất các phần tử theo chỉ số như slice với array. Thay vào đó, các phần tử trong map là các cặp khóa-giá trị, trong các ngôn ngữ khác thì chúng còn có cái tên như mảng liên kết, bảng băm, từ điển… Việc truy xuất các phần tử trong map được thực hiện thông qua khóa.

Trong Go chúng ta khai báo một map như sau:

var x map[string]int

Chúng ta dùng từ khóa map, sau đó là kiểu dữ liệu của khóa trong cặp dấu ngoặc vuông [], rồi đến kiểu dữ liệu của giá trị. Trong dòng code trên mỗi phần tử trong map x có khóa kiểu string mang giá trị kiểu int.

Hoặc chúng ta có thể tạo một map bằng cách dùng hàm make:

x := make(map[string]int)

Việc gán giá trị và truy xuất giá trị trong map cũng giống với array và slice, chỉ khác là thay vì dùng số thì bây giờ chúng ta dùng khóa. Ví dụ:

maps.go

package main
 
import "fmt"
 
func main() {
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])
}

Đoạn code trên sẽ in số 10 ra màn hình.

Chúng ta có thể xóa một phần tử trong map bằng hàm delete():

delete(x, "key")

Tham số cho hàm delete() gồm có tên map và khóa cần xóa.

Thực chất khi truy xuất một phần tử của map chúng ta nhận được 2 giá trị trả về chứ không chỉ có giá trị của khóa, giá trị thứ 2 là một giá trị kiểu boolean cho biết việc truy xuất có thành công hay không. Nếu chúng ta truy xuất một khóa có tồn tại trong map thì giá trị boolean trả về true, ngược lại trả về false, ví dụ:

false_key.go

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

Trong đoạn code trên, map x không có khóa key2, khi truy xuất giá trị sẽ trả về 0, ngoài ra giá trị boolean sẽ là false. Còn khóa key có tồn tại nên sẽ trả về số 10 và giá trị boolean của nó là true. Ngoài ra ở đây chúng ta còn biết được là Go cho phép chúng ta gán nhiều giá trị vào nhiều biến trên một dòng thông qua dấu phẩy ",".

Output

10 true
0 false

Giống như array, chúng ta có thể vừa khai báo vừa gán giá trị cho map như sau:

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

Lưu ý là sau phần tử cuối cùng vẫn có dấu phẩy “,” vì Go muốn giúp chúng ta khi cần comment một dòng thì không phải mất công xóa dấu phẩy ở dòng trước:


"F" : "Fluorine",
//"Ne" : "Neon",

Chúng ta còn có thể lồng các map vào nhau, tức là mỗi phần tử của map sẽ là một map khác, ví dụ:

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

Trong đoạn code trên, elements là một map, mỗi phần tử của elements là một map khác.