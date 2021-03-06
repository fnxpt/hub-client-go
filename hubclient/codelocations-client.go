// Copyright 2018 Synopsys, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hubclient

import (
	"github.com/blackducksoftware/hub-client-go/hubapi"
)

func (c *Client) ListAllCodeLocations(options *hubapi.GetListOptions) (*hubapi.CodeLocationList, error) {
	codeLocationsURL := c.baseURL + "/api/codelocations" + hubapi.ParameterString(options)

	var codeLocationsList hubapi.CodeLocationList
	err := c.GetPage(codeLocationsURL, options, &codeLocationsList)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to retrieve code locations list")
	}

	return &codeLocationsList, nil
}

func (c *Client) ListCodeLocations(link hubapi.ResourceLink, options *hubapi.GetListOptions) (*hubapi.CodeLocationList, error) {

	var codeLocationList hubapi.CodeLocationList
	err := c.GetPage(link.Href, options, &codeLocationList)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to retrieve code location list")
	}

	return &codeLocationList, nil
}

func (c *Client) GetCodeLocation(link hubapi.ResourceLink) (*hubapi.CodeLocation, error) {

	var codeLocation hubapi.CodeLocation
	err := c.HttpGetJSON(link.Href, &codeLocation, 200)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to retrieve a code location")
	}

	return &codeLocation, nil
}

// DeleteCodeLocation deletes a code location using
// https://<base_hub_URL>/api.html#!/composite45code45location45rest45server/deleteCodeLocationUsingDELETE
func (c *Client) DeleteCodeLocation(codeLocationURL string) error {
	return c.HttpDelete(codeLocationURL, "application/json", 204)
}

func (c *Client) ListScanSummaries(link hubapi.ResourceLink) (*hubapi.ScanSummaryList, error) {
	// TODO: Need offset/limit
	var scanSummaryList hubapi.ScanSummaryList
	err := c.GetPage(link.Href, nil, &scanSummaryList)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to retrieve scan summary list")
	}

	return &scanSummaryList, nil
}

func (c *Client) GetScanSummary(link hubapi.ResourceLink) (*hubapi.ScanSummary, error) {

	var scanSummary hubapi.ScanSummary
	err := c.HttpGetJSON(link.Href, &scanSummary, 200)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to retrieve a scan summary")
	}

	return &scanSummary, nil
}
