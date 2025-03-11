package manage

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	netUtil "github.com/shirou/gopsutil/net"
)

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

func NewMachine() (*Machine, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return &Machine{}, err
	}

	// conn, err := net.Dial("udp", "8.8.8.8:80")
	// if err != nil {
	// 	return &Machine{}, err
	// }
	// defer conn.Close()

	// localAddr := conn.LocalAddr().(*net.UDPAddr)
	// ip := localAddr.IP.String()
	machine := &Machine{
		OS:       hostInfo.OS,
		Hostname: hostInfo.Hostname,
		IP:       "43242352",
	}
	err = machine.GetFeatures()
	if err != nil {
		return &Machine{}, err
	}

	return machine, err
}

func (m *Machine) GetFeatures() error {
	physicalCnt, err := cpu.Counts(false)
	if err != nil {
		return err
	}
	logicalCnt, err := cpu.Counts(true)
	if err != nil {
		return err
	}

	mem, _ := mem.VirtualMemory()
	// fmt.Printf("%.1f", float64(mem.Total)/1024/1024/1024)

	features := Features{
		PhysicalCores: physicalCnt,
		LogicalCores:  logicalCnt,
		Memory:        float64(mem.Total) / 1024 / 1024 / 1024,
	}

	m.Features = features
	return nil
}

func (m *Machine) GetResourceUsage() error {
	memPercent, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	diskPercent, err := disk.Usage("/")
	if err != nil {
		return err
	}

	network, err := netUtil.IOCounters(false)
	if err != nil {
		return err
	}
	totalPercent, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		return err
	}
	// perPercents, err := cpu.Percent(3*time.Second, true)
	// if err != nil {
	// 	return err
	// }

	resources := Resource{
		CPUUsage: totalPercent[0],
		// CoreLoad:     perPercents,
		MemoryUsage:  memPercent.UsedPercent,
		DiskUsage:    diskPercent.UsedPercent,
		NetworkUsage: network[0].BytesSent,
	}
	m.UsageResource = resources
	return nil
}
