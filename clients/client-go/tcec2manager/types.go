// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcec2manager

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster/clients/client-go/v18"
)

type (
	// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items
	Entry1 struct {

		// EC2 instance type
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items/properties/instanceType
		InstanceType string `json:"instanceType,omitempty"`

		// Amount of dollars for an hour of usage for this configuration
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items/properties/price
		Price float64 `json:"price,omitempty"`

		// EC2 region
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items/properties/region
		Region string `json:"region,omitempty"`

		// Possible values:
		//   * "spot"
		//   * "ondemand"
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items/properties/type
		Type string `json:"type,omitempty"`

		// EC2 availability zone identifier
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#/items/properties/zone
		Zone string `json:"zone,omitempty"`
	}

	// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items
	Entry2 struct {

		// Possible values:
		//   * "instanceType"
		//   * "region"
		//   * "price"
		//   * "minPrice"
		//   * "maxPrice"
		//   * "zone"
		//   * "type"
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items/properties/key
		Key string `json:"key,omitempty"`

		// One of:
		//   * Var5
		//   * Var6
		//   * Var7
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items/properties/restriction
		Restriction json.RawMessage `json:"restriction,omitempty"`
	}

	// This method returns a list of errors.  It currently gives the error code only
	// because we're not sure of the security implications of exposing the full
	// message.  We do store complete error messages, but are figuring out how to
	// best expose them
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#
	Errors struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors
		Errors []Var4 `json:"errors,omitempty"`
	}

	// This method provides a summary of the health in the EC2 account being managed.
	// Values for the overall account are provided, broken down by Region, Availability
	// Zone and Instance Type.
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#
	HealthOfTheEC2Account struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth
		RequestHealth []json.RawMessage `json:"requestHealth,omitempty"`

		// An overview of currently running instances
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running
		Running []Var2 `json:"running,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth
		TerminationHealth []json.RawMessage `json:"terminationHealth,omitempty"`
	}

	// This is an EC2-Manager specific wrapping of the request body for the
	// upstream EC2 API.  Values from this are passed through verbatim.  A small
	// number of checks are done on the data before making the call, as well as
	// having some schema keys set to ensure certain values are either present
	// or absent
	//
	// Defined properties:
	//
	//  struct {
	//
	//  	// This is the AMI Identifier for this spot request.  This image must
	//  	// already exist and must be in the region of the request.  Note that
	//  	// AMI images are per-region, so you must copy or regenerate the image
	//  	// for each region.
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/ImageId
	//  	ImageID string `json:"ImageId,omitempty"`
	//
	//  	// The instance type to use for this spot request
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/InstanceType
	//  	InstanceType string `json:"InstanceType,omitempty"`
	//
	//  	// A valid EC2 KeyPair name.  The KeyPair must already exist
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/KeyName
	//  	KeyName string `json:"KeyName,omitempty"`
	//
	//  	// This is a list of the security groups this image will use.  These
	//  	// groups must already exist in the region.
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/SecurityGroups
	//  	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	//  }
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo
	LaunchInfo json.RawMessage

	// A list of prices for EC2
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/prices.json#
	ListOfPrices []Entry1

	// A list of prices for EC2
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#
	ListOfRestrictionsForPrices []Entry2

	// A list of names of worker types
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/list-worker-types.json#
	ListOfWorkerTypes []string

	// Presented here are the fields that are absolutely 100% required to make a
	// spot request.  The `LaunchSpecification` property is an opaque datastructure
	// from EC2, however the fields which we know are absolutely required are
	// described
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#
	MakeASpotRequest struct {

		// A ClientToken string per the implementation requirements of the EC2 api.
		// This string must be no more than 64 characters of ASCII.  We restrict the
		// client tokens further to alphanumeric ASCII with the addition of the `-`
		// and `_` characters
		//
		// Syntax:     ^[a-zA-Z0-0_-]{1,64}
		// Min length: 1
		// Max length: 64
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/ClientToken
		ClientToken string `json:"ClientToken,omitempty"`

		// This is an EC2-Manager specific wrapping of the request body for the
		// upstream EC2 API.  Values from this are passed through verbatim.  A small
		// number of checks are done on the data before making the call, as well as
		// having some schema keys set to ensure certain values are either present
		// or absent
		//
		// Defined properties:
		//
		//  struct {
		//
		//  	// This is the AMI Identifier for this spot request.  This image must
		//  	// already exist and must be in the region of the request.  Note that
		//  	// AMI images are per-region, so you must copy or regenerate the image
		//  	// for each region.
		//  	//
		//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/ImageId
		//  	ImageID string `json:"ImageId,omitempty"`
		//
		//  	// The instance type to use for this spot request
		//  	//
		//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/InstanceType
		//  	InstanceType string `json:"InstanceType,omitempty"`
		//
		//  	// A valid EC2 KeyPair name.  The KeyPair must already exist
		//  	//
		//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/KeyName
		//  	KeyName string `json:"KeyName,omitempty"`
		//
		//  	// This is a list of the security groups this image will use.  These
		//  	// groups must already exist in the region.
		//  	//
		//  	// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo/properties/SecurityGroups
		//  	SecurityGroups []string `json:"SecurityGroups,omitempty"`
		//  }
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/LaunchInfo
		LaunchInfo json.RawMessage `json:"LaunchInfo,omitempty"`

		// The EC2 region in which this spot request is to be made.  This should be
		// the lower case api-identifier.  For example `us-east-1`
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/Region
		Region string `json:"Region,omitempty"`

		// Specify whether to use a spot request or an on-demand instance.  This is
		// not inferred from the SpotPrice being set or not because we want to allow
		// for the default behaviour for spot prices, which is to use the spot
		// market with a default price of the on-demand price
		//
		// Possible values:
		//   * "spot"
		//   * "on-demand"
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/RequestType
		RequestType string `json:"RequestType,omitempty"`

		// The actual price of the bid.  This is passed directly to the EC2 api and
		// so should not have any internal multipliers (e.g. capacity or utility)
		// applied
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/run-instance-request.json#/properties/SpotPrice
		SpotPrice float64 `json:"SpotPrice,omitempty"`
	}

	// Overview of computational resources for a given worker type
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#
	OverviewOfComputationalResources struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/pending
		Pending []json.RawMessage `json:"pending,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/running
		Running []interface{} `json:"running,omitempty"`
	}

	// Overview of computational resources for a given worker type
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-state.json#
	OverviewOfComputationalResources1 struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-state.json#/properties/instances
		Instances []interface{} `json:"instances,omitempty"`
	}

	// See http://schemas.taskcluster.net/ec2-manager/v1/create-key-pair.json#
	SSHPublicKey struct {

		// An OpenSSH format Public Key as described by tools.ietf.org/html/rfc4253#section-6.6
		//
		// Syntax:     ^(ssh-\S*)\s*(\S*)\s*(.*)$
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/create-key-pair.json#/properties/pubkey
		Pubkey string `json:"pubkey,omitempty"`
	}

	// Defined properties:
	//
	//  struct {
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/pending/items/properties/count
	//  	Count float64 `json:"count,omitempty"`
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/pending/items/properties/instanceType
	//  	InstanceType string `json:"instanceType,omitempty"`
	//
	//  	// Possible values:
	//  	//   * "instance"
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/pending/items/properties/type
	//  	Type string `json:"type,omitempty"`
	//  }
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/worker-type-resources.json#/properties/pending/items
	Var json.RawMessage

	// This is a list of outcomes for a specific region, availability zone and
	// instance type.  These are calls to the EC2 runInstances method, which
	// is how we request instances.  If a call to this method is successful,
	// then we expect to get an instance to match
	//
	// Defined properties:
	//
	//  struct {
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/az
	//  	Az string `json:"az,omitempty"`
	//
	//  	// The number of calls failed due to a misconfiguration of the worker type.  Due to the large number of error codes the EC2 API might return, this is a best effort categorization.  It covers codes which are like "Invalid%" using SQL pattern mattching on the codes from https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html It is not categorized by which field was invalid in this response
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/configuration_issue
	//  	Configuration_Issue int64 `json:"configuration_issue,omitempty"`
	//
	//  	// The total number of calls which failed, inrespective of why
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/failed
	//  	Failed int64 `json:"failed,omitempty"`
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/instanceType
	//  	InstanceType string `json:"instanceType,omitempty"`
	//
	//  	// Number of runInstances calls which have failed because there aren't
	//  	// enough hosts for the resources to be allocated.
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/insufficient_capacity
	//  	Insufficient_Capacity int64 `json:"insufficient_capacity,omitempty"`
	//
	//  	// The number of calls which failed due to a limit being exceeded.
	//  	// Due to the large number of error codes the EC2 API might return,
	//  	// this is a best effort categorization.  It covers codes which are
	//  	// like "%LimitExceeded" using SQL pattern mattching, but not
	//  	// RequestLimitExceeded on the codes from
	//  	// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html
	//  	// It is not categorized by which limit was exceeded in this response
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/limit_exceeded
	//  	Limit_Exceeded int64 `json:"limit_exceeded,omitempty"`
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/region
	//  	Region string `json:"region,omitempty"`
	//
	//  	// The number of instances which have been requested successfully
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/successful
	//  	Successful int64 `json:"successful,omitempty"`
	//
	//  	// Number of calls which have been throttled in this region.  These
	//  	// are errors with the code RequestLimitExceeded.
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items/properties/throttled_calls
	//  	Throttled_Calls int64 `json:"throttled_calls,omitempty"`
	//  }
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/requestHealth/items
	Var1 json.RawMessage

	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running/items
	Var2 struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running/items/properties/az
		Az string `json:"az,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running/items/properties/instanceType
		InstanceType string `json:"instanceType,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running/items/properties/region
		Region string `json:"region,omitempty"`

		// The number of currently running instances in this configuration
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/running/items/properties/running
		Running int64 `json:"running,omitempty"`
	}

	// This is a list of summaries of instances which have terminated
	//
	// Defined properties:
	//
	//  struct {
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/az
	//  	Az string `json:"az,omitempty"`
	//
	//  	// A count of the instances which were shutdown cleanty.  For the
	//  	// purposes of this API, a clean shutdown is one which was initiated
	//  	// by us.  This includes API shutdowns or workers ending themselves.
	//  	// It does not mean the actual workload ran successfully, rather that
	//  	// we chose to terminate it
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/clean_shutdown
	//  	Clean_Shutdown int64 `json:"clean_shutdown,omitempty"`
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/instanceType
	//  	InstanceType string `json:"instanceType,omitempty"`
	//
	//  	// The number of instances which were terminated due to a lack of
	//  	// capacity.  More than likely, this will always be zero because the
	//  	// new spot service is now synchronous, so runInstances calls should
	//  	// fail
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/insufficient_capacity
	//  	Insufficient_Capacity int64 `json:"insufficient_capacity,omitempty"`
	//
	//  	// The number of instances which were terminated due to not being able
	//  	// to find the AMI
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/missing_ami
	//  	Missing_Ami int64 `json:"missing_ami,omitempty"`
	//
	//  	// The number of terminations which we cannot find a code.  This means
	//  	// we cannot determine whether this should be classified as a good or
	//  	// bad outcome.  The specific reason is that the code which polls for
	//  	// termination reason was not able to run before the EC2 API dropped
	//  	// the instance from its database
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/no_code
	//  	No_Code int64 `json:"no_code,omitempty"`
	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/region
	//  	Region string `json:"region,omitempty"`
	//
	//  	// The number of instances which were killed by the spot service
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/spot_kill
	//  	Spot_Kill int64 `json:"spot_kill,omitempty"`
	//
	//  	// The number of instances which failed to start, either because of an
	//  	// error on our side or on the EC2 side
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/startup_failed
	//  	Startup_Failed int64 `json:"startup_failed,omitempty"`
	//
	//  	// The number of terminations which have a code which this code does
	//  	// not recognize
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/unknown_code
	//  	Unknown_Code int64 `json:"unknown_code,omitempty"`
	//
	//  	// The number of instances which were terminated due to exceeding the
	//  	// limit for number of ebs volumes
	//  	//
	//  	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items/properties/volume_limit_exceeded
	//  	Volume_Limit_Exceeded int64 `json:"volume_limit_exceeded,omitempty"`
	//  }
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/ec2-manager/v1/health.json#/properties/terminationHealth/items
	Var3 json.RawMessage

	// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items
	Var4 struct {

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/az
		Az string `json:"az,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/code
		Code string `json:"code,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/instanceType
		InstanceType string `json:"instanceType,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/message
		Message string `json:"message,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/region
		Region string `json:"region,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/time
		Time tcclient.Time `json:"time,omitempty"`

		// Possible values:
		//   * "instance-request"
		//   * "termination"
		//
		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/type
		Type string `json:"type,omitempty"`

		// See http://schemas.taskcluster.net/ec2-manager/v1/errors.json#/properties/errors/items/properties/workerType
		WorkerType string `json:"workerType,omitempty"`
	}

	// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items/properties/restriction/oneOf[0]
	Var5 float64

	// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items/properties/restriction/oneOf[1]
	Var6 string

	// See http://schemas.taskcluster.net/ec2-manager/v1/prices-request.json#/items/properties/restriction/oneOf[2]
	Var7 []string
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// LaunchInfo is of type json.RawMessage...
func (this *LaunchInfo) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *LaunchInfo) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("LaunchInfo: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Var is of type json.RawMessage...
func (this *Var) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Var) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Var: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Var1 is of type json.RawMessage...
func (this *Var1) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Var1) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Var1: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// Var3 is of type json.RawMessage...
func (this *Var3) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *Var3) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("Var3: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
