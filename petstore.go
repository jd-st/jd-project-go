// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject

import (
	"context"
	"net/http"
	"slices"

	"github.com/jd-st/jd-project-go/internal/requestconfig"
	"github.com/jd-st/jd-project-go/option"
)

// PetstoreService contains methods and other services that help with interacting
// with the jd-project API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetstoreService] method instead.
type PetstoreService struct {
	Options []option.RequestOption
	Orders  PetstoreOrderService
}

// NewPetstoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPetstoreService(opts ...option.RequestOption) (r PetstoreService) {
	r = PetstoreService{}
	r.Options = opts
	r.Orders = NewPetstoreOrderService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *PetstoreService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *PetstoreListInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "petstore/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PetstoreListInventoryResponse map[string]int64
