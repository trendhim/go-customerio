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
		return nil, newCollectionError(statusCode, body)
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
		return nil, newCollectionError(statusCode, body)
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
		return nil, newCollectionError(statusCode, body)
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

// newCollectionError reads a failed response: the first error's detail when the body has one, otherwise the body as
// it came. An answer in another shape (a gateway's, or {"meta":{"error":...}}) used to leave the error without text.
func newCollectionError(statusCode int, body []byte) *CollectionError {
	var meta struct {
		Errors []struct {
			Detail string `json:"detail"`
		} `json:"errors"`
	}

	message := string(body)
	if err := json.Unmarshal(body, &meta); err == nil && len(meta.Errors) > 0 && meta.Errors[0].Detail != "" {
		message = meta.Errors[0].Detail
	}

	return &CollectionError{
		StatusCode: statusCode,
		Err:        message,
	}
}

func (e *CollectionError) Error() string {
	return e.Err
}
