package manage

import (
	"math"
	"net"
	"time"

	"github.com/p-hti/heimdallr-client/internal/domain/model"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	netUtil "github.com/shirou/gopsutil/net"
)

type Machine struct {
	model.Machine
}

func NewMachine() (*Machine, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return &Machine{}, err
	}

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return &Machine{}, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()
	machine := &Machine{
		Machine: model.Machine{
			OS:       hostInfo.OS,
			Hostname: hostInfo.Hostname,
			IP:       ip,
		},
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

	memory := math.Floor((float64(mem.Total)/1024/1024/1024)*10) / 10
	features := model.Features{
		PhysicalCores: physicalCnt,
		LogicalCores:  logicalCnt,
		Memory:        memory,
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
	totalPercent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return err
	}
	perPercents, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return err
	}

	resources := model.Resource{
		CPUUsage:     totalPercent[0],
		CoreLoad:     perPercents,
		MemoryUsage:  memPercent.UsedPercent,
		DiskUsage:    diskPercent.UsedPercent,
		NetworkUsage: network[0].BytesSent,
	}
	m.UsageResource = resources
	return nil
}
