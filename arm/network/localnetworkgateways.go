package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// LocalNetworkGateways Client
type LocalNetworkGatewaysClient struct {
	NetworkResourceProviderClient
}

func NewLocalNetworkGatewaysClient(subscriptionId string) LocalNetworkGatewaysClient {
	return NewLocalNetworkGatewaysClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewLocalNetworkGatewaysClientWithBaseUri(baseUri string, subscriptionId string) LocalNetworkGatewaysClient {
	return LocalNetworkGatewaysClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// CreateOrUpdate the Put LocalNetworkGateway operation creates/updates a
// local network gateway in the specified resource group through Network
// resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
// parameters is parameters supplied to the Begin Create or update Local
// Network Gateway operation through Network resource provider.
func (client LocalNetworkGatewaysClient) CreateOrUpdate(resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (result LocalNetworkGateway, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusCreated, http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) NewCreateOrUpdateRequest(resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"))
}

// Get the Get LocalNetworkGateway operation retrieves information about the
// specified local network gateway through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Get(resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGateway, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, localNetworkGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client LocalNetworkGatewaysClient) NewGetRequest(resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client LocalNetworkGatewaysClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"))
}

// Delete the Delete LocalNetworkGateway operation deletes the specifed local
// network Gateway through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Delete(resourceGroupName string, localNetworkGatewayName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(resourceGroupName, localNetworkGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK))
	if err == nil {
		err = client.IsPollingAllowed(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client LocalNetworkGatewaysClient) NewDeleteRequest(resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client LocalNetworkGatewaysClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}/"))
}

// List the List LocalNetworkGateways opertion retrieves all the local network
// gateways stored.
//
// resourceGroupName is the name of the resource group.
func (client LocalNetworkGatewaysClient) List(resourceGroupName string) (result LocalNetworkGatewayListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.LocalNetworkGatewaysClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client LocalNetworkGatewaysClient) NewListRequest(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client LocalNetworkGatewaysClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways"))
}
