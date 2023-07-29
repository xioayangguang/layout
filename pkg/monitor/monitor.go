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
		var cpuFile *os.File
		var memFile *os.File
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
					cpuFile, err = os.OpenFile(fmt.Sprintf("./log/pprof/%s-cpu%v.pprof", time.Now().Format("01-02-15-4-5"), int(cpuPercent[0])), os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						fmt.Println(err)
						continue
					}
					err = pprof.StartCPUProfile(cpuFile)
					if err != nil {
						fmt.Println(err)
						continue
					}
					memFile, err = os.OpenFile(fmt.Sprintf("./log/pprof/%s-mem%v.pprof", time.Now().Format("01-02-15-4-5"), int(v.UsedPercent)), os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						fmt.Println(err)
						continue
					}
					err = pprof.WriteHeapProfile(memFile)
					if err != nil {
						fmt.Println(err)
						continue
					}
					monitorStartTime = time.Now().Unix()
					monitorState = false
					fmt.Println("开始")
				}
			}
			if !monitorState {
				if (v.UsedPercent < 80 && cpuPercent[0] < 80) || (monitorStartTime+30 < time.Now().Unix()) {
					pprof.StopCPUProfile()
					monitorStartTime = 0
					monitorState = true
					err = cpuFile.Close()
					if err != nil {
						fmt.Println(err)
						continue
					}
					err = memFile.Close()
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Println("结束")
				}
			}
		}
	}()
}
