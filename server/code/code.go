package code

import (
	"errors"
)

var (
	HttpOk       = 200
	Unauthorized = 401
	JsonErr      = 402
	TypeErr      = 403 // 数据格式不正确
	//NotFound     = 404 // 404 Not found
	CodeIsHasErr = 405 // 存在重复的编码
	IdNotHas     = 406 // 无效的ID
	SaveErr      = 407 // 保存失败
	UpdateErr    = 408 // 更新失败
	SelectErr    = 409 // 查询失败
	DeleteErr    = 410 // 删除失败
	OptionErr    = 411 // 操作失败
	NameIsEmpty  = 412 // 姓名或名称不能为空
	NameIsHasErr = 413 // 名称已存在
	NameTypeErr  = 414 // 名称不能包含特殊符号
	PhoneErr     = 415 // 手机号不正确
	EmailErr     = 416 // 邮箱不正确
	ServerErr    = 500 // 系统内部错误
)

var ErrMsgMap = make(map[int]error, 100)

func init() {
	ErrMsgMap[JsonErr] = errors.New("请求参数不正确,请检查参数类型是否正确")
	ErrMsgMap[Unauthorized] = errors.New("服务器拒绝处理您的请求！您可能没有访问此操作的权限")
	ErrMsgMap[TypeErr] = errors.New("数据格式不正确")

	//ErrMsgMap[NotFound] = errors.New("")

	ErrMsgMap[CodeIsHasErr] = errors.New("存在重复的编码")
	ErrMsgMap[IdNotHas] = errors.New("无效的ID")
	ErrMsgMap[SaveErr] = errors.New("保存失败")
	ErrMsgMap[UpdateErr] = errors.New("更新失败")
	ErrMsgMap[SelectErr] = errors.New("查询失败")
	ErrMsgMap[DeleteErr] = errors.New("删除失败")
	ErrMsgMap[OptionErr] = errors.New("操作失败")
	ErrMsgMap[NameIsEmpty] = errors.New("姓名或名称不能为空")
	ErrMsgMap[NameIsHasErr] = errors.New("名称已存在")
	ErrMsgMap[NameTypeErr] = errors.New("名称不能包含特殊符号")
	ErrMsgMap[PhoneErr] = errors.New("手机号不正确")
	ErrMsgMap[EmailErr] = errors.New("邮箱不正确")

	ErrMsgMap[ServerErr] = errors.New("系统内部错误")
}
