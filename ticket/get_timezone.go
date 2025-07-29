package ticket

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTimezoneRequest struct {
}

func (r *GetTimezoneRequest) String() string {
	if r == nil {
		return ""
	}
	bs, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(bs)
}

type GetTimezoneResponse struct {
	// application name.
	Name string `json:"name"`

	// application description.
	Description string `json:"description"`

	// application tags.
	Tags []string `json:"tags"`
}

func (r *GetTimezoneResponse) String() string {
	if r == nil {
		return ""
	}
	bs, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(bs)
}

func (t *Ticket) GetTimezone(ctx context.Context, req *GetTimezoneRequest) (*GetTimezoneResponse, error) {
	httpReq, err := t.httpClient.NewRawRequest(ctx, http.MethodGet, "/api/ws/5/ticket/get_timezone", nil)
	if err != nil {
		return nil, err
	}

	var rsp GetTimezoneResponse
	err = t.httpClient.SendJSONRequest(httpReq, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
