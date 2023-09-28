package customerio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type UpdateCampaignActionRequest struct {
	ParentActionID int64             `json:"parent_action_id,omitempty"`
	Created        time.Time         `json:"created,omitempty"`
	Updated        time.Time         `json:"updated,omitempty"`
	Body           string            `json:"body,omitempty"`
	SendingState   string            `json:"sending_state,omitempty"`
	FromID         int64             `json:"from_id,omitempty"`
	ReplyToID      int64             `json:"reply_to_id,omitempty"`
	Recipient      string            `json:"recipient,omitempty"`
	Subject        string            `json:"subject,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
}

type UpdateCampaignActionResponse struct {
	ID             string    `json:"id"`
	CampaignID     int64     `json:"campaign_id"`
	ParentActionID int64     `json:"parent_action_id"`
	DeduplicateID  string    `json:"deduplicate_id"`
	Name           string    `json:"name"`
	Layout         string    `json:"layout"`
	Body           string    `json:"body"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
	Type           string    `json:"type"`
	SendingState   string    `json:"sending_state"`
	From           string    `json:"from"`
	FromID         int64     `json:"from_id"`
	ReplyTo        string    `json:"reply_to"`
	ReplyToID      int64     `json:"reply_to_id"`
	Preprocessor   string    `json:"preprocessor"`
	Recipient      string    `json:"recipient"`
	Subject        string    `json:"subject"`
	Bcc            string    `json:"bcc"`
	FakeBcc        bool      `json:"fake_bcc"`
	PreheaderText  string    `json:"preheader_text"`
}

func (t *UpdateCampaignActionResponse) UnmarshalJSON(b []byte) error {
	var r struct {
		Action struct {
			ID             string `json:"id"`
			CampaignID     int64  `json:"campaign_id"`
			ParentActionID int64  `json:"parent_action_id"`
			DeduplicateID  string `json:"deduplicate_id"`
			Name           string `json:"name"`
			Layout         string `json:"layout"`
			Body           string `json:"body"`
			Created        int64  `json:"created"`
			Updated        int64  `json:"updated"`
			Type           string `json:"type"`
			SendingState   string `json:"sending_state"`
			From           string `json:"from"`
			FromID         int64  `json:"from_id"`
			ReplyTo        string `json:"reply_to"`
			ReplyToID      int64  `json:"reply_to_id"`
			Preprocessor   string `json:"preprocessor"`
			Recipient      string `json:"recipient"`
			Subject        string `json:"subject"`
			Bcc            string `json:"bcc"`
			FakeBcc        bool   `json:"fake_bcc"`
			PreheaderText  string `json:"preheader_text"`
		} `json:"action"`
	}

	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	t.ID = r.Action.ID
	t.CampaignID = r.Action.CampaignID
	t.ParentActionID = r.Action.ParentActionID
	t.DeduplicateID = r.Action.DeduplicateID
	t.Name = r.Action.Name
	t.Layout = r.Action.Layout
	t.Body = r.Action.Body
	t.Created = time.Unix(r.Action.Created, 0)
	t.Updated = time.Unix(r.Action.Updated, 0)
	t.Type = r.Action.Type
	t.SendingState = r.Action.SendingState
	t.From = r.Action.From
	t.FromID = r.Action.FromID
	t.ReplyTo = r.Action.ReplyTo
	t.ReplyToID = r.Action.ReplyToID
	t.Preprocessor = r.Action.Preprocessor
	t.Recipient = r.Action.Recipient
	t.Subject = r.Action.Subject
	t.Bcc = r.Action.Bcc
	t.FakeBcc = r.Action.FakeBcc
	t.PreheaderText = r.Action.PreheaderText

	return nil
}

type DataError struct {
	Errors []Errors `json:"errors"`
}
type Errors struct {
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (e *DataError) Error() string {
	if len(e.Errors) == 0 {
		return "error message undefined"
	}
	return e.Errors[0].Detail
}

func (c *BetaAPIClient) UpdateCampaignLocalizedAction(ctx context.Context, campaignID string, actionID string, locale string, req *UpdateCampaignActionRequest) (*UpdateCampaignActionResponse, error) {
	var requestPath string

	if locale != "" {
		requestPath = fmt.Sprintf("/v1/api/campaigns/%s/actions/%s/language/%s", url.PathEscape(campaignID), url.PathEscape(actionID), url.PathEscape(locale))
	} else {
		requestPath = fmt.Sprintf("/v1/api/campaigns/%s/actions/%s", url.PathEscape(campaignID), url.PathEscape(actionID))
	}

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

	var result UpdateCampaignActionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *BetaAPIClient) UpdateCampaignAction(ctx context.Context, campaignID string, actionID string, req *UpdateCampaignActionRequest) (*UpdateCampaignActionResponse, error) {
	requestPath := fmt.Sprintf("/v1/api/campaigns/%s/actions/%s", url.PathEscape(campaignID), url.PathEscape(actionID))

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

	var result UpdateCampaignActionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
