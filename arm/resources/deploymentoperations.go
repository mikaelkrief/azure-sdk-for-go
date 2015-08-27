package resources

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

// DeploymentOperations Client
type DeploymentOperationsClient struct {
	ResourceManagementClient
}

func NewDeploymentOperationsClient(subscriptionId string) DeploymentOperationsClient {
	return NewDeploymentOperationsClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewDeploymentOperationsClientWithBaseUri(baseUri string, subscriptionId string) DeploymentOperationsClient {
	return DeploymentOperationsClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// Get get a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. operationId is
// operation Id.
func (client DeploymentOperationsClient) Get(resourceGroupName string, deploymentName string, operationId string) (result DeploymentOperation, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, deploymentName, operationId)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client DeploymentOperationsClient) NewGetRequest(resourceGroupName string, deploymentName string, operationId string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
		"operationId":       url.QueryEscape(operationId),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
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
func (client DeploymentOperationsClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}"))
}

// List gets a list of deployments operations.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. deploymentName is the name of the deployment. top is query
// parameters.
func (client DeploymentOperationsClient) List(resourceGroupName string, deploymentName string, top int) (result DeploymentOperationsListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName, deploymentName, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", "Failure preparing request")
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
			ae = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.DeploymentOperationsClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client DeploymentOperationsClient) NewListRequest(resourceGroupName string, deploymentName string, top int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    url.QueryEscape(deploymentName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client DeploymentOperationsClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations"))
}
