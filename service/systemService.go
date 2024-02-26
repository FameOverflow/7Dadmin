package service

import (
	"7D-admin/utils/systemUtils"
	"log"
	"runtime"
	"sync"
)

type SystemInfo struct {
	HostInfo      *systemUtils.HostInfo `json:"host"`
	CpuInfo       *systemUtils.CpuInfo  `json:"cpu"`
	MemInfo       *systemUtils.MemInfo  `json:"mem"`
	DiskInfo      *systemUtils.DiskInfo `json:"disk"`
	PanelMemUsage uint64                `json:"panelMemUsage"`
	PanelCpuUsage float64               `json:"panelCpuUsage"`
}

func (g *GameService) GetSystemInfo(clusterName string) *SystemInfo {
	var wg sync.WaitGroup
	wg.Add(5)

	dashboardVO := SystemInfo{}
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		dashboardVO.HostInfo = systemUtils.GetHostInfo()
	}()

	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		dashboardVO.CpuInfo = systemUtils.GetCpuInfo()
	}()

	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		dashboardVO.MemInfo = systemUtils.GetMemInfo()
	}()

	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		dashboardVO.DiskInfo = systemUtils.GetDiskInfo()
	}()

	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		dashboardVO.PanelMemUsage = m.Alloc / 1024 // 将字节转换为MB

		// 获取当前程序使用的CPU信息
		//startCPU, _ := cpu.Percent(time.Second, false)
		//time.Sleep(1 * time.Second) // 假设程序运行1秒
		//endCPU, _ := cpu.Percent(time.Second, false)
		//cpuUsage := endCPU[0] - startCPU[0]
		//dashboardVO.PanelCpuUsage = cpuUsage

	}()

	wg.Wait()
	return &dashboardVO
}