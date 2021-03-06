# golang-basic
Go là một ngôn ngữ lập trình được thiết kế dựa trên tư duy lập trình hệ thống. Go được phát triển bởi Robert Griesemer, Rob Pike và Ken Thompson tại Google vào năm 2007. Điểm mạnh của Go là bộ thu gom rác và hỗ trợ lập trình đồng thời (tương tự như đa luồng – multithreading). Go là một ngôn ngữ biên dịch như C/C++, Java, Pascal… Go được giới thiệu vào năm 2009 và được sử dụng hầu hết trong các sản phẩm của Google.

Một số đặc điểm
- Hỗ trợ khai báo kiểu dữ liệu động
- Tốc độ biên dịch nhanh
- Hỗ trợ các tác vụ đồng thời
- Ngôn ngữ đơn giản, ngắn gọn

Tuy nhiên chính vì muốn ngôn ngữ này trở nên cực kỳ đơn giản mà các nhà phát triển đã loại bỏ một số tính năng (mà mình cho là hữu ích) có trong các ngôn ngữ khác như:

- Không hỗ trợ thừa kế
- Không hỗ trợ quá tải toán tử hoặc ghi đè phương thức
- Không hỗ trợ thao tác trên con trỏ (vì lý do bảo mật)
- Không hỗ trợ kiểu Generic (giống như template trong C++)

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