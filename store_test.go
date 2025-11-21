// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/jd-project-go"
	"github.com/stainless-sdks/jd-project-go/internal/testutil"
	"github.com/stainless-sdks/jd-project-go/option"
)

func TestStoreListInventory(t *testing.T) {
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
	_, err := client.Store.ListInventory(context.TODO())
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
