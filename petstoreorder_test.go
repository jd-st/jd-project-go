// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/jd-st/jd-project-go"
	"github.com/jd-st/jd-project-go/internal/testutil"
	"github.com/jd-st/jd-project-go/option"
	"github.com/jd-st/jd-project-go/shared"
)

func TestPetstoreOrderNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jdproject.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Petstore.Orders.New(context.TODO(), jdproject.PetstoreOrderNewParams{
		Order: shared.OrderParam{
			ID:       jdproject.Int(10),
			Complete: jdproject.Bool(true),
			PetID:    jdproject.Int(198772),
			Quantity: jdproject.Int(7),
			ShipDate: jdproject.Time(time.Now()),
			Status:   shared.OrderStatusApproved,
		},
	})
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetstoreOrderGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jdproject.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Petstore.Orders.Get(context.TODO(), 0)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetstoreOrderDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jdproject.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Petstore.Orders.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
