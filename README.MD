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