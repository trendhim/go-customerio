package customerio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TestGroup struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Label      string `json:"label"`
	Winner     bool   `json:"winner"`
	ContentIds []int  `json:"content_ids"`
}

type ListNewsletterTestGroupsResponse struct {
	TestGroups []TestGroup `json:"test_groups"`
}

type UpdateNewsletterVariantRequest struct {
	Body          string `json:"body,omitempty"`
	BodyAmp       string `json:"body_amp,omitempty"`
	FromID        int64  `json:"from_id,omitempty"`
	ReplyToID     int64  `json:"reply_to_id,omitempty"`
	Recipient     string `json:"recipient,omitempty"`
	Subject       string `json:"subject,omitempty"`
	PreheaderText string `json:"preheader_text,omitempty"`
	Headers       []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers,omitempty"`
}

type UpdateNewsletterVariantResponse struct {
	ID            int    `json:"id"`
	NewsletterID  int    `json:"newsletter_id"`
	DeduplicateID string `json:"deduplicate_id"`
	Name          string `json:"name"`
	Layout        string `json:"layout"`
	Body          string `json:"body"`
	BodyAmp       string `json:"body_amp"`
	Language      string `json:"language"`
	Type          string `json:"type"`
	From          string `json:"from"`
	FromID        int    `json:"from_id"`
	ReplyTo       string `json:"reply_to"`
	ReplyToID     int    `json:"reply_to_id"`
	Preprocessor  string `json:"preprocessor"`
	Recipient     string `json:"recipient"`
	Subject       string `json:"subject"`
	Bcc           string `json:"bcc"`
	FakeBcc       bool   `json:"fake_bcc"`
	PreheaderText string `json:"preheader_text"`
}

func (t *UpdateNewsletterVariantResponse) UnmarshalJSON(b []byte) error {
	var r struct {
		Content struct {
			UpdateNewsletterVariantResponse
		} `json:"content"`
	}
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	t.ID = r.Content.ID
	t.NewsletterID = r.Content.NewsletterID
	t.DeduplicateID = r.Content.DeduplicateID
	t.Name = r.Content.Name
	t.Layout = r.Content.Layout
	t.Body = r.Content.Body
	t.BodyAmp = r.Content.BodyAmp
	t.Language = r.Content.Language
	t.Type = r.Content.Type
	t.From = r.Content.From
	t.FromID = r.Content.FromID
	t.ReplyTo = r.Content.ReplyTo
	t.ReplyToID = r.Content.ReplyToID
	t.Preprocessor = r.Content.Preprocessor
	t.Recipient = r.Content.Recipient
	t.Subject = r.Content.Subject
	t.Bcc = r.Content.Bcc
	t.FakeBcc = r.Content.FakeBcc
	t.PreheaderText = r.Content.PreheaderText

	return nil
}

func (c *APIClient) UpdateNewsletterVariant(ctx context.Context, newsletterID string, contentID string, req *UpdateNewsletterVariantRequest) (*UpdateNewsletterVariantResponse, error) {
	requestPath := fmt.Sprintf("/v1/newsletters/%s/contents/%s", url.PathEscape(newsletterID), url.PathEscape(contentID))

	body, statusCode, err := c.doRequest(ctx, "PUT", requestPath, req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var tmpError = &DataError{}
		if err := json.Unmarshal(body, &tmpError); err != nil {
			error := Errors{
				Status: statusCode,
				Detail: string(body),
			}
			return nil, &DataError{[]Errors{error}}
		}
		return nil, tmpError
	}

	var result UpdateNewsletterVariantResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) UpdateNewsletterVariantTranslation(ctx context.Context, newsletterID string, language string, req *UpdateNewsletterVariantRequest) (*UpdateNewsletterVariantResponse, error) {
	requestPath := fmt.Sprintf("/v1/newsletters/%s/language/%s", url.PathEscape(newsletterID), url.PathEscape(language))

	body, statusCode, err := c.doRequest(ctx, "PUT", requestPath, req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var tmpError = &DataError{}
		if err := json.Unmarshal(body, &tmpError); err != nil {
			error := Errors{
				Status: statusCode,
				Detail: string(body),
			}
			return nil, &DataError{[]Errors{error}}
		}
		return nil, tmpError
	}

	var result UpdateNewsletterVariantResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) UpdateNewsletterTestGroupTranslation(ctx context.Context, newsletterID string, language string, testGroupID string, req *UpdateNewsletterVariantRequest) (*UpdateNewsletterVariantResponse, error) {
	requestPath := fmt.Sprintf("/v1/newsletters/%s/test_group/%s/language/%s", url.PathEscape(newsletterID), url.PathEscape(testGroupID), url.PathEscape(language))

	body, statusCode, err := c.doRequest(ctx, "PUT", requestPath, req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var tmpError = &DataError{}
		if err := json.Unmarshal(body, &tmpError); err != nil {
			error := Errors{
				Status: statusCode,
				Detail: string(body),
			}
			return nil, &DataError{[]Errors{error}}
		}
		return nil, tmpError
	}

	var result UpdateNewsletterVariantResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *APIClient) ListNewsletterTestGroups(ctx context.Context, newsletterID string) (*ListNewsletterTestGroupsResponse, error) {
	requestPath := fmt.Sprintf("/v1/newsletters/%s/test_groups", url.PathEscape(newsletterID))

	body, statusCode, err := c.doRequest(ctx, http.MethodGet, requestPath, nil)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var tmpError = &DataError{}
		if err := json.Unmarshal(body, &tmpError); err != nil {
			error := Errors{
				Status: statusCode,
				Detail: string(body),
			}
			return nil, &DataError{[]Errors{error}}
		}
		return nil, tmpError
	}

	var result ListNewsletterTestGroupsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
