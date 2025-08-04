package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type QueryTicketPushFailInfosRequest struct {
	core.BaseRequest
	// 开始时间，例如：2020-04-08 00:00:00
	StartDatetime string `json:"start_datetime"`
	// 结束时间，例如：2020-04-08 23:59:59
	EndDatetime string `json:"end_datetime"`
	// 当前页数， 默认第一页
	PageNo int `json:"page_no"`
	// 当前页显示条数，默认每页显示15条，最大每页显示50条
	PageSize int `json:"page_size"`
}

type QueryTicketPushFailInfosResponse struct {
	core.BasePageResponse
	// 返回集合
	Item []struct {
		// 工单ID
		TicketId string `json:"ticketid"`
		// 公司ID
		CompanyId string `json:"companyid"`
		// 工单标题
		TicketTitle string `json:"ticket_title"`
		// 问题描述
		TicketContent string `json:"ticket_content"`
		// 工单编号
		TicketCode string `json:"ticket_code"`
		// 工单来源
		TicketFrom string `json:"ticket_from"`
		// 工单优先级， 0低，1中，2高，3紧急
		TicketLevel string `json:"ticket_level"`
		// 工单状态，0尚未受理，1受理中，2等待回复，3已解决，99已关闭，98已删除，包括自定义工单状态
		TicketStatus string `json:"ticket_status"`
		// 工单分类名称，显示格式：一级/二级/三级
		TicketTypeName string `json:"ticket_type_name"`
		// 受理坐席ID
		DealAgentId string `json:"deal_agentid,omitempty"`
		// 受理坐席名称
		DealAgentName string `json:"deal_agent_name,omitempty"`
		// 受理技能组ID
		DealGroupId string `json:"deal_groupid,omitempty"`
		// 受理技能组名称
		DealGroupName string `json:"deal_group_name,omitempty"`
		// 工单回复内容
		ReplyContent string `json:"reply_content,omitempty"`
		// 工单更新时间，例如：2020-04-08 13:00:00
		UpdateDatetime string `json:"update_datetime"`
		// 操作坐席名称
		UpdateAgentName string `json:"update_agent_name"`
		// 工单关联客户ID
		UserId string `json:"userid"`
		// 客户昵称
		UserNick string `json:"user_nick"`
		// 客户名称
		UserName string `json:"user_name,omitempty"`
		// 客户电话，多个电话号码，号码之间采用英文逗号,隔开
		UserTels string `json:"user_tels,omitempty"`
		// 客户邮箱，多个邮箱，邮箱之间采用英文逗号,隔开
		UserEmails string `json:"user_emails,omitempty"`
		// 工单关联记录ID，工单来源是呼叫中心这个是呼叫记录ID，来源是在线工作台-桌面网站客服提交这个是会话记录ID
		RecordId string `json:"recordid,omitempty"`
		// 附件路径
		FileStr string `json:"file_str,omitempty"`
		// 录音文件
		VoiceUrl string `json:"voice_url,omitempty"`
		// SLA规定的首次响应时间，例如：2020-04-08 13:00:00
		HopeAcceptTime string `json:"hope_accept_time,omitempty"`
		// SLA规定的首次解决时间，例如：2020-04-08 13:00:00
		HopeCompleteTime string `json:"hope_complete_time,omitempty"`
		// 催单信息标识，是否催单信息 0 否 1 是
		IsReminder string `json:"is_reminder"`
		// 催单坐席id
		ReminderAgentId string `json:"reminder_agentid,omitempty"`
		// 催单坐席名称
		ReminderAgentName string `json:"reminder_agent_name,omitempty"`
		// 催单备注
		ReminderRemark string `json:"reminder_remark,omitempty"`
		// 催单时间，具体的时间点，例如：2020-04-20 12:35:30
		ReminderTime string `json:"reminder_time,omitempty"`
		// 关联客户对接id
		PartnerId string `json:"partnerid,omitempty"`
		// 工单自定义字段
		ExtendFieldsList []*struct {
			// 自定义字段ID
			FieldId string `json:"fieldid"`
			// 自定义字段名称
			FieldName string `json:"field_name"`
			// 自定义字段类型，1单行文本，2多行文本，3日期，4时间，5 数值，6下拉列表，7复选框，8单选框 9级联 12 日期+时间
			FieldType string `json:"field_type"`
			// 选择型字段选项文本值
			FieldText string `json:"field_text,omitempty"`
			// 自定义字段值
			FieldValue string `json:"field_value"`
		} `json:"extend_fields_list,omitempty"`
	} `json:"items"`
}

// 查询工单消息推送异常数据
// 返回工单消息推送失败的数据
func (t *Ticket) QueryTicketPushFailInfos(ctx context.Context, req *QueryTicketPushFailInfosRequest) (*QueryTicketPushFailInfosResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/query_ticket_push_fail_infos", nil)
	if err != nil {
		return nil, err
	}

	var rsp QueryTicketPushFailInfosResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
