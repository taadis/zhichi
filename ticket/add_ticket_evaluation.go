package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type AddTicketEvaluationRequest struct {
	core.BaseRequest
	// 工单ID
	TicketId string `json:"ticketId"`
	// 公司ID
	CompanyId string `json:"companyId"`
	// 评分，1 1星，2 2星，3 3星，4 4星，5 5星
	Score string `json:"score"`
	// 评语
	Remark string `json:"remark,omitempty"`
}

type AddTicketEvaluationResponse struct {
	core.BaseResponse
}

func (t *Ticket) AddTicketEvaluation(ctx context.Context, req *AddTicketEvaluationRequest) (*AddTicketEvaluationResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/add_ticket_evaluation", req)
	if err != nil {
		return nil, err
	}

	var rsp AddTicketEvaluationResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
