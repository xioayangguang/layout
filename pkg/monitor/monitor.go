package monitor

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"os"
	"runtime/pprof"
	"time"
)

func init() {
	go func() {
		_ = os.MkdirAll("./log/pprof/", os.ModePerm)
		var monitorStartTime int64 = 0
		var monitorState = true
		interval := time.Second * 5
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			cpuPercent, err := cpu.Percent(interval, false)
			if err != nil {
				continue
			}
			v, err := mem.VirtualMemory()
			if err != nil {
				continue
			}
			if monitorState {
				if v.UsedPercent >= 80 || cpuPercent[0] >= 80 {
					cpuFile, err := os.OpenFile(fmt.Sprintf("./log/pprof/%s-cpu%v.pprof", time.Now().Format("01-02-15-4-5"), int(cpuPercent[0])), os.O_CREATE, 0666)
					if err != nil {
						return
					}
					_ = pprof.StartCPUProfile(cpuFile)
					memFile, err := os.OpenFile(fmt.Sprintf("./log/pprof/%s-mem%v.pprof", time.Now().Format("01-02-15-4-5"), v.UsedPercent), os.O_CREATE, 0666)
					if err != nil {
						return
					}
					_ = pprof.WriteHeapProfile(memFile)
					_ = cpuFile.Close()
					_ = memFile.Close()
					monitorStartTime = time.Now().Unix()
					monitorState = false
				}
			}
			if !monitorState {
				if (v.UsedPercent < 80 && cpuPercent[0] < 80) || (monitorStartTime+30 < time.Now().Unix()) {
					pprof.StopCPUProfile()
					monitorStartTime = 0
					monitorState = true
				}
			}
		}
	}()
}
