package machine_info

import (
	"DistriAI-Node/machine_info/cpu"
	"DistriAI-Node/machine_info/disk"
	"DistriAI-Node/machine_info/flops"
	"DistriAI-Node/machine_info/gpu"
	"DistriAI-Node/machine_info/location"
	"DistriAI-Node/machine_info/machine_uuid"
	"DistriAI-Node/machine_info/memory"
	"DistriAI-Node/machine_info/speedtest"
	logs "DistriAI-Node/utils/log_utils"
)

type MachineInfo struct {
	MachineUUID     machine_uuid.MachineUUID `json:"MachineUUID"`
	Addr            string                   `json:"Addr"`
	CPUInfo         cpu.InfoCPU              `json:"CPUInfo"`
	DiskInfo        disk.InfoDisk            `json:"DiskInfo"`
	Score           float64                  `json:"Score"`
	MemoryInfo      memory.InfoMemory        `json:"InfoMemory"`
	GPUInfo         gpu.InfoGPU              `json:"GPUInfo"`
	LocationInfo    location.InfoLocation    `json:"LocationInfo"`
	SpeedInfo       speedtest.InfoSpeed      `json:"SpeedInfo"`
	FlopsInfo       flops.InfoFlop           `json:"InfoFlop"`
	MachineAccounts string                   `json:"MachineAccounts"`
}

func GetMachineInfo() (MachineInfo, error) {
	var hwInfo MachineInfo

	cpuInfo, err := cpu.GetCPUInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.CPUInfo = cpuInfo

	memInfo, err := memory.GetMemoryInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.MemoryInfo = memInfo

	gpuInfo, _ := gpu.GetGPUInfo()
	// if err != nil {
	// 	return hwInfo, err
	// }
	hwInfo.GPUInfo = gpuInfo

	locationInfo, err := location.GetLocationInfo()
	if err != nil {
		return hwInfo, err
	}
	hwInfo.LocationInfo = locationInfo

	// Easy debugging
	speedInfo, err := speedtest.GetSpeedInfo()
	if err != nil {
		logs.Error(err.Error())
	}
	// speedInfo := speedtest.InfoSpeed{
	// 	Download: "1000 Mbit/s",
	// 	Upload:   "1000 Mbit/s",
	// }
	hwInfo.SpeedInfo = speedInfo

	flopsInfo := flops.GetFlopsInfo(gpuInfo.Model)
	hwInfo.FlopsInfo = flopsInfo

	return hwInfo, nil
}
