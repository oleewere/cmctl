/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client

// Utilization report information of a tenant of Yarn application.
type ApiYarnTenantUtilization struct {
	// Name of the tenant.
	TenantName string `json:"tenantName,omitempty"`
	// Average number of VCores allocated to YARN applications of the tenant.
	AvgYarnCpuAllocation float32 `json:"avgYarnCpuAllocation,omitempty"`
	// Average number of VCores used by YARN applications of the tenant.
	AvgYarnCpuUtilization float32 `json:"avgYarnCpuUtilization,omitempty"`
	// Average unused VCores of the tenant.
	AvgYarnCpuUnusedCapacity float32 `json:"avgYarnCpuUnusedCapacity,omitempty"`
	// Average steady fair share VCores.
	AvgYarnCpuSteadyFairShare float32 `json:"avgYarnCpuSteadyFairShare,omitempty"`
	// Average allocated Vcores with pending containers.
	AvgYarnPoolAllocatedCpuDuringContention float32 `json:"avgYarnPoolAllocatedCpuDuringContention,omitempty"`
	// Average fair share VCores with pending containers.
	AvgYarnPoolFairShareCpuDuringContention float32 `json:"avgYarnPoolFairShareCpuDuringContention,omitempty"`
	// Average steady fair share VCores with pending containers.
	AvgYarnPoolSteadyFairShareCpuDuringContention float32 `json:"avgYarnPoolSteadyFairShareCpuDuringContention,omitempty"`
	// Average percentage of pending containers for the pool during periods of contention.
	AvgYarnContainerWaitRatio float32 `json:"avgYarnContainerWaitRatio,omitempty"`
	// Average memory allocated to YARN applications of the tenant.
	AvgYarnMemoryAllocation float32 `json:"avgYarnMemoryAllocation,omitempty"`
	// Average memory used by YARN applications of the tenant.
	AvgYarnMemoryUtilization float32 `json:"avgYarnMemoryUtilization,omitempty"`
	// Average unused memory of the tenant.
	AvgYarnMemoryUnusedCapacity float32 `json:"avgYarnMemoryUnusedCapacity,omitempty"`
	// Average steady fair share memory.
	AvgYarnMemorySteadyFairShare float32 `json:"avgYarnMemorySteadyFairShare,omitempty"`
	// Average allocated memory with pending containers.
	AvgYarnPoolAllocatedMemoryDuringContention float32 `json:"avgYarnPoolAllocatedMemoryDuringContention,omitempty"`
	// Average fair share memory with pending containers.
	AvgYarnPoolFairShareMemoryDuringContention float32 `json:"avgYarnPoolFairShareMemoryDuringContention,omitempty"`
	// Average steady fair share memory with pending containers.
	AvgYarnPoolSteadyFairShareMemoryDuringContention float32 `json:"avgYarnPoolSteadyFairShareMemoryDuringContention,omitempty"`
}