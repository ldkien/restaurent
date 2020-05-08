package error

import pb "restaurant/backend-entity/entities"

const (
	SUCCESS          = 0
	UNAUTHENTICATED  = -1
	SYSTEM_ERROR     = 1
	CONNECTOR_ERROR  = 2
	NOT_USER         = 3
	LOGIN_ERROR      = 4
	INVALID_PARAMS   = 5
	INVALID_USERNAME = 6
	EXIST_USERNAME   = 7
	DB_ERROR         = 8
)

var ErrorDes = map[int32]string{
	0:  "Thành công",
	-1: "Xác thực không hợp lê",
	1:  "Đã có lỗi trong quá trình xử lý. Xin vui lòng thử lại",
	2:  "Lỗi hệ thống. Xin vui lòng thử lại",
	3:  "Người dùng không hợp lệ",
	4:  "Tên đăng nhập hoặc mật khẩu không chính xác",
	5:  "Dữ liệu không hợp lệ",
	6:  "Tên đăng nhập không hợp lệ",
	7:  "Tên đăng nhập đã tồn tại. Xin vui lòng thử lại",
	8:  "Đã có lỗi trong quá trình xử lý. Xin vui lòng thử lại",
	50: "L",
	40: "XL",
	10: "X",
	9:  "IX",
}

func GetError(errorCode int32) *pb.Error {
	return &pb.Error{
		ErrorCode: errorCode,
		ErrorDes:  ErrorDes[errorCode],
	}
}
