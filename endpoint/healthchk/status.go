package healthchk

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct {
	Log         logging.Logger
	ServiceName string
}

type GetRequest struct {
	Name string
}

func (gl *GetLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, gr *GetRequest) {
	res.Body = map[string]interface{}{"message": "All okay.", "service_name": gl.ServiceName, "query_name": gr.Name}
	res.HTTPStatus = 200
}
