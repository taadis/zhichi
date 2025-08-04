package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type QueryTicketStatusInfosRequest struct {
	core.BaseRequest
	// 开始时间，例如：2019-09-18 00:00:00
	StartDatetime string `json:"start_datetime"`
	// 结束时间，例如：2019-09-18 23:59:59
	EndDatetime string `json:"end_datetime"`
}

type QueryTicketStatusInfosResponse struct {
	core.BaseResponse
	// 返回集合
	Item []struct {
		// 状态id，如果未开启自定义状态管理，该字段为空
		StatusId string `json:"status_id"`
		// 状态编码，工单状态属性的具体值
		StatusCode int `json:"status_code"`
		// 自定义工单状态名称
		StatusName string `json:"status_name"`
		// 状态类型 1-新建，2-受理中 3-等待客户回复 4-暂停 5-已解决 6-已关闭 7-已删除
		StatusType int `json:"status_type"`
		// 自定义工单状态描述
		StatusDesc string `json:"status_desc"`
		// 自定义工单状态在客户侧展示的状态名称
		CustomerStatusName string `json:"customer_status_name"`
		// 是否删除 1-是 0-否
		IsDelete int `json:"is_delete"`
	} `json:"item"`
}

// 查询自定义工单状态
// 返回自定义工单状态的基本信息
func (t *Ticket) QueryTicketStatusInfos(ctx context.Context, req *QueryTicketStatusInfosRequest) (*QueryTicketStatusInfosResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/query_ticket_status_infos", nil)
	if err != nil {
		return nil, err
	}

	var rsp QueryTicketStatusInfosResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
