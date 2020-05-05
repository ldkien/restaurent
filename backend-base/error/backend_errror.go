package error

import pb "restaurant/backend-entity/entities"

const (
	SUCCESS         = 0
	UNAUTHENTICATED = -1
	SYSTEM_ERROR    = 1
	CONNECTOR_ERROR = 2
	NOT_USER        = 3
	LOGIN_ERROR     = 4
)

var ErrorDes = map[int32]string{
	0:   "Thành công",
	-1:  "Xác thực không hợp lê",
	1:   "Đã có lỗi trong quá trình xử lý. Xin vui lòng thử lại",
	2:   "Lỗi hệ thống. Xin vui lòng thử lại",
	3:   "Người dùng không hợp lệ",
	4:   "Tên đăng nhập hoặc mật khẩu không chính xác",
	500: "D",
	400: "CD",
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	5:   "V",
}

func GetError(errorCode int32) *pb.Error {
	return &pb.Error{
		ErrorCode: errorCode,
		ErrorDes:  ErrorDes[errorCode],
	}
}
