package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type GetTicketByIdRequest struct {
	core.BaseRequest
	// 工单ID
	TicketId string `json:"ticketid,omitempty"`
	// 工单编号（工单ID与工单编号不可同时为空，同时传参以工单ID为准）
	TicketCode string `json:"ticket_code,omitempty"`
}

type GetTicketByIdResponse struct {
	core.BasePageResponse
	Item []*struct {
		// 公司ID
		CompanyId string `json:"companyid"`
		// 工单ID
		Ticket string `json:"ticketid"`
		// 工单标题
		TicketTitle string `json:"ticket_title"`
		// 工单编号
		TicketCode string `json:"ticket_code"`
		// 工单来源
		TicketFrom string `json:"ticket_from"`
		// 工单优先级，0低，1中，2高，3紧急
		TicketLevel string `json:"ticket_level"`
		// 工单状态，0尚未受理，1受理中，2等待回复，3已解决，99已关闭，98已删除，包括自定义工单状态
		TicketStatus string `json:"ticket_status"`
		// 工单发起人类型，0坐席，1客户
		StartType string `json:"start_type"`
		// 工单发起人名称
		StartName string `json:"start_name"`
		// 工单分类名称，显示格式：一级/二级/三级
		TicketTypeName string `json:"ticket_type_name"`
		// 受理坐席名称
		DealAgentName string `json:"deal_agent_name"`
		// 受理技能组名称
		DealGroupName string `json:"deal_group_name"`
		// 工单创建时间，例如：2018-09-18 12:00:00
		CreateDatetime string `json:"create_datetime"`
		// 工单更新时间，例如：2018-09-18 13:00:00
		UpdateDatetime string `json:"update_datetime"`
		// 客户昵称
		UserNick string `json:"user_nick"`
		// 客户名称
		UserName string `json:"user_name,omitempty"`
		// 客户电话,多个电话号码，号码之间采用英文逗号","隔开
		UserTels string `json:"user_tels,omitempty"`
		// 客户邮箱,多个邮箱，邮箱之间采用英文逗号","隔开
		UserEmails string `json:"user_emails,omitempty"`
		// 工单评价时间, 例如：2018-09-18 13:00:00
		EvaluationDatetime string `json:"evaluation_datetime,omitempty"`
		// 评分,1 1星，2 2星，3 3星，4 4星，5 5星
		Score int `json:"score,omitempty"`
		// 评语
		Remark string `json:"remark,omitempty"`
	} `json:"item,omitempty"`
}

// 查询工单详情页
// 返回该条工单的详细信息（包括工单回复、操作历史）。
func (t *Ticket) GetTicketById(ctx context.Context, req *GetTicketByIdRequest) (*GetTicketByIdResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/get_ticket_by_id", nil)
	if err != nil {
		return nil, err
	}

	var rsp GetTicketByIdResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
