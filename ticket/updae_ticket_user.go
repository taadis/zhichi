package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type UpdateTicketUserRequest struct {
	core.BaseRequest
	// 工单ID(工单编码和工单ID至少填一个)
	TicketId string `json:"ticket_id,omitempty"`
	// 工单编码(工单编码和工单ID至少填一个)
	TicketCode string `json:"ticket_code,omitempty"`
	// 客户ID
	UserId string `json:"user_id"`
}

type UpdateTicketUserResponse struct {
	core.BaseResponse
}

func (t *Ticket) UpdateTicketUser(ctx context.Context, req *UpdateTicketUserRequest) (*UpdateTicketUserResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/update_ticket_user", req)
	if err != nil {
		return nil, err
	}

	var rsp UpdateTicketUserResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
