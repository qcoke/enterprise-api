package response

import "github.com/gogf/gf/net/ghttp"

// JsonResponse 系统封装的 JSON 结构
type JsonResponse struct {
	Code    int         `json:"code"`    //错误码（0：失败，1：成功，> 1 各种系统错误码)
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"date"`    // 返回数据（业务具体结构数据定义）
}

// Json 标准返回结果数据结构封装
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}
