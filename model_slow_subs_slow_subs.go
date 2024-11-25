/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SlowSubsSlowSubs struct {
	// Enable this feature
	Enable bool `json:"enable,omitempty"`
	// The latency threshold for statistics
	Threshold string `json:"threshold,omitempty"`
	// The eviction time of the record, which in the statistics record table
	ExpireInterval string `json:"expire_interval,omitempty"`
	// The maximum number of records in the slow subscription statistics record table
	TopKNum int32 `json:"top_k_num,omitempty"`
	// The method to calculate the latency
	StatsType string `json:"stats_type,omitempty"`
}
