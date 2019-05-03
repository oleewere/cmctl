/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Request object containing information needed for querying timeseries data.  Available since API v11.
type ApiTimeSeriesRequest struct {
	// tsquery to run against the CM time-series data store. Please see the <a href=\"http://tiny.cloudera.com/cm_tsquery\"> tsquery language documentation</a>.<p/>
	Query string `json:"query,omitempty"`
	// Start of the period to query in ISO 8601 format (defaults to 5 minutes before the end of the period).
	From string `json:"from,omitempty"`
	// End of the period to query in ISO 8601 format (defaults to current time).
	To string `json:"to,omitempty"`
	// contentType to return the response in. The content types \"application/json\" and \"text/csv\" are supported. This defaults to \"application/json\". If \"text/csv\" is specified then we return one row per time series data point, and we don't return any of the metadata.
	ContentType string `json:"contentType,omitempty"`
	// Aggregate rollup level desired for the response data. Valid values are RAW, TEN_MINUTELY, HOURLY, SIX_HOURLY, DAILY, and WEEKLY. Note that if the mustUseDesiredRollup parameter is not set, then the monitoring server can decide to return a different rollup level.
	DesiredRollup string `json:"desiredRollup,omitempty"`
	// If set to true, then the tsquery will return data with the desired aggregate rollup level.
	MustUseDesiredRollup bool `json:"mustUseDesiredRollup,omitempty"`
}
