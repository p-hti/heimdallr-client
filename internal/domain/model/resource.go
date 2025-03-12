package model

type Resource struct {
	CPUUsage     float64   `json:"cpu_usage"`
	CoreLoad     []float64 `json:"core_load"`
	GPUUsage     float64   `json:"gpu_usage,omitempty"`
	MemoryUsage  float64   `json:"memory_usage"`
	DiskUsage    float64   `json:"disk_usage"`
	NetworkUsage uint64    `json:"network_usage"`
}

type Features struct {
	PhysicalCores int     `json:"physical_cores"`
	LogicalCores  int     `json:"logical_cores"`
	GPU           float64 `json:"gpu,omitempty"`
	Memory        float64 `json:"memory"`
	DiskSpace     float64 `json:"disk,omitempty"`
}

type Machine struct {
	OS            string   `json:"os"`
	Hostname      string   `json:"hostname"`
	IP            string   `json:"ip"`
	UsageResource Resource `json:"usage_resource,omitempty"`
	Features      Features `json:"features,omitempty"`
}
