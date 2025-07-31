package zhichi

import (
	"fmt"
)

// see more: https://developer.zhichi.com/pages/950d89/#%E6%8E%A5%E5%8F%A3%E8%BF%94%E5%9B%9E%E7%8A%B6%E6%80%81%E7%BC%96%E7%A0%81

/*
outputs:

curl "https://www.sobot.com/api/get_token?appid=1&create_time=1569397773&sign=258eec3118705112b2c53dc8043d4d34"
{"ret_code":"900004","ret_msg":"没有找到公司的api配置信息!"}

*/

type ErrorCode string

const (
	// 操作成功（除此编码以外的编码为错误编码）
	Success string = "000000"

	//
	ErrCodeVersionNotSupported ErrorCode = "700046"
	ErrCodeTokenEmpty          ErrorCode = "900001"
	ErrCodeTokenInvalid        ErrorCode = "900002"
	ErrCodeSignatureInvalid    ErrorCode = "900003"
	ErrCodeAPIConfigNotFound   ErrorCode = "900004"
	ErrCodeUnknown             ErrorCode = "999999"

	//
	ErrCodeTagGroupNotExist        ErrorCode = "510001"
	ErrCodeTagGroupNameExists      ErrorCode = "510002"
	ErrCodeTagNameExists           ErrorCode = "510003"
	ErrCodeTagParamExistsOrNot     ErrorCode = "510004"
	ErrCodeTagGroupNameEmpty       ErrorCode = "510005"
	ErrCodeTagNameEmpty            ErrorCode = "510006"
	ErrCodeNameTooLong             ErrorCode = "510007"
	ErrCodeParamNameTooLong        ErrorCode = "510008"
	ErrCodeBlacklistTypeRequired   ErrorCode = "500152"
	ErrCodeBlacklistReasonRequired ErrorCode = "500153"
	ErrCodeInvalidBlacklistType    ErrorCode = "500154"
	ErrCodeInvalidTimeRange        ErrorCode = "500155"
	ErrCodeCustomerNotBlacklisted  ErrorCode = "500156"
	ErrCodeCustomerInBlacklisted   ErrorCode = "500157"
)

var (
	//
	ErrVersionNotSupported = NewError(ErrCodeVersionNotSupported, "当前企业所购的版本不支持使用接口（免费版、海外团队版不可使用）")
	ErrTokenEmpty          = NewError(ErrCodeTokenEmpty, "token为空")
	ErrTokenInvalid        = NewError(ErrCodeTokenInvalid, "token已失效，请重新获取")
	ErrSignatureInvalid    = NewError(ErrCodeSignatureInvalid, "signature错误")
	ErrAPIConfigNotFound   = NewError(ErrCodeAPIConfigNotFound, "没有找到公司的api配置信息")
	ErrUnknown             = NewError(ErrCodeUnknown, "系统未知异常")

	//
	ErrTagGroupNotExist        = NewError(ErrCodeTagGroupNotExist, "标签组不存在")
	ErrTagGroupNameExists      = NewError(ErrCodeTagGroupNameExists, "标签组名称已存在")
	ErrTagNameExists           = NewError(ErrCodeTagNameExists, "标签名称已存在")
	ErrTagParamExistsOrNot     = NewError(ErrCodeTagParamExistsOrNot, "标签参数名已存在或标签已不存在")
	ErrTagGroupNameEmpty       = NewError(ErrCodeTagGroupNameEmpty, "标签组名称不能为空")
	ErrTagNameEmpty            = NewError(ErrCodeTagNameEmpty, "标签名称不能为空")
	ErrNameTooLong             = NewError(ErrCodeNameTooLong, "名称不能超过30个字")
	ErrParamNameTooLong        = NewError(ErrCodeParamNameTooLong, "参数名称不能超过30个字")
	ErrBlacklistTypeRequired   = NewError(ErrCodeBlacklistTypeRequired, "拉黑类型为必填字段")
	ErrBlacklistReasonRequired = NewError(ErrCodeBlacklistReasonRequired, "拉黑原因为必填字段")
	ErrInvalidBlacklistType    = NewError(ErrCodeInvalidBlacklistType, "非法的拉黑类型")
	ErrInvalidTimeRange        = NewError(ErrCodeInvalidTimeRange, "时间为秒级时间戳且结束时间大于开始时间")
	ErrCustomerNotBlacklisted  = NewError(ErrCodeCustomerNotBlacklisted, "客户不在黑名单")
	ErrCustomerInBlacklisted   = NewError(ErrCodeCustomerInBlacklisted, "客户在黑名单")
)

type Error struct {
	Code ErrorCode
	//
	Message string
}

// Error implements the std error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

func NewError(errCode ErrorCode, errMsg string) *Error {
	return &Error{
		Code:    errCode,
		Message: errMsg,
	}
}

func Is(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}
	if customErr, ok := err.(*Error); ok {
		return customErr.Code == code
	}
	return false
}
