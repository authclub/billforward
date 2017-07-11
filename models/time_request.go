package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimeRequest Request object for advancing a subscription's flow through time.
// swagger:model TimeRequest
type TimeRequest struct {

	// {"default":true,"description":"When advancing onto an instant in time:<br><span class=\"label label-default\">true</span> &mdash; Action the events scheduled for your destination time. Amendments scheduled at your destined time will be actioned during this time travel. Advancing to a period boundary will promote your subscription to the period on the future side of that boundary.<br><span class=\"label label-default\">false</span> &mdash; Do not action events scheduled for your destination time. Amendments scheduled at your destined time will not be actioned during this time travel. Advancing to a period boundary will result in your subscription's remaining within the period on the past side of that boundary.","verbs":["POST","GET"]}
	AdvanceInclusively *bool `json:"advanceInclusively,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created *strfmt.DateTime `json:"created,omitempty"`

	// {"default":false,"description":"Changes described in the response:<br><span class=\"label label-default\">true</span> &mdash; Are not actually performed; your subscription remains unchanged, no billing events run, and no invoices are executed.<br><span class=\"label label-default\">false</span> &mdash; Are actually performed and committed.","verbs":["POST","GET"]}
	DryRun *bool `json:"dryRun,omitempty"`

	// {"default":"<span class=\"label label-default\">SingleAttempt</span>","description":"What strategy to use when executing any invoices raised as time advances:<br><span class=\"label label-default\">SingleAttempt</span> &mdash; Execute any invoice just once.<br><span class=\"label label-default\">FollowDunning</span> &mdash; Apply the existing dunning strategy when executing invoices.<br><span class=\"label label-default\">None</span>: Do not execute invoices.","verbs":["POST","GET"]}
	ExecutionStrategy string `json:"executionStrategy,omitempty"`

	// {"default":false,"description":"Once the subscription is advanced through time:<br><span class=\"label label-default\">true</span> &mdash; Freeze the subscription.<br><span class=\"label label-default\">false</span> &mdash; Do not freeze the subscription.","verbs":["POST","GET"]}
	FreezeOnCompletion *bool `json:"freezeOnCompletion,omitempty"`

	// {"default":true,"description":"As time scrubs forward:<br><span class=\"label label-default\">true</span> &mdash; Run any amendments that were scheduled to action.<br><span class=\"label label-default\">false</span> &mdash; Do not run any amendments that were scheduled to action.","verbs":["POST","GET"]}
	HandleAmendments *bool `json:"handleAmendments,omitempty"`

	// {"description":"(Required: one of [`periods`, `to`])<br>The number of period boundaries up to which the subscription should advance.
	// A 1-value advances the subscription to the end of its current service period.
	// Higher values advance the subscription to subsequent period boundaries."verbs":["POST","GET"]}
	Periods *int32 `json:"periods,omitempty"`

	// {"default":false,"description":"As time scrubs forward:<br><span class=\"label label-default\">true</span> &mdash; Raise no invoice upon advancing over a period boundary.<br><span class=\"label label-default\">false</span> &mdash; Raise invoices for any period that is entered.","verbs":["POST","GET"]}
	SkipIntermediatePeriods *bool `json:"skipIntermediatePeriods,omitempty"`

	// {"description":"(Required: one of [`periods`, `to`])<br>The time up until which the subscription should be fast-forwarded.","verbs":["POST","GET"]}
	To *strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this time request
func (m *TimeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var timeRequestTypeExecutionStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SingleAttempt","FollowDunning","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeRequestTypeExecutionStrategyPropEnum = append(timeRequestTypeExecutionStrategyPropEnum, v)
	}
}

const (
	// TimeRequestExecutionStrategySingleAttempt captures enum value "SingleAttempt"
	TimeRequestExecutionStrategySingleAttempt string = "SingleAttempt"
	// TimeRequestExecutionStrategyFollowDunning captures enum value "FollowDunning"
	TimeRequestExecutionStrategyFollowDunning string = "FollowDunning"
	// TimeRequestExecutionStrategyNone captures enum value "None"
	TimeRequestExecutionStrategyNone string = "None"
)

// prop value enum
func (m *TimeRequest) validateExecutionStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, timeRequestTypeExecutionStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TimeRequest) validateExecutionStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutionStrategyEnum("executionStrategy", "body", m.ExecutionStrategy); err != nil {
		return err
	}

	return nil
}
