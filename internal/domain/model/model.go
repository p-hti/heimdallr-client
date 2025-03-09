package model

type Resource struct {
	CPUUsage     float64 `json:"cpu_usage"`
	GPUUsage     float64 `json:"gpu_usage,omitempty"`
	MemoryUsage  float64 `json:"memory_usage"`
	DiskUsage    float64 `json:"disk_usage"`
	NetworkUsage uint64  `json:"network_usage"`
}

type Features struct {
	CPU       float64 `json:"CPU"`
	GLU       float64 `json:"GPU,omitempty"`
	Memory    float64 `json:"Memory"`
	DiskSpace float64 `json:"Disk"`
}

type Machine struct {
	OS            string `json:"os"`
	Hostname      string `json:"hostname"`
	IP            string `json:"ip"`
	UsageResource Resource
	Features      Features
}
