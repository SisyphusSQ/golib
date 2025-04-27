package ps

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	Pid            int32                    `json:"pid"`
	CpuUsage       float64                  `json:"cpuUsage"`
	MemUsage       float32                  `json:"memUsage"`
	NumFDs         int32                    `json:"numFDs"`
	NumThreads     int32                    `json:"numThreads"`
	Connections    []net.ConnectionStat     `json:"connections"`
	MemoryInfoStat *process.MemoryInfoStat  `json:"memoryInfoStat"`
	IOCountersStat *process.IOCountersStat  `json:"ioCountersStat"`
	OpenFilesStats []process.OpenFilesStat  `json:"openFilesStat"`
	Threads        map[int32]*cpu.TimesStat `json:"threads"`

	ps *process.Process
}

func GetPsInfoWithoutThrds(pid int32) (*ProcessInfo, []error) {
	var (
		err    error
		errs   = make([]error, 0)
		psInfo = &ProcessInfo{
			Pid: pid,
		}
	)

	psInfo.CpuUsage, err = getProcessCPUUsage(pid)
	if err != nil {
		err = fmt.Errorf("cpuUsage err: %w", err)
		errs = append(errs, err)
	}

	p, err := process.NewProcess(pid)
	if err != nil {
		err = fmt.Errorf("newProcess err: %w", err)
		errs = append(errs, err)
		return nil, errs
	}
	psInfo.ps = p

	psInfo.MemUsage, err = p.MemoryPercent()
	if err != nil {
		err = fmt.Errorf("memoryPercent err: %w", err)
		errs = append(errs, err)
	}

	psInfo.MemoryInfoStat, err = p.MemoryInfo()
	if err != nil {
		err = fmt.Errorf("memoryInfo err: %w", err)
		errs = append(errs, err)
	}

	psInfo.Connections, err = p.Connections()
	if err != nil {
		err = fmt.Errorf("connections err: %w", err)
		errs = append(errs, err)
	}

	psInfo.NumFDs, err = p.NumFDs()
	if err != nil {
		err = fmt.Errorf("numFDs err: %w", err)
		errs = append(errs, err)
	}

	psInfo.NumThreads, err = p.NumThreads()
	if err != nil {
		err = fmt.Errorf("numThreads err: %w", err)
		errs = append(errs, err)
	}

	psInfo.IOCountersStat, err = p.IOCounters()
	if err != nil {
		err = fmt.Errorf("ioCountersStat err: %w", err)
		errs = append(errs, err)
	}

	psInfo.OpenFilesStats, err = p.OpenFiles()
	if err != nil {
		err = fmt.Errorf("openFiles err: %w", err)
		errs = append(errs, err)
	}

	return psInfo, errs
}

func GetPsInfo(pid int32) (*ProcessInfo, []error) {
	var err error
	psInfo, errs := GetPsInfoWithoutThrds(pid)
	if psInfo == nil {
		return nil, errs
	}

	psInfo.Threads, err = psInfo.ps.Threads()
	if err != nil {
		err = fmt.Errorf("threads err: %w", err)
		errs = append(errs, err)
	}
	return psInfo, errs
}

func getProcessCPUUsage(pid int32) (float64, error) {
	cmd := exec.Command("top", "-b", "-n", "1", "-p", strconv.Itoa(int(pid)))
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, strconv.Itoa(int(pid))) {
			fields := strings.Fields(line)
			if len(fields) > 8 {
				cpuUsage, err := strconv.ParseFloat(fields[8], 64)
				if err != nil {
					return 0, err
				}
				return cpuUsage, nil
			}
		}
	}

	return 0, fmt.Errorf("failed to find CPU usage for PID %d", pid)
}
