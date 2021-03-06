// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateProvisioningTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the default provisioning template version.
	DefaultVersionId *int64 `locationName:"defaultVersionId" type:"integer"`

	// The description of the fleet provisioning template.
	Description *string `locationName:"description" type:"string"`

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string `locationName:"provisioningRoleArn" min:"20" type:"string"`

	// The name of the fleet provisioning template.
	//
	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"templateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateProvisioningTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProvisioningTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProvisioningTemplateInput"}
	if s.ProvisioningRoleArn != nil && len(*s.ProvisioningRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningRoleArn", 20))
	}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProvisioningTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DefaultVersionId != nil {
		v := *s.DefaultVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultVersionId", protocol.Int64Value(v), metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.ProvisioningRoleArn != nil {
		v := *s.ProvisioningRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "provisioningRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "templateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateProvisioningTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateProvisioningTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProvisioningTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateProvisioningTemplate = "UpdateProvisioningTemplate"

// UpdateProvisioningTemplateRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates a fleet provisioning template.
//
//    // Example sending a request using UpdateProvisioningTemplateRequest.
//    req := client.UpdateProvisioningTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateProvisioningTemplateRequest(input *UpdateProvisioningTemplateInput) UpdateProvisioningTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateProvisioningTemplate,
		HTTPMethod: "PATCH",
		HTTPPath:   "/provisioning-templates/{templateName}",
	}

	if input == nil {
		input = &UpdateProvisioningTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateProvisioningTemplateOutput{})
	return UpdateProvisioningTemplateRequest{Request: req, Input: input, Copy: c.UpdateProvisioningTemplateRequest}
}

// UpdateProvisioningTemplateRequest is the request type for the
// UpdateProvisioningTemplate API operation.
type UpdateProvisioningTemplateRequest struct {
	*aws.Request
	Input *UpdateProvisioningTemplateInput
	Copy  func(*UpdateProvisioningTemplateInput) UpdateProvisioningTemplateRequest
}

// Send marshals and sends the UpdateProvisioningTemplate API request.
func (r UpdateProvisioningTemplateRequest) Send(ctx context.Context) (*UpdateProvisioningTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProvisioningTemplateResponse{
		UpdateProvisioningTemplateOutput: r.Request.Data.(*UpdateProvisioningTemplateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProvisioningTemplateResponse is the response type for the
// UpdateProvisioningTemplate API operation.
type UpdateProvisioningTemplateResponse struct {
	*UpdateProvisioningTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProvisioningTemplate request.
func (r *UpdateProvisioningTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
