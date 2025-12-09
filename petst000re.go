// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject

import (
	"context"
	"net/http"
	"slices"

	"github.com/jd-st/jd-project-go/internal/requestconfig"
	"github.com/jd-st/jd-project-go/option"
)

// Petst000reService contains methods and other services that help with interacting
// with the jd-project API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetst000reService] method instead.
type Petst000reService struct {
	Options []option.RequestOption
	Orders  Petst000reOrderService
}

// NewPetst000reService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPetst000reService(opts ...option.RequestOption) (r Petst000reService) {
	r = Petst000reService{}
	r.Options = opts
	r.Orders = NewPetst000reOrderService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *Petst000reService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *Petst000reListInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "petst000re/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Petst000reListInventoryResponse map[string]int64
