// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutManagedScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the ID of an EMR cluster where the managed scaling policy is attached.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// Specifies the constraints for the managed scaling policy.
	//
	// ManagedScalingPolicy is a required field
	ManagedScalingPolicy *ManagedScalingPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutManagedScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutManagedScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutManagedScalingPolicyInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if s.ManagedScalingPolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("ManagedScalingPolicy"))
	}
	if s.ManagedScalingPolicy != nil {
		if err := s.ManagedScalingPolicy.Validate(); err != nil {
			invalidParams.AddNested("ManagedScalingPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutManagedScalingPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutManagedScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutManagedScalingPolicy = "PutManagedScalingPolicy"

// PutManagedScalingPolicyRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Creates or updates a managed scaling policy for an Amazon EMR cluster. The
// managed scaling policy defines the limits for resources, such as EC2 instances
// that can be added or terminated from a cluster. The policy only applies to
// the core and task nodes. The master node cannot be scaled after initial configuration.
//
//    // Example sending a request using PutManagedScalingPolicyRequest.
//    req := client.PutManagedScalingPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/PutManagedScalingPolicy
func (c *Client) PutManagedScalingPolicyRequest(input *PutManagedScalingPolicyInput) PutManagedScalingPolicyRequest {
	op := &aws.Operation{
		Name:       opPutManagedScalingPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutManagedScalingPolicyInput{}
	}

	req := c.newRequest(op, input, &PutManagedScalingPolicyOutput{})
	return PutManagedScalingPolicyRequest{Request: req, Input: input, Copy: c.PutManagedScalingPolicyRequest}
}

// PutManagedScalingPolicyRequest is the request type for the
// PutManagedScalingPolicy API operation.
type PutManagedScalingPolicyRequest struct {
	*aws.Request
	Input *PutManagedScalingPolicyInput
	Copy  func(*PutManagedScalingPolicyInput) PutManagedScalingPolicyRequest
}

// Send marshals and sends the PutManagedScalingPolicy API request.
func (r PutManagedScalingPolicyRequest) Send(ctx context.Context) (*PutManagedScalingPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutManagedScalingPolicyResponse{
		PutManagedScalingPolicyOutput: r.Request.Data.(*PutManagedScalingPolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutManagedScalingPolicyResponse is the response type for the
// PutManagedScalingPolicy API operation.
type PutManagedScalingPolicyResponse struct {
	*PutManagedScalingPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutManagedScalingPolicy request.
func (r *PutManagedScalingPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
