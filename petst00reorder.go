// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/jd-st/jd-project-go/internal/encoding/json"
	"github.com/jd-st/jd-project-go/internal/requestconfig"
	"github.com/jd-st/jd-project-go/option"
	"github.com/jd-st/jd-project-go/shared"
)

// Petst00reOrderService contains methods and other services that help with
// interacting with the jd-project API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetst00reOrderService] method instead.
type Petst00reOrderService struct {
	Options []option.RequestOption
}

// NewPetst00reOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPetst00reOrderService(opts ...option.RequestOption) (r Petst00reOrderService) {
	r = Petst00reOrderService{}
	r.Options = opts
	return
}

// Place a new order in the store
func (r *Petst00reOrderService) New(ctx context.Context, body Petst00reOrderNewParams, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "petst00re/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// For valid response try integer IDs with value <= 5 or > 10. Other values will
// generate exceptions.
func (r *Petst00reOrderService) Get(ctx context.Context, orderID int64, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("petst00re/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or
// nonintegers will generate API errors
func (r *Petst00reOrderService) Delete(ctx context.Context, orderID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("petst00re/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Petst00reOrderNewParams struct {
	Order shared.OrderParam
	paramObj
}

func (r Petst00reOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Order)
}
func (r *Petst00reOrderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Order)
}
