// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject_test

import (
	"context"
	"os"
	"testing"

	"github.com/jd-st/jd-project-go"
	"github.com/jd-st/jd-project-go/internal/testutil"
	"github.com/jd-st/jd-project-go/option"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Petstore.Orders.New(context.TODO(), jdproject.PetstoreOrderNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", order.ID)
}
