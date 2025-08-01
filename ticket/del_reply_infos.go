package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type DelReplyInfosRequest struct {
	core.BaseRequest
	// 工单ID
	TicketId string `json:"ticket_id"`
	// 工单回复ID，删除工单回复信息对应的回复记录ID
	ReplyId string `json:"replyid"`
	// 操作坐席ID
	AgentId string `json:"agentid,omitempty"`
}

type DelReplyInfosResponse struct {
	core.BaseResponse
}

// 工单回复信息删除
// 删除工单的回复信息。
func (t *Ticket) DelReplyInfos(ctx context.Context, req *DelReplyInfosRequest) (*DelReplyInfosResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/del_reply_infos", req)
	if err != nil {
		return nil, err
	}

	var rsp DelReplyInfosResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
