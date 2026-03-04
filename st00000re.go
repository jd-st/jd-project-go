// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject

import (
	"context"
	"net/http"
	"slices"

	"github.com/jd-st/jd-project-go/internal/requestconfig"
	"github.com/jd-st/jd-project-go/option"
)

// Access to Petstore orders
//
// St00000reService contains methods and other services that help with interacting
// with the jd-project API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSt00000reService] method instead.
type St00000reService struct {
	Options []option.RequestOption
	// Access to Petstore orders
	Orders St00000reOrderService
}

// NewSt00000reService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSt00000reService(opts ...option.RequestOption) (r St00000reService) {
	r = St00000reService{}
	r.Options = opts
	r.Orders = NewSt00000reOrderService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *St00000reService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *St00000reListInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "st00000re/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type St00000reListInventoryResponse map[string]int64
