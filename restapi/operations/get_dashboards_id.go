package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDashboardsIDHandlerFunc turns a function with the right signature into a get dashboards ID handler
type GetDashboardsIDHandlerFunc func(context.Context, GetDashboardsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDashboardsIDHandlerFunc) Handle(ctx context.Context, params GetDashboardsIDParams) middleware.Responder {
	return fn(ctx, params)
}

// GetDashboardsIDHandler interface for that can handle valid get dashboards ID params
type GetDashboardsIDHandler interface {
	Handle(context.Context, GetDashboardsIDParams) middleware.Responder
}

// NewGetDashboardsID creates a new http.Handler for the get dashboards ID operation
func NewGetDashboardsID(ctx *middleware.Context, handler GetDashboardsIDHandler) *GetDashboardsID {
	return &GetDashboardsID{Context: ctx, Handler: handler}
}

/*GetDashboardsID swagger:route GET /dashboards/{id} getDashboardsId

Specific pre-configured dashboard containing cells and queries.

dashboards will hold information about how to layout the page of graphs.


*/
type GetDashboardsID struct {
	Context *middleware.Context
	Handler GetDashboardsIDHandler
}

func (o *GetDashboardsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDashboardsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
