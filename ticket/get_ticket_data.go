package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type GetTicketDataRequest struct {
	core.BaseRequest
	// 开始时间，例如：2019-09-18 00:00:00
	StartDatetime string `json:"start_datetime"`
	// 结束时间，例如：2019-09-18 23:59:59
	EndDatetime string `json:"end_datetime"`
}

type GetTicketDataResponse struct {
	core.BaseResponse
	// 未分配工单数量
	UnassignedTicketNum string `json:"unassigned_ticket_num"`
	// 未解决工单数量
	UnsolvedTicketNum string `json:"unsolved_ticket_num"`
	// 平均首次解决时长
	AvgFirstSolvedTime string `json:"avg_first_solved_time"`
	// 平均首次回复时长
	AvgFirstResponseTime string `json:"avg_first_response_time"`
}

// 查询工单统计概览
// 获取某个时间段工单对应的数量。
func (t *Ticket) GetTicketData(ctx context.Context, req *GetTicketDataRequest) (*GetTicketDataResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/get_ticket_data", nil)
	if err != nil {
		return nil, err
	}

	var rsp GetTicketDataResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
