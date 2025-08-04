package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type SaveAgentTicketRequest struct {
	core.BaseRequest
	// 企业ID
	CompanyId string `json:"companyid"`
	// 工单标题
	TicketTitle string `json:"ticket_title"`
	// 客户ID
	UserId string `json:"userid,omitempty"`
	// 对接ID
	PartnerId string `json:"partnerid,omitempty"`
	// 工单问题描述
	TicketContent string `json:"ticket_content"`
	// 客户邮箱
	UserEmail string `json:"user_email,omitempty"`
	// 客户电话
	UserTels string `json:"user_tels,omitempty"`
	// 工单分类ID，叶子节点的分类ID
	TicketTypeId string `json:"ticket_typeid"`
	// 工单来源
	TicketFrom string `json:"ticket_from"`
	// 附件路径，多个附件，附件之间采用英文分号";"隔开
	FileStr string `json:"file_str,omitempty"`
	// 工单自定义字段信息
	ExtendFields []struct {
		// 自定义字段ID
		FieldId string `json:"fieldid"`
		// 自定义字段值
		// 级联字段field_value只需填写最子级节点的data_value；
		// 地区字段按照 省/市/县（区）/街道 码值填写，
		// 示例：10002,10048,10501,14865，
		// 详情请见**查询地区型字段信息**接口
		FieldValue string `json:"field_value"`
	} `json:"extend_fields,omitempty"`
	// 优先联系方式的类型 1:邮箱 2: 手机号
	FirstContactType int `json:"first_contact_type,omitempty"`
	// 优先联系方式的值
	FirstContactInfo string `json:"first_contact_info,omitempty"`
}

type SaveAgentTicketResponse struct {
	core.BaseResponse
}

// 创建工单（坐席）
// 坐席创建工单并把工单关联到相对应的客户下
func (t *Ticket) SaveAgentTicket(ctx context.Context, req *SaveAgentTicketRequest) (*SaveAgentTicketResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/save_agent_ticket", req)
	if err != nil {
		return nil, err
	}

	var rsp SaveAgentTicketResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
