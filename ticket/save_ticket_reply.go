package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type SaveTicketReplyRequest struct {
	core.BaseRequest
	// 工单标题
	TicketTitle string `json:"ticket_title"`
	// 工单问题描述
	TicketContent string `json:"ticket_content"`
	// 工单ID
	TicketId string `json:"ticketid"`
	// 获取工单信息时间，例如：2019-09-19 13:00:00 (当前时间)
	GetTicketDatetime string `json:"get_ticket_datetime"`
	// 工单回复内容
	ReplyContent string `json:"reply_content,omitempty"`
	// 工单回复类型，0 所有人可见，1 仅坐席可见
	ReplyType string `json:"reply_type"`
	// 受理技能组ID
	DealGroupId string `json:"deal_groupid,omitempty"`
	// 受理技能组名称
	DealGroupName string `json:"deal_group_name,omitempty"`
	// 受理坐席ID
	DealAgentId string `json:"deal_agentid,omitempty"`
	// 受理坐席名称
	DealAgentName string `json:"deal_agent_name,omitempty"`
	// 回复坐席ID，指当前处理回复的坐席
	ReplyAgentId string `json:"reply_agentid"`
	// 回复坐席名称
	ReplyAgentName string `json:"reply_agent_name"`
	// 工单状态，0尚未受理，1受理中，2等待回复，3已解决，99已关闭，98已删除，包括自定义工单状态
	TicketStatus string `json:"ticket_status"`
	// 工单优先级， 0低，1中，2高，3紧急
	TicketLevel string `json:"ticket_level"`
	// 回复附件路径，多个附件，附件之间使用英文分号";"隔开
	ReplyFileStr string `json:"reply_file_str,omitempty"`
	// 抄送坐席
	CopyAgent []struct {
		// 坐席名称
		AgentName string `json:"agent_name"`
		// 坐席邮箱
		AgentMail string `json:"agent_mail"`
	} `json:"copy_agent,omitempty"`
	// 工单自定义字段信息
	ExtendFields []struct {
		// 自定义字段ID
		FieldId string `json:"fieldid"`
		// 自定义字段值
		FieldValue string `json:"field_value"`
	} `json:"extend_fields,omitempty"`
	TemplateId string `json:"templateid,omitempty"`
}

type SaveTicketReplyResponse struct {
	core.BaseResponse
}

// 回复工单（坐席）
// 坐席侧回复工单（坐席回复分为仅回复坐席和回复所有人）
func (t *Ticket) SaveTicketReply(ctx context.Context, req *SaveTicketReplyRequest) (*SaveTicketReplyResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/save_ticket_reply", req)
	if err != nil {
		return nil, err
	}

	var rsp SaveTicketReplyResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
