// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jdproject_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/jd-st/jd-project-go"
	"github.com/jd-st/jd-project-go/internal/testutil"
	"github.com/jd-st/jd-project-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), jdproject.PetNewParams{
		Pet: jdproject.PetParam{
			Name:      "take five",
			PhotoURLs: []string{"take five"},
			ID:        jdproject.Int(1),
			Category: jdproject.CategoryParam{
				ID:   jdproject.Int(1),
				Name: jdproject.String("take five"),
			},
			Status: jdproject.PetStatusAvailable,
			Tags: []jdproject.PetTagParam{{
				ID:   jdproject.Int(1),
				Name: jdproject.String("take five"),
			}},
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), jdproject.PetUpdateParams{
		Pet: jdproject.PetParam{
			Name:      "take five",
			PhotoURLs: []string{"take five"},
			ID:        jdproject.Int(1),
			Category: jdproject.CategoryParam{
				ID:   jdproject.Int(1),
				Name: jdproject.String("take five"),
			},
			Status: jdproject.PetStatusAvailable,
			Tags: []jdproject.PetTagParam{{
				ID:   jdproject.Int(1),
				Name: jdproject.String("take five"),
			}},
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), jdproject.PetFindByStatusParams{
		Status: jdproject.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), jdproject.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		0,
		jdproject.PetUpdateByIDParams{
			Name:   jdproject.String("name"),
			Status: jdproject.String("status"),
		},
	)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		jdproject.PetUploadImageParams{
			AdditionalMetadata: jdproject.String("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *jdproject.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
