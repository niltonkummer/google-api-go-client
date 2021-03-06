// Package cloudmonitoring provides access to the Cloud Monitoring API.
//
// See https://cloud.google.com/monitoring/v2beta2/
//
// Usage example:
//
//   import "google.golang.org/api/cloudmonitoring/v2beta2"
//   ...
//   cloudmonitoringService, err := cloudmonitoring.New(oauthHttpClient)
package cloudmonitoring

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "cloudmonitoring:v2beta2"
const apiName = "cloudmonitoring"
const apiVersion = "v2beta2"
const basePath = "https://www.googleapis.com/cloudmonitoring/v2beta2/projects/"

// OAuth2 scopes used by this API.
const (
	// View and write monitoring data for all of your Google and third-party
	// Cloud and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.MetricDescriptors = NewMetricDescriptorsService(s)
	s.Timeseries = NewTimeseriesService(s)
	s.TimeseriesDescriptors = NewTimeseriesDescriptorsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	MetricDescriptors *MetricDescriptorsService

	Timeseries *TimeseriesService

	TimeseriesDescriptors *TimeseriesDescriptorsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewMetricDescriptorsService(s *Service) *MetricDescriptorsService {
	rs := &MetricDescriptorsService{s: s}
	return rs
}

type MetricDescriptorsService struct {
	s *Service
}

func NewTimeseriesService(s *Service) *TimeseriesService {
	rs := &TimeseriesService{s: s}
	return rs
}

type TimeseriesService struct {
	s *Service
}

func NewTimeseriesDescriptorsService(s *Service) *TimeseriesDescriptorsService {
	rs := &TimeseriesDescriptorsService{s: s}
	return rs
}

type TimeseriesDescriptorsService struct {
	s *Service
}

type DeleteMetricDescriptorResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#deleteMetricDescriptorResponse".
	Kind string `json:"kind,omitempty"`
}

type ListMetricDescriptorsRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listMetricDescriptorsRequest".
	Kind string `json:"kind,omitempty"`
}

type ListMetricDescriptorsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listMetricDescriptorsResponse".
	Kind string `json:"kind,omitempty"`

	// Metrics: The returned metric descriptors.
	Metrics []*MetricDescriptor `json:"metrics,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, pass this value to the pageToken query
	// parameter.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListTimeseriesDescriptorsRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesDescriptorsRequest".
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesDescriptorsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesDescriptorsResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, set this value to the pageToken query
	// parameter.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Oldest: The oldest timestamp of the interval of this query, as an RFC
	// 3339 string.
	Oldest string `json:"oldest,omitempty"`

	// Timeseries: The returned time series descriptors.
	Timeseries []*TimeseriesDescriptor `json:"timeseries,omitempty"`

	// Youngest: The youngest timestamp of the interval of this query, as an
	// RFC 3339 string.
	Youngest string `json:"youngest,omitempty"`
}

type ListTimeseriesRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesRequest".
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, set the pageToken query parameter to this
	// value. All of the points of a time series will be returned before
	// returning any point of the subsequent time series.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Oldest: The oldest timestamp of the interval of this query as an RFC
	// 3339 string.
	Oldest string `json:"oldest,omitempty"`

	// Timeseries: The returned time series.
	Timeseries []*Timeseries `json:"timeseries,omitempty"`

	// Youngest: The youngest timestamp of the interval of this query as an
	// RFC 3339 string.
	Youngest string `json:"youngest,omitempty"`
}

type MetricDescriptor struct {
	// Description: Description of this metric.
	Description string `json:"description,omitempty"`

	// Labels: Labels defined for this metric.
	Labels []*MetricDescriptorLabelDescriptor `json:"labels,omitempty"`

	// Name: The name of this metric.
	Name string `json:"name,omitempty"`

	// Project: The project ID to which the metric belongs.
	Project string `json:"project,omitempty"`

	// TypeDescriptor: Type description for this metric.
	TypeDescriptor *MetricDescriptorTypeDescriptor `json:"typeDescriptor,omitempty"`
}

type MetricDescriptorLabelDescriptor struct {
	// Description: Label description.
	Description string `json:"description,omitempty"`

	// Key: Label key.
	Key string `json:"key,omitempty"`
}

type MetricDescriptorTypeDescriptor struct {
	// MetricType: The method of collecting data for the metric. See Metric
	// types.
	MetricType string `json:"metricType,omitempty"`

	// ValueType: The data type of of individual points in the metric's time
	// series. See Metric value types.
	ValueType string `json:"valueType,omitempty"`
}

type Point struct {
	// BoolValue: The value of this data point. Either "true" or "false".
	BoolValue bool `json:"boolValue,omitempty"`

	// DistributionValue: The value of this data point as a distribution. A
	// distribution value can contain a list of buckets and/or an
	// underflowBucket and an overflowBucket. The values of these points can
	// be used to create a histogram.
	DistributionValue *PointDistribution `json:"distributionValue,omitempty"`

	// DoubleValue: The value of this data point as a double-precision
	// floating-point number.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// End: The interval [start, end] is the time period to which the
	// point's value applies. For gauge metrics, whose values are
	// instantaneous measurements, this interval should be empty (start
	// should equal end). For cumulative metrics (of which deltas and rates
	// are special cases), the interval should be non-empty. Both start and
	// end are RFC 3339 strings.
	End string `json:"end,omitempty"`

	// Int64Value: The value of this data point as a 64-bit integer.
	Int64Value int64 `json:"int64Value,omitempty,string"`

	// Start: The interval [start, end] is the time period to which the
	// point's value applies. For gauge metrics, whose values are
	// instantaneous measurements, this interval should be empty (start
	// should equal end). For cumulative metrics (of which deltas and rates
	// are special cases), the interval should be non-empty. Both start and
	// end are RFC 3339 strings.
	Start string `json:"start,omitempty"`

	// StringValue: The value of this data point in string format.
	StringValue string `json:"stringValue,omitempty"`
}

type PointDistribution struct {
	// Buckets: The finite buckets.
	Buckets []*PointDistributionBucket `json:"buckets,omitempty"`

	// OverflowBucket: The overflow bucket.
	OverflowBucket *PointDistributionOverflowBucket `json:"overflowBucket,omitempty"`

	// UnderflowBucket: The underflow bucket.
	UnderflowBucket *PointDistributionUnderflowBucket `json:"underflowBucket,omitempty"`
}

type PointDistributionBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// LowerBound: The lower bound of the value interval of this bucket
	// (inclusive).
	LowerBound float64 `json:"lowerBound,omitempty"`

	// UpperBound: The upper bound of the value interval of this bucket
	// (exclusive).
	UpperBound float64 `json:"upperBound,omitempty"`
}

type PointDistributionOverflowBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// LowerBound: The lower bound of the value interval of this bucket
	// (inclusive).
	LowerBound float64 `json:"lowerBound,omitempty"`
}

type PointDistributionUnderflowBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// UpperBound: The upper bound of the value interval of this bucket
	// (exclusive).
	UpperBound float64 `json:"upperBound,omitempty"`
}

type Timeseries struct {
	// Points: The data points of this time series. The points are listed in
	// order of their end timestamp, from younger to older.
	Points []*Point `json:"points,omitempty"`

	// TimeseriesDesc: The descriptor of this time series.
	TimeseriesDesc *TimeseriesDescriptor `json:"timeseriesDesc,omitempty"`
}

type TimeseriesDescriptor struct {
	// Labels: The label's name.
	Labels map[string]string `json:"labels,omitempty"`

	// Metric: The name of the metric.
	Metric string `json:"metric,omitempty"`

	// Project: The Developers Console project number to which this time
	// series belongs.
	Project string `json:"project,omitempty"`
}

type TimeseriesDescriptorLabel struct {
	// Key: The label's name.
	Key string `json:"key,omitempty"`

	// Value: The label's value.
	Value string `json:"value,omitempty"`
}

type TimeseriesPoint struct {
	// Point: The data point in this time series snapshot.
	Point *Point `json:"point,omitempty"`

	// TimeseriesDesc: The descriptor of this time series.
	TimeseriesDesc *TimeseriesDescriptor `json:"timeseriesDesc,omitempty"`
}

type WriteTimeseriesRequest struct {
	// CommonLabels: The label's name.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Timeseries: Provide time series specific labels and the data points
	// for each time series. The labels in timeseries and the common_labels
	// should form a complete list of labels that required by the metric.
	Timeseries []*TimeseriesPoint `json:"timeseries,omitempty"`
}

type WriteTimeseriesResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#writeTimeseriesResponse".
	Kind string `json:"kind,omitempty"`
}

// method id "cloudmonitoring.metricDescriptors.create":

type MetricDescriptorsCreateCall struct {
	s                *Service
	project          string
	metricdescriptor *MetricDescriptor
	opt_             map[string]interface{}
}

// Create: Create a new metric.
func (r *MetricDescriptorsService) Create(project string, metricdescriptor *MetricDescriptor) *MetricDescriptorsCreateCall {
	c := &MetricDescriptorsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metricdescriptor = metricdescriptor
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsCreateCall) Fields(s ...googleapi.Field) *MetricDescriptorsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MetricDescriptorsCreateCall) Do() (*MetricDescriptor, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.metricdescriptor)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/metricDescriptors")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *MetricDescriptor
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new metric.",
	//   "httpMethod": "POST",
	//   "id": "cloudmonitoring.metricDescriptors.create",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project id. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors",
	//   "request": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "response": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.metricDescriptors.delete":

type MetricDescriptorsDeleteCall struct {
	s       *Service
	project string
	metric  string
	opt_    map[string]interface{}
}

// Delete: Delete an existing metric.
func (r *MetricDescriptorsService) Delete(project string, metric string) *MetricDescriptorsDeleteCall {
	c := &MetricDescriptorsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metric = metric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsDeleteCall) Fields(s ...googleapi.Field) *MetricDescriptorsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MetricDescriptorsDeleteCall) Do() (*DeleteMetricDescriptorResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/metricDescriptors/{metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DeleteMetricDescriptorResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete an existing metric.",
	//   "httpMethod": "DELETE",
	//   "id": "cloudmonitoring.metricDescriptors.delete",
	//   "parameterOrder": [
	//     "project",
	//     "metric"
	//   ],
	//   "parameters": {
	//     "metric": {
	//       "description": "Name of the metric.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which the metric belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors/{metric}",
	//   "response": {
	//     "$ref": "DeleteMetricDescriptorResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.metricDescriptors.list":

type MetricDescriptorsListCall struct {
	s                            *Service
	project                      string
	listmetricdescriptorsrequest *ListMetricDescriptorsRequest
	opt_                         map[string]interface{}
}

// List: List metric descriptors that match the query. If the query is
// not set, then all of the metric descriptors will be returned. Large
// responses will be paginated, use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.
func (r *MetricDescriptorsService) List(project string, listmetricdescriptorsrequest *ListMetricDescriptorsRequest) *MetricDescriptorsListCall {
	c := &MetricDescriptorsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.listmetricdescriptorsrequest = listmetricdescriptorsrequest
	return c
}

// Count sets the optional parameter "count": Maximum number of metric
// descriptors per page. Used for pagination. If not specified, count =
// 100.
func (c *MetricDescriptorsListCall) Count(count int64) *MetricDescriptorsListCall {
	c.opt_["count"] = count
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *MetricDescriptorsListCall) PageToken(pageToken string) *MetricDescriptorsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Query sets the optional parameter "query": The query used to search
// against existing metrics. Separate keywords with a space; the service
// joins all keywords with AND, meaning that all keywords must match for
// a metric to be returned. If this field is omitted, all metrics are
// returned. If an empty string is passed with this field, no metrics
// are returned.
func (c *MetricDescriptorsListCall) Query(query string) *MetricDescriptorsListCall {
	c.opt_["query"] = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsListCall) Fields(s ...googleapi.Field) *MetricDescriptorsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MetricDescriptorsListCall) Do() (*ListMetricDescriptorsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["query"]; ok {
		params.Set("query", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/metricDescriptors")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListMetricDescriptorsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List metric descriptors that match the query. If the query is not set, then all of the metric descriptors will be returned. Large responses will be paginated, use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.metricDescriptors.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "count": {
	//       "default": "100",
	//       "description": "Maximum number of metric descriptors per page. Used for pagination. If not specified, count = 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project id. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "The query used to search against existing metrics. Separate keywords with a space; the service joins all keywords with AND, meaning that all keywords must match for a metric to be returned. If this field is omitted, all metrics are returned. If an empty string is passed with this field, no metrics are returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors",
	//   "request": {
	//     "$ref": "ListMetricDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListMetricDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseries.list":

type TimeseriesListCall struct {
	s                     *Service
	project               string
	metric                string
	youngest              string
	listtimeseriesrequest *ListTimeseriesRequest
	opt_                  map[string]interface{}
}

// List: List the data points of the time series that match the metric
// and labels values and that have data points in the interval. Large
// responses are paginated; use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.
func (r *TimeseriesService) List(project string, metric string, youngest string, listtimeseriesrequest *ListTimeseriesRequest) *TimeseriesListCall {
	c := &TimeseriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metric = metric
	c.youngest = youngest
	c.listtimeseriesrequest = listtimeseriesrequest
	return c
}

// Aggregator sets the optional parameter "aggregator": The aggregation
// function that will reduce the data points in each window to a single
// point. This parameter is only valid for non-cumulative metrics with a
// value type of INT64 or DOUBLE.
//
// Possible values:
//   "max"
//   "mean"
//   "min"
//   "sum"
func (c *TimeseriesListCall) Aggregator(aggregator string) *TimeseriesListCall {
	c.opt_["aggregator"] = aggregator
	return c
}

// Count sets the optional parameter "count": Maximum number of data
// points per page, which is used for pagination of results.
func (c *TimeseriesListCall) Count(count int64) *TimeseriesListCall {
	c.opt_["count"] = count
	return c
}

// Labels sets the optional parameter "labels": A collection of labels
// for the matching time series, which are represented as:
// - key==value: key equals the value
// - key=~value: key regex matches the value
// - key!=value: key does not equal the value
// - key!~value: key regex does not match the value  For example, to
// list all of the time series descriptors for the region us-central1,
// you could
// specify:
// label=cloud.googleapis.com%2Flocation=~us-central1.*
func (c *TimeseriesListCall) Labels(labels string) *TimeseriesListCall {
	c.opt_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest": Start of the time
// interval (exclusive), which is expressed as an RFC 3339 timestamp. If
// neither oldest nor timespan is specified, the default time interval
// will be (youngest - 4 hours, youngest]
func (c *TimeseriesListCall) Oldest(oldest string) *TimeseriesListCall {
	c.opt_["oldest"] = oldest
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *TimeseriesListCall) PageToken(pageToken string) *TimeseriesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Timespan sets the optional parameter "timespan": Length of the time
// interval to query, which is an alternative way to declare the
// interval: (youngest - timespan, youngest]. The timespan and oldest
// parameters should not be used together. Units:
// - s: second
// - m: minute
// - h: hour
// - d: day
// - w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for
// example: 2w3d is not allowed; you should use 17d instead.
//
// If neither oldest nor timespan is specified, the default time
// interval will be (youngest - 4 hours, youngest].
func (c *TimeseriesListCall) Timespan(timespan string) *TimeseriesListCall {
	c.opt_["timespan"] = timespan
	return c
}

// Window sets the optional parameter "window": The sampling window. At
// most one data point will be returned for each window in the requested
// time interval. This parameter is only valid for non-cumulative metric
// types. Units:
// - m: minute
// - h: hour
// - d: day
// - w: week  Examples: 3m, 4w. Only one unit is allowed, for example:
// 2w3d is not allowed; you should use 17d instead.
func (c *TimeseriesListCall) Window(window string) *TimeseriesListCall {
	c.opt_["window"] = window
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesListCall) Fields(s ...googleapi.Field) *TimeseriesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TimeseriesListCall) Do() (*ListTimeseriesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("youngest", fmt.Sprintf("%v", c.youngest))
	if v, ok := c.opt_["aggregator"]; ok {
		params.Set("aggregator", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["oldest"]; ok {
		params.Set("oldest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["timespan"]; ok {
		params.Set("timespan", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["window"]; ok {
		params.Set("window", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/timeseries/{metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTimeseriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the data points of the time series that match the metric and labels values and that have data points in the interval. Large responses are paginated; use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.timeseries.list",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "aggregator": {
	//       "description": "The aggregation function that will reduce the data points in each window to a single point. This parameter is only valid for non-cumulative metrics with a value type of INT64 or DOUBLE.",
	//       "enum": [
	//         "max",
	//         "mean",
	//         "min",
	//         "sum"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "count": {
	//       "default": "6000",
	//       "description": "Maximum number of data points per page, which is used for pagination of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "12000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "description": "A collection of labels for the matching time series, which are represented as:  \n- key==value: key equals the value \n- key=~value: key regex matches the value \n- key!=value: key does not equal the value \n- key!~value: key regex does not match the value  For example, to list all of the time series descriptors for the region us-central1, you could specify:\nlabel=cloud.googleapis.com%2Flocation=~us-central1.*",
	//       "location": "query",
	//       "pattern": "(.+?)(==|=~|!=|!~)(.+)",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "metric": {
	//       "description": "Metric names are protocol-free URLs as listed in the Supported Metrics page. For example, compute.googleapis.com/instance/disk/read_ops_count.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "description": "Start of the time interval (exclusive), which is expressed as an RFC 3339 timestamp. If neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest]",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which this time series belongs. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "description": "Length of the time interval to query, which is an alternative way to declare the interval: (youngest - timespan, youngest]. The timespan and oldest parameters should not be used together. Units:  \n- s: second \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.\n\nIf neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest].",
	//       "location": "query",
	//       "pattern": "[0-9]+[smhdw]?",
	//       "type": "string"
	//     },
	//     "window": {
	//       "description": "The sampling window. At most one data point will be returned for each window in the requested time interval. This parameter is only valid for non-cumulative metric types. Units:  \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.",
	//       "location": "query",
	//       "pattern": "[0-9]+[mhdw]?",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "description": "End of the time interval (inclusive), which is expressed as an RFC 3339 timestamp.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseries/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseries.write":

type TimeseriesWriteCall struct {
	s                      *Service
	project                string
	writetimeseriesrequest *WriteTimeseriesRequest
	opt_                   map[string]interface{}
}

// Write: Put data points to one or more time series for one or more
// metrics. If a time series does not exist, a new time series will be
// created. It is not allowed to write a time series point that is older
// than the existing youngest point of that time series. Points that are
// older than the existing youngest point of that time series will be
// discarded silently. Therefore, users should make sure that points of
// a time series are written sequentially in the order of their end
// time.
func (r *TimeseriesService) Write(project string, writetimeseriesrequest *WriteTimeseriesRequest) *TimeseriesWriteCall {
	c := &TimeseriesWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.writetimeseriesrequest = writetimeseriesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesWriteCall) Fields(s ...googleapi.Field) *TimeseriesWriteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TimeseriesWriteCall) Do() (*WriteTimeseriesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writetimeseriesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/timeseries:write")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *WriteTimeseriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Put data points to one or more time series for one or more metrics. If a time series does not exist, a new time series will be created. It is not allowed to write a time series point that is older than the existing youngest point of that time series. Points that are older than the existing youngest point of that time series will be discarded silently. Therefore, users should make sure that points of a time series are written sequentially in the order of their end time.",
	//   "httpMethod": "POST",
	//   "id": "cloudmonitoring.timeseries.write",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseries:write",
	//   "request": {
	//     "$ref": "WriteTimeseriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteTimeseriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseriesDescriptors.list":

type TimeseriesDescriptorsListCall struct {
	s                                *Service
	project                          string
	metric                           string
	youngest                         string
	listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest
	opt_                             map[string]interface{}
}

// List: List the descriptors of the time series that match the metric
// and labels values and that have data points in the interval. Large
// responses are paginated; use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.
func (r *TimeseriesDescriptorsService) List(project string, metric string, youngest string, listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest) *TimeseriesDescriptorsListCall {
	c := &TimeseriesDescriptorsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.metric = metric
	c.youngest = youngest
	c.listtimeseriesdescriptorsrequest = listtimeseriesdescriptorsrequest
	return c
}

// Aggregator sets the optional parameter "aggregator": The aggregation
// function that will reduce the data points in each window to a single
// point. This parameter is only valid for non-cumulative metrics with a
// value type of INT64 or DOUBLE.
//
// Possible values:
//   "max"
//   "mean"
//   "min"
//   "sum"
func (c *TimeseriesDescriptorsListCall) Aggregator(aggregator string) *TimeseriesDescriptorsListCall {
	c.opt_["aggregator"] = aggregator
	return c
}

// Count sets the optional parameter "count": Maximum number of time
// series descriptors per page. Used for pagination. If not specified,
// count = 100.
func (c *TimeseriesDescriptorsListCall) Count(count int64) *TimeseriesDescriptorsListCall {
	c.opt_["count"] = count
	return c
}

// Labels sets the optional parameter "labels": A collection of labels
// for the matching time series, which are represented as:
// - key==value: key equals the value
// - key=~value: key regex matches the value
// - key!=value: key does not equal the value
// - key!~value: key regex does not match the value  For example, to
// list all of the time series descriptors for the region us-central1,
// you could
// specify:
// label=cloud.googleapis.com%2Flocation=~us-central1.*
func (c *TimeseriesDescriptorsListCall) Labels(labels string) *TimeseriesDescriptorsListCall {
	c.opt_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest": Start of the time
// interval (exclusive), which is expressed as an RFC 3339 timestamp. If
// neither oldest nor timespan is specified, the default time interval
// will be (youngest - 4 hours, youngest]
func (c *TimeseriesDescriptorsListCall) Oldest(oldest string) *TimeseriesDescriptorsListCall {
	c.opt_["oldest"] = oldest
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *TimeseriesDescriptorsListCall) PageToken(pageToken string) *TimeseriesDescriptorsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Timespan sets the optional parameter "timespan": Length of the time
// interval to query, which is an alternative way to declare the
// interval: (youngest - timespan, youngest]. The timespan and oldest
// parameters should not be used together. Units:
// - s: second
// - m: minute
// - h: hour
// - d: day
// - w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for
// example: 2w3d is not allowed; you should use 17d instead.
//
// If neither oldest nor timespan is specified, the default time
// interval will be (youngest - 4 hours, youngest].
func (c *TimeseriesDescriptorsListCall) Timespan(timespan string) *TimeseriesDescriptorsListCall {
	c.opt_["timespan"] = timespan
	return c
}

// Window sets the optional parameter "window": The sampling window. At
// most one data point will be returned for each window in the requested
// time interval. This parameter is only valid for non-cumulative metric
// types. Units:
// - m: minute
// - h: hour
// - d: day
// - w: week  Examples: 3m, 4w. Only one unit is allowed, for example:
// 2w3d is not allowed; you should use 17d instead.
func (c *TimeseriesDescriptorsListCall) Window(window string) *TimeseriesDescriptorsListCall {
	c.opt_["window"] = window
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesDescriptorsListCall) Fields(s ...googleapi.Field) *TimeseriesDescriptorsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TimeseriesDescriptorsListCall) Do() (*ListTimeseriesDescriptorsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("youngest", fmt.Sprintf("%v", c.youngest))
	if v, ok := c.opt_["aggregator"]; ok {
		params.Set("aggregator", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["count"]; ok {
		params.Set("count", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["labels"]; ok {
		params.Set("labels", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["oldest"]; ok {
		params.Set("oldest", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["timespan"]; ok {
		params.Set("timespan", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["window"]; ok {
		params.Set("window", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/timeseriesDescriptors/{metric}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTimeseriesDescriptorsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the descriptors of the time series that match the metric and labels values and that have data points in the interval. Large responses are paginated; use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.timeseriesDescriptors.list",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "aggregator": {
	//       "description": "The aggregation function that will reduce the data points in each window to a single point. This parameter is only valid for non-cumulative metrics with a value type of INT64 or DOUBLE.",
	//       "enum": [
	//         "max",
	//         "mean",
	//         "min",
	//         "sum"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "count": {
	//       "default": "100",
	//       "description": "Maximum number of time series descriptors per page. Used for pagination. If not specified, count = 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "description": "A collection of labels for the matching time series, which are represented as:  \n- key==value: key equals the value \n- key=~value: key regex matches the value \n- key!=value: key does not equal the value \n- key!~value: key regex does not match the value  For example, to list all of the time series descriptors for the region us-central1, you could specify:\nlabel=cloud.googleapis.com%2Flocation=~us-central1.*",
	//       "location": "query",
	//       "pattern": "(.+?)(==|=~|!=|!~)(.+)",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "metric": {
	//       "description": "Metric names are protocol-free URLs as listed in the Supported Metrics page. For example, compute.googleapis.com/instance/disk/read_ops_count.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "description": "Start of the time interval (exclusive), which is expressed as an RFC 3339 timestamp. If neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest]",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which this time series belongs. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "description": "Length of the time interval to query, which is an alternative way to declare the interval: (youngest - timespan, youngest]. The timespan and oldest parameters should not be used together. Units:  \n- s: second \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.\n\nIf neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest].",
	//       "location": "query",
	//       "pattern": "[0-9]+[smhdw]?",
	//       "type": "string"
	//     },
	//     "window": {
	//       "description": "The sampling window. At most one data point will be returned for each window in the requested time interval. This parameter is only valid for non-cumulative metric types. Units:  \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.",
	//       "location": "query",
	//       "pattern": "[0-9]+[mhdw]?",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "description": "End of the time interval (inclusive), which is expressed as an RFC 3339 timestamp.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseriesDescriptors/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}
