// Package tools 存放辅助方法
package tools

import (
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"

	"github.com/TheTNB/panel/pkg/shell"
	"github.com/TheTNB/panel/pkg/types"
)

// CurrentInfo 获取监控数据
func CurrentInfo(nets, disks []string) types.CurrentInfo {
	var res types.CurrentInfo
	res.Cpus, _ = cpu.Info()
	res.Percents, _ = cpu.Percent(100*time.Millisecond, true)
	percent, _ := cpu.Percent(100*time.Millisecond, false)
	if len(percent) > 0 {
		res.Percent = percent[0]
	}
	res.Load, _ = load.Avg()
	res.Host, _ = host.Info()
	res.Mem, _ = mem.VirtualMemory()
	res.Swap, _ = mem.SwapMemory()
	// 硬盘IO
	ioCounters, _ := disk.IOCounters(disks...)
	for _, info := range ioCounters {
		res.DiskIO = append(res.DiskIO, info)
	}
	// 硬盘使用
	var excludes = []string{"/dev", "/boot", "/sys", "/dev", "/run", "/proc", "/usr", "/var", "/snap"}
	excludes = append(excludes, "/mnt/cdrom") // CDROM
	excludes = append(excludes, "/mnt/wsl")   // Windows WSL
	res.Disk, _ = disk.Partitions(false)
	res.Disk = slices.DeleteFunc(res.Disk, func(d disk.PartitionStat) bool {
		for _, exclude := range excludes {
			if strings.HasPrefix(d.Mountpoint, exclude) {
				return true
			}
			// 去除内存盘和overlay容器盘
			if slices.Contains([]string{"tmpfs", "overlay"}, d.Fstype) {
				continue
			}
		}
		return false
	})
	// 分区使用
	for _, partition := range res.Disk {
		usage, _ := disk.Usage(partition.Mountpoint)
		res.DiskUsage = append(res.DiskUsage, *usage)
	}
	// 网络
	if len(nets) == 0 {
		netInfo, _ := net.IOCounters(false)
		res.Net = netInfo
	} else {
		var netStats []net.IOCountersStat
		netInfo, _ := net.IOCounters(true)
		for _, state := range netInfo {
			if slices.Contains(nets, state.Name) {
				netStats = append(netStats, state)
			}
		}
		res.Net = netStats
	}

	res.Time = time.Now()
	return res
}

// RestartPanel 重启面板
func RestartPanel() {
	_ = shell.ExecfAsync("sleep 1 && systemctl restart panel")
}

// IsChina 是否中国大陆
func IsChina() bool {
	client := resty.New()
	client.SetTimeout(5 * time.Second)
	client.SetRetryCount(2)

	resp, err := client.R().Get("https://www.cloudflare-cn.com/cdn-cgi/trace")
	if err != nil || !resp.IsSuccess() {
		return false
	}

	if strings.Contains(resp.String(), "loc=CN") {
		return true
	}

	return false
}

// GetPublicIP 获取公网IP
func GetPublicIP() (string, error) {
	client := resty.New()
	client.SetTimeout(5 * time.Second)
	client.SetRetryCount(2)

	resp, err := client.R().Get("https://www.cloudflare-cn.com/cdn-cgi/trace")
	if err != nil || !resp.IsSuccess() {
		return "", errors.New("获取公网IP失败")
	}

	return strings.TrimPrefix(strings.Split(resp.String(), "\n")[2], "ip="), nil
}
