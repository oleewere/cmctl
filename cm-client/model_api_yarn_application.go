/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Represents a Yarn application
type ApiYarnApplication struct {
	// The sum of memory in MB allocated to the application's running containers Available since v12.
	AllocatedMB float32 `json:"allocatedMB,omitempty"`
	// The sum of virtual cores allocated to the application's running containers Available since v12.
	AllocatedVCores float32 `json:"allocatedVCores,omitempty"`
	// The number of containers currently running for the application Available since v12.
	RunningContainers float32 `json:"runningContainers,omitempty"`
	// List of YARN application tags. Available since v12.
	ApplicationTags []string `json:"applicationTags,omitempty"`
	// Allocated memory to the application in units of mb-secs. Available since v12.
	AllocatedMemorySeconds float32 `json:"allocatedMemorySeconds,omitempty"`
	// Allocated vcore-secs to the application. Available since v12.
	AllocatedVcoreSeconds float32 `json:"allocatedVcoreSeconds,omitempty"`
	// The application id.
	ApplicationId string `json:"applicationId,omitempty"`
	// The name of the application.
	Name string `json:"name,omitempty"`
	// The time the application was submitted.
	StartTime string `json:"startTime,omitempty"`
	// The time the application finished. If the application hasn't finished this will return null.
	EndTime string `json:"endTime,omitempty"`
	// The user who submitted the application.
	User string `json:"user,omitempty"`
	// The pool the application was submitted to.
	Pool string `json:"pool,omitempty"`
	// The progress, as a percentage, the application has made. This is only set if the application is currently executing.
	Progress float32 `json:"progress,omitempty"`
	// A map of additional application attributes which is generated by Cloudera Manager. For example MR2 job counters are exposed as key/value pairs here. For more details see the Cloudera Manager documentation.
	Attributes map[string]string `json:"attributes,omitempty"`
	// 
	Mr2AppInformation *ApiMr2AppInformation `json:"mr2AppInformation,omitempty"`
	// 
	State string `json:"state,omitempty"`
	// Actual memory (in MB-secs) used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12.
	ContainerUsedMemorySeconds float32 `json:"containerUsedMemorySeconds,omitempty"`
	// Maximum memory used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics Available since v16
	ContainerUsedMemoryMax float32 `json:"containerUsedMemoryMax,omitempty"`
	// Actual CPU (in percent-secs) used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12.
	ContainerUsedCpuSeconds float32 `json:"containerUsedCpuSeconds,omitempty"`
	// Actual VCore-secs used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12.
	ContainerUsedVcoreSeconds float32 `json:"containerUsedVcoreSeconds,omitempty"`
	// Total memory (in mb-secs) allocated to containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12.
	ContainerAllocatedMemorySeconds float32 `json:"containerAllocatedMemorySeconds,omitempty"`
	// Total vcore-secs allocated to containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12.
	ContainerAllocatedVcoreSeconds float32 `json:"containerAllocatedVcoreSeconds,omitempty"`
}
