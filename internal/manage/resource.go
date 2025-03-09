package manage

import (
	"github.com/p-hti/heimdallr-client/internal/domain/model"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func GetResourceUsage() (model.Resource, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return model.Resource{}, err
	}

	memPercent, err := mem.VirtualMemory()
	if err != nil {
		return model.Resource{}, err
	}

	diskPercent, err := disk.Usage("/")
	if err != nil {
		return model.Resource{}, err
	}

	network, err := net.IOCounters(false)
	if err != nil {
		return model.Resource{}, err
	}

	resources := model.Resource{
		CPUUsage:     cpuPercent[0],
		MemoryUsage:  memPercent.UsedPercent,
		DiskUsage:    diskPercent.UsedPercent,
		NetworkUsage: network[0].BytesSent,
	}
	return resources, nil
}

func GetMachineInfo() (model.Machine, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return model.Machine{}, err
	}
	return model.Machine{}, nil
}
