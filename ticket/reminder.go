package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type ReminderRequest struct {
	core.BaseRequest
	// 工单ID(工单编码和工单ID至少填一个)
	TicketId string `json:"ticket_id,omitempty"`
	// 工单编号（工单ID与工单编号不可同时为空，同时传参以工单ID为准）
	TicketCode string `json:"ticket_code,omitempty"`
	// 催单人（客服ID）
	ReminderAgentid string `json:"reminder_agentid"`
	// 催单原因
	ReminderRemark string `json:"reminder_remark"`
}

type ReminderResponse struct {
	core.BaseResponse
}

func (t *Ticket) Reminder(ctx context.Context, req *ReminderRequest) (*ReminderResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/reminder", req)
	if err != nil {
		return nil, err
	}

	var rsp ReminderResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
