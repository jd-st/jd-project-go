// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject

import (
	"context"
	"net/http"
	"slices"

	"github.com/jd-st/jd-project-go/internal/requestconfig"
	"github.com/jd-st/jd-project-go/option"
)

// St0reService contains methods and other services that help with interacting with
// the jd-project API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSt0reService] method instead.
type St0reService struct {
	Options []option.RequestOption
	Orders  St0reOrderService
}

// NewSt0reService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSt0reService(opts ...option.RequestOption) (r St0reService) {
	r = St0reService{}
	r.Options = opts
	r.Orders = NewSt0reOrderService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *St0reService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *St0reListInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "st0re/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type St0reListInventoryResponse map[string]int64
