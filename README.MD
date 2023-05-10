POST /user/create: Tạo tài khoản người dùng mới.
POST /user/login: Đăng nhập bằng tài khoản người dùng đã có.
GET /user/info/flight: Lấy thông tin về các chuyến bay có sẵn.
POST /user/register/ticket: Đặt vé cho một chuyến bay cụ thể.
PATCH /user/cancel/ticket/:phone_number/:booking_id: Hủy vé đã đặt bằng cách cung cấp số điện thoại và mã đặt chỗ.
GET /user/info/ticket: Lấy thông tin về tất cả các vé đã đặt.
GET /user/info/ticket/phone_number/:phone_number: Lấy thông tin về các vé đã đặt cho một người dùng cụ thể, được xác định bằng số điện thoại của họ.

Các endpoint cho người dùng quản trị:
POST /admin/create: Tạo tài khoản quản trị mới.
POST /admin/login: Đăng nhập bằng tài khoản quản trị đã có.
GET /admin/info/user: Lấy thông tin về tất cả các tài khoản người dùng.
POST /admin/create/flight: Tạo lịch trình chuyến bay mới.
PATCH /admin/update/flight/:flight_id/:name_flight: Cập nhật lịch trình chuyến bay hiện có bằng cách cung cấp ID chuyến bay và tên chuyến bay mới.
DELETE /admin/delete/flight/:flight_id/:name_flight: Xóa lịch trình chuyến bay hiện có bằng cách cung cấp ID chuyến bay và tên chuyến bay.
GET /admin/info/flight: Lấy thông tin về tất cả các chuyến bay có sẵn.




Bảng "users" chứa thông tin về người dùng đăng ký trên hệ thống.

user_id: là khóa chính (primary key) của bảng, định danh người dùng.
user_name: tên người dùng, không được để trống.
phone_number: số điện thoại người dùng, không được để trống.
email: địa chỉ email người dùng, không được để trống và phải là duy nhất.
password: mật khẩu người dùng, không được để trống.
address: địa chỉ người dùng, không được để trống.
date_birth: ngày tháng năm sinh người dùng, không được để trống.
number_cmnd: số CMND/Hộ chiếu của người dùng, không được để trống.
nationality: quốc tịch của người dùng, không được để trống.
language: ngôn ngữ sử dụng trên hệ thống của người dùng, không được để trống.
amount: số tiền trong tài khoản của người dùng, không được để trống và phải lớn hơn hoặc bằng 0.
Bảng "flights" chứa thông tin về các chuyến bay.

flight_id: là khóa chính (primary key) của bảng, định danh chuyến bay.
name_airline: tên hãng hàng không, không được để trống.
destination: điểm đến của chuyến bay, không được để trống.
departure: điểm khởi hành của chuyến bay, không được để trống.
destination_time: thời gian đến của chuyến bay, không được để trống.
departure_time: thời gian khởi hành của chuyến bay, không được để trống.
remaining_seats: số lượng ghế trống còn lại trên chuyến bay.
ticket_type: loại vé (ví dụ: hạng thương gia, hạng phổ thông), không được để trống.
fare: giá tiền của vé, không được để trống và phải lớn hơn hoặc bằng 0.
status: trạng thái của chuyến bay (ví dụ: đã hoàn tất, bị hủy), không được để trống.
name_flight: tên của chuyến bay.
Bảng "bookings" chứa thông tin về các đặt chỗ.

booking_id: là khóa chính (primary key) của bảng, định danh đặt chỗ.
user_id: là khóa ngoại (foreign key) tham chiếu đến bảng "users", định danh người dùng đặt chỗ.
flight_id: là khóa ngoại (foreign key) tham chiếu đến bảng "flights", định danh chuyến bay được đặt chỗ.
number_of_seats: số lượng ghế được đặt chỗ.
amount: tổng số tiền phải trả cho đặt chỗ, không được để trống và phải




User
tiếp tục
ChatGPT
Tiếp tục với bảng "flights":

flight_id: Định danh (ID) của chuyến bay, đây là trường khóa chính của bảng.
name_airline: Tên hãng hàng không cho chuyến bay.
destination: Điểm đến của chuyến bay.
departure: Điểm khởi hành của chuyến bay.
destination_time: Thời gian dự kiến đến điểm đến.
departure_time: Thời gian dự kiến khởi hành.
remaining_seats: Số chỗ ngồi còn lại trong chuyến bay.
ticket_type: Loại vé được bán cho chuyến bay.
fare: Giá vé của chuyến bay.
status: Trạng thái hiện tại của chuyến bay.
name_flight: Tên chuyến bay.
Bảng "bookings":

booking_id: Định danh (ID) của đơn đặt chỗ, đây là trường khóa chính của bảng.
user_id: ID của người dùng đặt vé, đây là trường tham chiếu đến bảng "users".
flight_id: ID của chuyến bay đặt vé, đây là trường tham chiếu đến bảng "flights".
number_of_seats: Số lượng chỗ ngồi được đặt trong đơn đặt chỗ.
amount: Tổng số tiền phải trả cho đơn đặt chỗ.
user_name: Tên người dùng đặt vé.
phone_number: Số điện thoại người dùng đặt vé.
address: Địa chỉ người dùng đặt vé.
number_cmnd: Số chứng minh nhân dân người dùng đặt vé.
nationality: Quốc tịch người dùng đặt vé.
language: Ngôn ngữ được sử dụng bởi người dùng đặt vé.
name_airline: Tên hãng hàng không của chuyến bay được đặt vé.
destination: Điểm đến của chuyến bay được đặt vé.
departure: Điểm khởi hành của chuyến bay được đặt vé.
destination_time: Thời gian dự kiến đến điểm đến của chuyến bay được đặt vé.
departure_time: Thời gian dự kiến khởi hành của chuyến bay được đặt vé.
remaining_seats: Số chỗ ngồi còn lại trong chuyến bay được đặt vé.
ticket_type: Loại vé của chuyến bay được đặt vé.
fare: Giá vé của chuyến bay được đặt vé.
status: Trạng thái hiện tại của chuyến bay được đặt vé.
register_time: Thời gian đăng ký đơn đặt chỗ.
Bảng "payments":

payment_id: Định danh (ID) của thanh toán, đây là trường khóa chính của bảng.
booking_id: