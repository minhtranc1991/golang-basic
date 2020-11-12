# Go – Hàm băm và mã hóa
Hàm băm (tiếng Anh: Hash) là các hàm làm công việc biến một chuỗi giá trị bình thường thành một chuỗi giá trị khác ngắn hơn và có chiều dài cố định, chuỗi này còn gọi là giá trị băm. Giá trị băm sẽ được dùng để tìm kiếm và đánh chỉ mục trong cơ sở dữ liệu, bởi vì khi chúng ta cần tìm một thứ gì đó trong cơ sở dữ liệu thì việc tìm thông qua các giá trị băm sẽ nhanh hơn nhiều so với tìm chuỗi gốc. Ngoài ra hàm băm còn được dùng nhiều trong các thuật toán mã hóa do đó hàm băm còn được chia làm 2 loại là loại có mã hóa và loại không có mã hóa.

Ví dụ chúng ta có cơ sở dữ liệu lưu trữ tên người gồm: Abernathy, Sara Epperdingle, Roscoe Moore, Wilfred Smith, David, khi chúng ta cần tìm tên của ai đó trong cơ sở dữ liệu, như bình thường thì chúng ta sẽ so sánh chuỗi xem từ khóa tìm có giống với tên người đó hay không, nếu giống thì trả về, không thì tiếp tục tìm cho đến hết. Tuy nhiên thay vì lưu trữ tên gốc như thế, chúng ta có thể dùng hàm băm để chuyển các tên người đó thành các giá trị là số có 4 chữ số như Abernathy – 7864, Sara 9802 Epperdingle – 9802, Roscoe Moore – 1990, Wilfred Smith – 8822 , David – 7822 (tất nhiên trong thực tế thì giá trị này không chỉ có 4 chữ số), và khi tìm kiếm thì thay vì tìm kiếm các chuỗi gốc, chúng ta sẽ tìm kiếm bằng cách tính giá trị băm cho từ khóa rồi so sánh với các giá trị băm trong cơ sở dữ liệu.

### Hàm băm không mã hóa

Trong Go có sẵn một số gói chứa hàm băm không mã hóa là crc32, adler32, crc64 và fnv, các gói này đều nằm trong gói hash. Ví dụ về crc32:

crc32.go
```go
package main
 
import "fmt"
import "hash/crc32"
 
func main() {
  
    str1 := []byte("There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain") 
    hash1 := crc32.NewIEEE()
    hash1.Write(str1)
    fmt.Println(hash1.Sum32())
  
    str2 := []byte("There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain...")
    hash2 := crc32.NewIEEE()
    hash2.Write(str2)
    fmt.Println(hash2.Sum32())
  
    if hash1.Sum32() != hash2.Sum32() {
        fmt.Println("Different")
    } else {
        fmt.Println("Same")
    }
}
```
Trong đoạn code trên, chúng ta tính giá trị băm cho 2 chuỗi bằng cách dùng gói hash/crc32 và so sánh xem 2 chuỗi đó có giống nhau hay không (chuỗi thứ 2 có dấu 3 chấm ở cuối chuỗi).

```go
hash1 := crc32.NewIEEE()
```
Đầu tiên chúng ta dùng hàm crc32.NewIEEE(), hàm này sẽ tạo một đối tượng kiểu hash.Hash32.

```go
hash1.Write(str1)
```
Tiếp theo chúng ta gọi hàm Write() để ghi dữ liệu trong biến str1 vào đối tượng Hash32 đó.

```go
fmt.Println(hash1.Sum32())
```
Cuối cùng hàm Sum32() sẽ tính giá trị băm trả về giá trị đó là một số nguyên 32-bit kiểu uint32.

```go
if hash1.Sum32() != hash2.Sum32() {
    fmt.Println("Different")
} else {
    fmt.Println("Same")
}
```
Và chúng ta chỉ cần so sánh giá trị uint32 này là có thể biết được 2 chuỗi giống nhau hay khác nhau.

Output
```
3455434559
2617155743
Different
```
Nếu bạn muốn biết thuật toán CRC32 hoạt động như thế nào hay muốn tìm hiểu xem tại sao hàm băm này lại có thể tạo ra các chuỗi khác nhau riêng biệt mà không sợ trùng thì tìm hiểu thêm trên internet, ở đây mình không đi sâu.

###Hàm băm có mã hóa

Các hàm băm có mã hóa thì cũng giống hàm băm không mã hóa, chỉ khác là chúng không thể bị đảo ngược, tức là chúng ta không có cách nào để chuyển từ một giá trị băm về chuỗi giá trị gốc được. Do đó các hàm này thường dùng trong các ứng dụng bảo mật (thay vì dùng cho tìm kiếm hoặc kiểm tra xem dữ liệu có thay đổi hay không như CRC32 ở trên).

Một số thuật toán băm có mã hóa thông dụng hiện nay là SHA-1, SHA-256, MD5… Ví dụ về sha-1:

sha1.go
```go
package main
 
import "fmt"
import "crypto/sha1"
 
func main() {
  
    str1 := []byte("There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain") 
    crypto := sha1.New()
    crypto.Write(str1)
    fmt.Println(crypto.Sum([]byte{}))
    fmt.Printf("%x", crypto.Sum([]byte{}))
}
```
Gói crypto trong Go chứa gói sha1 dùng trong việc tính giá trị băm theo thuật toán SHA-1.

Cách sử dụng các hàm băm này cũng tương tự như với hàm băm CRC32, chúng ta dùng hàm sha1.New() để tạo một đối tượng hash.Hash. Sau đó dùng hàm Write() để ghi dữ liệu vào đối tượng Hash này, rồi dùng hàm Sum() để tính giá trị băm.

Một điểm khác của hàm băm SHA-1 với CRC32 là SHA-1 sẽ trả về giá trị 160 bit chứ không phải chỉ có 32 bit, và vì trong Go không có kiểu int nào lớn tới 160 bit nên hàm Sum() sẽ trả về một slice chứa 20 phần tử kiểu byte (mỗi byte 8 bit). Ngoài ra hàm Sum() còn nhận một tham số đầu vào là một slice kiểu byte nữa, ý nghĩa của tham số này là giá trị salt trong thuật toán SHA-1. Thường dùng trong xác thực mật khẩu, ở đây chúng ta để trống, nếu muốn bạn có thể tìm hiểu thêm trên internet.

Ở đoạn code trên chúng ta in giá trị băm dưới dạng số nguyên gốc và số hệ 16.

Output
```
[29 187 202 66 188 186 226 160 230 54 155 114 184 195 234 87 232 6 193 155]
1dbbca42bcbae2a0e6369b72b8c3ea57e806c19b
```