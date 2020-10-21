package sd

import (
	"errors"
	"net/http"
	"time"

	"github.com/lexkong/log"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/viper"

	"github.com/shirou/gopsutil/cpu"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
)

type diskMonitor struct {
	Used    int `json:"UsedDisk"`
	Total   int `json:"TotalDisk"`
	Level   int `json:"WarnLevel"`
	Percent int `json:"UsedPercent"`
}

type cpuMonitor struct {
	Level   int `json:"WarnLevel"`
	Percent int `json:"UsedPercent"`
}

type memMonitor struct {
	Used    int `json:"UsedMemory"`
	Total   int `json:"TotalMemory"`
	Level   int `json:"WarnLevel"`
	Percent int `json:"UsedPercent"`
}

type serverMonitor struct {
	Disk diskMonitor `json:"Disk"`
	Cpu  cpuMonitor  `json:"CPU"`
	Mem  memMonitor  `json:"Memory"`
}

func DiskCheck() diskMonitor {
	diskData, _ := disk.Usage("/")
	ret := diskMonitor{
		Used:    int(diskData.Used) / MB,
		Total:   int(diskData.Total) / MB,
		Percent: int(diskData.UsedPercent),
		Level:   int(0),
	}
	if ret.Percent >= 95 {
		ret.Level = 2
	} else if ret.Percent >= 90 {
		ret.Level = 1
	}
	return ret
}

func CPUCheck() cpuMonitor {
	cpuData, _ := cpu.Percent(time.Second, false)
	ret := cpuMonitor{
		Percent: int(cpuData[0]),
		Level:   int(0),
	}
	if ret.Percent >= 95 {
		ret.Level = 2
	} else if ret.Percent >= 90 {
		ret.Level = 1
	}
	return ret
}

func MemCheck() memMonitor {
	memData, _ := mem.VirtualMemory()
	ret := memMonitor{
		Used:    int(memData.Used) / MB,
		Total:   int(memData.Total) / MB,
		Percent: int(memData.UsedPercent),
		Level:   int(0),
	}
	if ret.Percent >= 95 {
		ret.Level = 2
	} else if ret.Percent >= 90 {
		ret.Level = 1
	}
	return ret
}

func HealthCheck(ctx *gin.Context) {
	message := "Status OK!"
	ctx.String(http.StatusOK, message)
}

func MonitorCheck(ctx *gin.Context) {
	data := serverMonitor{}

	data.Disk = DiskCheck()
	data.Cpu = CPUCheck()
	data.Mem = MemCheck()
	status := http.StatusOK
	maxLevel := func(x, y, z int) int {
		if x >= y && x >= z {
			return x
		} else if y >= z {
			return y
		} else {
			return z
		}
	}(data.Disk.Percent, data.Cpu.Percent, data.Mem.Percent)
	if maxLevel == 2 {
		status = http.StatusInternalServerError
	} else if maxLevel == 1 {
		status = http.StatusTooManyRequests
	}
	ctx.JSON(status, data)
}

func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
