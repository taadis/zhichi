package ticket

import (
	"context"
	"net/http"

	"github.com/taadis/zhichi/core"
)

type QueryAgentJobsRequest struct {
	core.BaseRequest
	// 开始时间，例如：2019-09-18 00:00:00
	StartDatetime string `json:"start_datetime"`
	// 结束时间，例如：2019-09-18 23:59:59
	EndDatetime string `json:"end_datetime"`
}

type QueryAgentJobsResponse struct {
	core.BaseResponse
	Item []struct {
		// 坐席ID
		AgentId string `json:"agentid"`
		// 坐席名称
		AgentName string `json:"agent_name"`
		// 处理工单数量
		DealTicketNum string `json:"deal_ticket_num"`
		// 工单完结率
		TicketEndRate string `json:"ticket_end_rate"`
		// 工单回复数
		ReplyNum string `json:"reply_num"`
	} `json:"item"`
}

// 查询工单坐席工作量
// 获取某个时间段内坐席工作量。
func (t *Ticket) QueryAgentJobs(ctx context.Context, req *QueryAgentJobsRequest) (*QueryAgentJobsResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/query_agent_jobs", nil)
	if err != nil {
		return nil, err
	}

	var rsp QueryAgentJobsResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
