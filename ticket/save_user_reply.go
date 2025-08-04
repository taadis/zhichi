package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type SaveUserReplyRequest struct {
	core.BaseRequest
	// 工单ID
	TicketId string `json:"ticketid"`
	// 工单回复内容
	ReplyContent string `json:"reply_content"`
	// 回复附件路径，多个附件，附件之间使用英文分号";"隔开
	FileStr string `json:"file_str,omitempty"`
}

type SaveUserReplyResponse struct {
	core.BaseResponse
}

// 回复工单（客户）
// 客户侧回复工单（客户可回复所有人）
func (t *Ticket) SaveUserReply(ctx context.Context, req *SaveUserReplyRequest) (*SaveUserReplyResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/save_user_reply", req)
	if err != nil {
		return nil, err
	}

	var rsp SaveUserReplyResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
