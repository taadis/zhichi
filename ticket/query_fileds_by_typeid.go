package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type QueryFiledsByTypeIdRequest struct {
	core.BaseRequest
	// 工单分类ID
	TicketTypeId string `json:"ticket_typeid"`
}

type QueryFiledsByTypeIdResponse struct {
	core.BaseResponse
	// 返回集合
	Item []struct {
		// 自定义字段ID
		FieldId string `json:"fieldid"`
		// 自定义字段名称
		FieldName string `json:"field_name"`
		// 自定义字段类型,1单行文本，2多行文本，3日期，4时间，5 数值，6下拉列表，7复选框，8单选框 9 级联 11 地区 12 日期+时间
		FieldType string `json:"field_type"`
		// 自定义字段是否必填,0 否,1 是
		FillFlag string `json:"fill_flag"`
		// 选择型字段的选项信息
		FieldDataList []struct {
			// 自定义字段选项名称
			DataName string `json:"data_name"`
			// 自定义字段选项CODE值
			DataValue string `json:"data_value"`
		} `json:"field_data_list,omitempty"`
	} `json:"item"`
}

// 查询工单分类关联的工单模板
// 返回工单分类关联的工单模板中的自定义字段信息
func (t *Ticket) QueryFiledsByTypeId(ctx context.Context, req *QueryFiledsByTypeIdRequest) (*QueryFiledsByTypeIdResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/query_fileds_by_typeid", nil)
	if err != nil {
		return nil, err
	}

	var rsp QueryFiledsByTypeIdResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
