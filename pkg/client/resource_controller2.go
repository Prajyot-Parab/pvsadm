package client

import (
	"context"
	"encoding/json"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/common"
	rcv2 "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type ResourceControllerV2 struct {
	ResourceControllerV2 *rcv2.ResourceControllerV2
}

func NewResourceControllerV2(r *rcv2.ResourceControllerV2Options) (*ResourceControllerV2, error) {
	rc, err := rcv2.NewResourceControllerV2(r)
	return &ResourceControllerV2{
		rc,
	}, err
}

func (resourceController *ResourceControllerV2) CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions) (result *rcv2.ResourceKey, response *core.DetailedResponse, err error) {
	return resourceController.CreateResourceKeyWithContext(context.Background(), createResourceKeyOptions)
}

type CreateResourceKeyOptions struct {
	*rcv2.CreateResourceKeyOptions

	// Overriding the Parameters to accommodate the HMAC parameter
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Overriding the CreateResourceKeyWithContext function from the ResourceControllerV2 code to work with HMAC parameter
func (resourceController *ResourceControllerV2) CreateResourceKeyWithContext(ctx context.Context, createResourceKeyOptions *CreateResourceKeyOptions) (result *rcv2.ResourceKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceKeyOptions, "createResourceKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceKeyOptions, "createResourceKeyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = resourceController.ResourceControllerV2.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(resourceController.ResourceControllerV2.Service.Options.URL, `/v2/resource_keys`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceKeyOptions.CreateResourceKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_controller", "V2", "CreateResourceKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createResourceKeyOptions.Name != nil {
		body["name"] = createResourceKeyOptions.Name
	}
	if createResourceKeyOptions.Source != nil {
		body["source"] = createResourceKeyOptions.Source
	}
	if createResourceKeyOptions.Parameters != nil {
		body["parameters"] = createResourceKeyOptions.Parameters
	}
	if createResourceKeyOptions.Role != nil {
		body["role"] = createResourceKeyOptions.Role
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = resourceController.ResourceControllerV2.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, rcv2.UnmarshalResourceKey)
	if err != nil {
		return
	}
	response.Result = result

	return
}
