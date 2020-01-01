// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceComplianceUserOverviewRequestBuilder is request builder for DeviceComplianceUserOverview
type DeviceComplianceUserOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceComplianceUserOverviewRequest
func (b *DeviceComplianceUserOverviewRequestBuilder) Request() *DeviceComplianceUserOverviewRequest {
	return &DeviceComplianceUserOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceComplianceUserOverviewRequest is request for DeviceComplianceUserOverview
type DeviceComplianceUserOverviewRequest struct{ BaseRequest }

// Get performs GET request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Get(ctx context.Context) (resObj *DeviceComplianceUserOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Update(ctx context.Context, reqObj *DeviceComplianceUserOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}