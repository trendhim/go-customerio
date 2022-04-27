package customerio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type UpdateCollectionActionRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type UpdateCollectionActionResponse struct {
	Collection struct {
		Bytes     int      `json:"bytes"`
		CreatedAt int      `json:"created_at"`
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		Rows      int      `json:"rows"`
		Schema    []string `json:"schema"`
		UpdatedAt int      `json:"updated_at"`
	} `json:"collection"`
}

type ListCollectionsActionResponse struct {
	Collections []struct {
		Bytes     int      `json:"bytes"`
		CreatedAt int      `json:"created_at"`
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		Rows      int      `json:"rows"`
		Schema    []string `json:"schema"`
		UpdatedAt int      `json:"updated_at"`
	} `json:"collections"`
}

func (c *APIClient) CreateCollectionAction(ctx context.Context, req *UpdateCollectionActionRequest) (*UpdateCollectionActionResponse, error) {
	body, statusCode, err := c.doRequest(ctx, "POST", "/v1/collections", req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var meta struct {
			Errors []struct {
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Status string `json:"status"`
			} `json:"errors"`
		}

		if err := json.Unmarshal(body, &meta); err != nil {
			return nil, &CollectionError{
				StatusCode: statusCode,
				Err:        string(body),
			}
		}

		var firstErrorMessage string
		if len(meta.Errors) > 0 {
			firstErrorMessage = meta.Errors[0].Detail
		}

		return nil, &CollectionError{
			StatusCode: statusCode,
			Err:        firstErrorMessage,
		}
	}

	var result UpdateCollectionActionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) ListCollectionsAction(ctx context.Context) (*ListCollectionsActionResponse, error) {
	body, statusCode, err := c.doRequest(ctx, "GET", "/v1/collections", nil)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var meta struct {
			Errors []struct {
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Status string `json:"status"`
			} `json:"errors"`
		}

		if err := json.Unmarshal(body, &meta); err != nil {
			return nil, &CollectionError{
				StatusCode: statusCode,
				Err:        string(body),
			}
		}

		var firstErrorMessage string
		if len(meta.Errors) > 0 {
			firstErrorMessage = meta.Errors[0].Detail
		}

		return nil, &CollectionError{
			StatusCode: statusCode,
			Err:        firstErrorMessage,
		}
	}

	var result ListCollectionsActionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) UpdateCollectionAction(ctx context.Context, collectionID string, req *UpdateCollectionActionRequest) (*UpdateCollectionActionResponse, error) {
	requestPath := fmt.Sprintf("/v1/collections/%s", url.PathEscape(collectionID))

	body, statusCode, err := c.doRequest(ctx, "PUT", requestPath, req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var meta struct {
			Errors []struct {
				Detail string `json:"detail"`
				Source struct {
					Pointer string `json:"pointer"`
				} `json:"source"`
				Status string `json:"status"`
			} `json:"errors"`
		}

		if err := json.Unmarshal(body, &meta); err != nil {
			return nil, &CollectionError{
				StatusCode: statusCode,
				Err:        string(body),
			}
		}

		var firstErrorMessage string
		if len(meta.Errors) > 0 {
			firstErrorMessage = meta.Errors[0].Detail
		}

		return nil, &CollectionError{
			StatusCode: statusCode,
			Err:        firstErrorMessage,
		}
	}

	var result UpdateCollectionActionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type CollectionError struct {
	// Err is a more specific error message.
	Err string
	// StatusCode is the http status code for the error.
	StatusCode int
}

func (e *CollectionError) Error() string {
	return e.Err
}
