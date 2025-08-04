package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type SaveOptionDataValueRequest struct {
	core.BaseRequest
	// 工单自定义字段ID
	FieldId string `json:"fieldid"`
	// 工单自定义字段选项信息
	FieldDataValue []struct {
		// 字段选项值
		FieldValue string `json:"field_value"`
		// 字段文本值
		FieldText string `json:"field_text"`
	} `json:"field_data_value"`
	// 操作坐席ID
	AgentId string `json:"agentid"`
	// 操作坐席名称
	AgentName string `json:"agent_name"`
}

type SaveOptionDataValueResponse struct {
	core.BaseResponse
}

// 工单自定义字段（选择型）添加选项信息
// 工单选择型自定义字段批量添加选项信息
func (t *Ticket) SaveOptionDataValue(ctx context.Context, req *SaveOptionDataValueRequest) (*SaveOptionDataValueResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodPost, "/api/ws/5/ticket/save_option_data_value", req)
	if err != nil {
		return nil, err
	}

	var rsp SaveOptionDataValueResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
