package service

import (
	"github.com/saichler/netop/go/types"
)

// GetMockDevices returns a comprehensive list of mock network devices
func GetMockDevices() []*types.Device {
	devices := []*types.Device{
		{
			Id:              "SW-001",
			Name:            "Core Switch A",
			Type:            types.DeviceType_DEVICE_TYPE_SWITCH,
			Status:          types.DeviceStatus_DEVICE_STATUS_ONLINE,
			IpAddress:       "192.168.1.10",
			Location:        "Datacenter A",
			Model:           "Cisco Catalyst 9300",
			SerialNumber:    "FCW2148G0AB",
			FirmwareVersion: "IOS-XE 17.6.4",
			MacAddress:      "00:1B:54:C2:3A:4E",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.1.1",
			Vlan:            "10",
			Uptime:          "45d 12h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    48.5,
				MemoryUsagePercent: 62.3,
				TemperatureCelsius: 42.1,
			},
			Ports: []*types.Port{
				{
					Name:   "Uplink Port 1",
					Status: types.PortStatus_PORT_STATUS_ACTIVE,
					Speed:  "1000",
					Vlan:   "10",
				},
				{
					Name:   "Uplink Port 2",
					Status: types.PortStatus_PORT_STATUS_ACTIVE,
					Speed:  "1000",
					Vlan:   "10",
				},
				{
					Name:   "Management Port",
					Status: types.PortStatus_PORT_STATUS_ACTIVE,
					Speed:  "1000",
					Vlan:   "1",
				},
			},
		},
		{
			Id:              "SW-002",
			Name:            "Core Switch B",
			Type:            types.DeviceType_DEVICE_TYPE_SWITCH,
			Status:          types.DeviceStatus_DEVICE_STATUS_ONLINE,
			IpAddress:       "192.168.1.11",
			Location:        "Datacenter B",
			Model:           "Cisco Catalyst 9300",
			SerialNumber:    "FCW2149G0AB",
			FirmwareVersion: "IOS-XE 17.6.4",
			MacAddress:      "00:1B:54:C2:3A:4F",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.1.1",
			Vlan:            "10",
			Uptime:          "30d 5h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    45.2,
				MemoryUsagePercent: 58.7,
				TemperatureCelsius: 41.8,
			},
		},
		{
			Id:              "RT-042",
			Name:            "Edge Router 42",
			Type:            types.DeviceType_DEVICE_TYPE_ROUTER,
			Status:          types.DeviceStatus_DEVICE_STATUS_WARNING,
			IpAddress:       "192.168.2.1",
			Location:        "Building B",
			Model:           "Cisco ISR 4331",
			SerialNumber:    "FJC2148W0LA",
			FirmwareVersion: "IOS-XE 16.12.4",
			MacAddress:      "00:1A:2B:3C:4D:42",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.2.254",
			Vlan:            "20",
			Uptime:          "12d 8h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    87.3,
				MemoryUsagePercent: 78.9,
				TemperatureCelsius: 58.2,
			},
		},
		{
			Id:              "RT-001",
			Name:            "Main Router",
			Type:            types.DeviceType_DEVICE_TYPE_ROUTER,
			Status:          types.DeviceStatus_DEVICE_STATUS_ONLINE,
			IpAddress:       "192.168.0.254",
			Location:        "Datacenter A",
			Model:           "Cisco ISR 4451",
			SerialNumber:    "FJC2149W0LB",
			FirmwareVersion: "IOS-XE 16.12.5",
			MacAddress:      "00:1A:2B:3C:4D:5E",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.0.1",
			Vlan:            "1",
			Uptime:          "90d 12h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    35.1,
				MemoryUsagePercent: 52.4,
				TemperatureCelsius: 38.7,
			},
		},
		{
			Id:              "FW-003",
			Name:            "Main Firewall",
			Type:            types.DeviceType_DEVICE_TYPE_FIREWALL,
			Status:          types.DeviceStatus_DEVICE_STATUS_ONLINE,
			IpAddress:       "192.168.0.1",
			Location:        "Datacenter A",
			Model:           "Fortinet FortiGate 200F",
			SerialNumber:    "FG200FTK20012346",
			FirmwareVersion: "FortiOS 7.2.5",
			MacAddress:      "00:09:0F:09:00:03",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.0.254",
			Vlan:            "1",
			Uptime:          "120d 5h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    28.6,
				MemoryUsagePercent: 44.2,
				TemperatureCelsius: 36.9,
			},
		},
		{
			Id:              "AP-101",
			Name:            "WiFi AP Floor 1",
			Type:            types.DeviceType_DEVICE_TYPE_ACCESS_POINT,
			Status:          types.DeviceStatus_DEVICE_STATUS_ONLINE,
			IpAddress:       "192.168.10.101",
			Location:        "Building A",
			Model:           "Cisco Aironet 2802i",
			SerialNumber:    "FCZ2049W0XY",
			FirmwareVersion: "AireOS 8.10.151.0",
			MacAddress:      "00:3A:9A:47:C5:B2",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.10.1",
			Vlan:            "100",
			Uptime:          "7d 3h",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    22.1,
				MemoryUsagePercent: 31.8,
				TemperatureCelsius: 34.2,
			},
		},
		{
			Id:              "FW-002",
			Name:            "Backup Firewall",
			Type:            types.DeviceType_DEVICE_TYPE_FIREWALL,
			Status:          types.DeviceStatus_DEVICE_STATUS_MAINTENANCE,
			IpAddress:       "192.168.0.2",
			Location:        "Datacenter B",
			Model:           "Fortinet FortiGate 100F",
			SerialNumber:    "FG100FTK20012345",
			FirmwareVersion: "FortiOS 7.2.4",
			MacAddress:      "00:09:0F:09:00:02",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.0.1",
			Vlan:            "1",
			Uptime:          "N/A",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    0,
				MemoryUsagePercent: 0,
				TemperatureCelsius: 0,
			},
		},
		{
			Id:              "SW-015",
			Name:            "Edge Switch 15",
			Type:            types.DeviceType_DEVICE_TYPE_SWITCH,
			Status:          types.DeviceStatus_DEVICE_STATUS_OFFLINE,
			IpAddress:       "192.168.3.15",
			Location:        "Building C",
			Model:           "Cisco Catalyst 2960X",
			SerialNumber:    "FOC2049Y0ZB",
			FirmwareVersion: "IOS 15.2(7)E3",
			MacAddress:      "00:1C:F6:C0:33:15",
			SubnetMask:      "255.255.255.0",
			Gateway:         "192.168.3.1",
			Vlan:            "30",
			Uptime:          "N/A",
			Performance: &types.PerformanceMetrics{
				CpuUsagePercent:    0,
				MemoryUsagePercent: 0,
				TemperatureCelsius: 0,
			},
		},
	}

	return devices
}

// GetMockIncidents returns mock incident data
func GetMockIncidents() []*types.Incident {
	return []*types.Incident{
		{
			Id:          "INC-001",
			Title:       "Core Switch SW-001 Down",
			Description: "Connection lost to core switch in datacenter A",
			Priority:    types.IncidentPriority_INCIDENT_PRIORITY_CRITICAL,
			Status:      types.IncidentStatus_INCIDENT_STATUS_IN_PROGRESS,
			AssignedTo:  "Network Team",
		},
		{
			Id:          "INC-002",
			Title:       "High CPU Usage on Router R-042",
			Description: "CPU utilization exceeds 85% threshold",
			Priority:    types.IncidentPriority_INCIDENT_PRIORITY_WARNING,
			Status:      types.IncidentStatus_INCIDENT_STATUS_IN_PROGRESS,
			AssignedTo:  "Performance Team",
		},
		{
			Id:          "INC-003",
			Title:       "Intermittent Connectivity Issues",
			Description: "Users reporting slow internet access",
			Priority:    types.IncidentPriority_INCIDENT_PRIORITY_INFO,
			Status:      types.IncidentStatus_INCIDENT_STATUS_RESOLVED,
			AssignedTo:  "Support Team",
		},
	}
}

// GetMockNetworkStats returns mock network statistics
func GetMockNetworkStats() *types.NetworkStats {
	return &types.NetworkStats{
		ActiveDevices:        247,
		NetworkUptimePercent: 99.98,
		ActiveAlerts:         8,
		AverageLatencyMs:     12,
		TotalBandwidthGbps:   8.5,
	}
}

// GetMockTrafficData returns mock traffic data points
func GetMockTrafficData() []*types.TrafficDataPoint {
	var points []*types.TrafficDataPoint

	// Generate 24 hours of traffic data
	for i := 0; i < 24; i++ {
		points = append(points, &types.TrafficDataPoint{
			Timestamp:    int64(1700000000 + i*3600), // Unix timestamp with hourly intervals
			InboundMbps:  float64(500 + (i*50)%1500),
			OutboundMbps: float64(300 + (i*40)%1200),
		})
	}

	return points
}

// GetMockBandwidthByProtocol returns bandwidth usage by protocol
func GetMockBandwidthByProtocol() []*types.BandwidthByProtocol {
	return []*types.BandwidthByProtocol{
		{Protocol: "HTTP/HTTPS", Percentage: 35.0, BytesPerSecond: 1487500.0},
		{Protocol: "FTP", Percentage: 10.0, BytesPerSecond: 425000.0},
		{Protocol: "Email", Percentage: 15.0, BytesPerSecond: 637500.0},
		{Protocol: "VoIP", Percentage: 12.0, BytesPerSecond: 510000.0},
		{Protocol: "Video Streaming", Percentage: 20.0, BytesPerSecond: 850000.0},
		{Protocol: "Other", Percentage: 8.0, BytesPerSecond: 340000.0},
	}
}

// GetMockDeviceOverview returns device status overview
func GetMockDeviceOverview() []*types.DeviceStatusOverview {
	return []*types.DeviceStatusOverview{
		{
			DeviceType:   types.DeviceType_DEVICE_TYPE_ROUTER,
			TotalCount:   25,
			OnlineCount:  23,
			OfflineCount: 2,
		},
		{
			DeviceType:   types.DeviceType_DEVICE_TYPE_SWITCH,
			TotalCount:   50,
			OnlineCount:  49,
			OfflineCount: 1,
		},
		{
			DeviceType:   types.DeviceType_DEVICE_TYPE_FIREWALL,
			TotalCount:   4,
			OnlineCount:  4,
			OfflineCount: 0,
		},
		{
			DeviceType:   types.DeviceType_DEVICE_TYPE_ACCESS_POINT,
			TotalCount:   168,
			OnlineCount:  146,
			OfflineCount: 22,
		},
	}
}

// GetMockServiceHealth returns service health status
func GetMockServiceHealth() []*types.ServiceHealth {
	return []*types.ServiceHealth{
		{
			ServiceName: "DNS Service",
			Status:      types.ServiceHealthStatus_SERVICE_HEALTH_HEALTHY,
			Description: "DNS resolution service running normally",
		},
		{
			ServiceName: "DHCP Service",
			Status:      types.ServiceHealthStatus_SERVICE_HEALTH_HEALTHY,
			Description: "DHCP lease management operating correctly",
		},
		{
			ServiceName: "Web Portal",
			Status:      types.ServiceHealthStatus_SERVICE_HEALTH_DEGRADED,
			Description: "Web portal experiencing slow response times",
		},
		{
			ServiceName: "Monitoring System",
			Status:      types.ServiceHealthStatus_SERVICE_HEALTH_HEALTHY,
			Description: "Network monitoring system operational",
		},
	}
}

// GetMockDashboardData returns complete dashboard data
func GetMockDashboardData() *types.DashboardData {
	return &types.DashboardData{
		NetworkStats:        GetMockNetworkStats(),
		TrafficData:         GetMockTrafficData(),
		BandwidthByProtocol: GetMockBandwidthByProtocol(),
		DeviceOverview:      GetMockDeviceOverview(),
		RecentIncidents:     GetMockIncidents(),
	}
}

// GetMockLogEntries returns mock log entries for a device
func GetMockLogEntries(deviceId string) []*types.LogEntry {
	return []*types.LogEntry{
		{
			Timestamp: 1700000300, // 5 minutes ago
			Level:     types.LogLevel_LOG_LEVEL_ERROR,
			Message:   "Interface GigabitEthernet0/0/1 changed state to down",
			Source:    deviceId,
		},
		{
			Timestamp: 1700000720, // 12 minutes ago
			Level:     types.LogLevel_LOG_LEVEL_WARNING,
			Message:   "CPU utilization high: 87%",
			Source:    deviceId,
		},
		{
			Timestamp: 1700001800, // 30 minutes ago
			Level:     types.LogLevel_LOG_LEVEL_INFO,
			Message:   "OSPF neighbor 192.168.1.11 Up",
			Source:    deviceId,
		},
		{
			Timestamp: 1700003600, // 1 hour ago
			Level:     types.LogLevel_LOG_LEVEL_INFO,
			Message:   "Configuration changed by admin",
			Source:    deviceId,
		},
	}
}

// GetMockSwitchConfiguration returns mock switch configuration
func GetMockSwitchConfiguration() string {
	return `version 17.6
!
hostname CoreSwitch-A
!
enable secret 5 $1$xyz$abcdefghijklmnop
!
interface GigabitEthernet0/0/1
 description Uplink to Router
 switchport mode trunk
 switchport trunk allowed vlan 10,20,30
!
interface GigabitEthernet0/0/2
 description Backup Uplink
 switchport mode trunk
 switchport trunk allowed vlan 10,20,30
 shutdown
!
vlan 10
 name Management
!
vlan 20
 name Users
!
vlan 30
 name Servers
!
ip default-gateway 192.168.1.1
!
end`
}

// GetMockRouterConfiguration returns mock router configuration
func GetMockRouterConfiguration() string {
	return `version 16.12
!
hostname EdgeRouter-42
!
enable secret 5 $1$abc$defghijklmnopqrst
!
interface GigabitEthernet0/0/0
 description WAN Interface
 ip address 203.0.113.10 255.255.255.252
 no shutdown
!
interface GigabitEthernet0/0/1
 description LAN Interface
 ip address 192.168.2.1 255.255.255.0
 no shutdown
!
router ospf 1
 network 192.168.2.0 0.0.0.255 area 0
 network 203.0.113.8 0.0.0.3 area 0
!
ip route 0.0.0.0 0.0.0.0 203.0.113.9
!
end`
}

// GetMockFirewallConfiguration returns mock firewall configuration
func GetMockFirewallConfiguration() string {
	return `config system interface
    edit "port1"
        set ip 192.168.0.1 255.255.255.0
        set allowaccess https ssh ping
    next
    edit "port2"
        set ip 203.0.113.5 255.255.255.252
        set allowaccess ping
    next
end

config firewall policy
    edit 1
        set name "LAN-to-WAN"
        set srcintf "port1"
        set dstintf "port2"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
    next
end`
}

// GetMockAccessPointConfiguration returns mock access point configuration
func GetMockAccessPointConfiguration() string {
	return `System Name: WiFi-AP-Floor1
Controller IP: 192.168.10.1
AP IP Address: 192.168.10.101
Subnet Mask: 255.255.255.0
Gateway: 192.168.10.1

Radio 2.4GHz:
  Channel: 6
  Power: 100%
  SSID: CorpWiFi-24
  Security: WPA2-PSK

Radio 5GHz:
  Channel: 149
  Power: 100%
  SSID: CorpWiFi-5G
  Security: WPA2-PSK

802.11r Fast Transition: Enabled
Band Steering: Enabled
Load Balancing: Enabled`
}

// GetMockLatencyData returns mock latency data for different locations
func GetMockLatencyData() []*types.LatencyData {
	return []*types.LatencyData{
		{
			Location:         "DC-A",
			AverageLatencyMs: 5.2,
			MinLatencyMs:     3.1,
			MaxLatencyMs:     8.4,
		},
		{
			Location:         "DC-B",
			AverageLatencyMs: 8.1,
			MinLatencyMs:     5.3,
			MaxLatencyMs:     12.8,
		},
		{
			Location:         "Branch-1",
			AverageLatencyMs: 15.3,
			MinLatencyMs:     11.2,
			MaxLatencyMs:     22.5,
		},
		{
			Location:         "Branch-2",
			AverageLatencyMs: 22.7,
			MinLatencyMs:     18.1,
			MaxLatencyMs:     28.9,
		},
		{
			Location:         "Branch-3",
			AverageLatencyMs: 18.9,
			MinLatencyMs:     14.2,
			MaxLatencyMs:     25.7,
		},
		{
			Location:         "Remote-1",
			AverageLatencyMs: 45.2,
			MinLatencyMs:     38.5,
			MaxLatencyMs:     56.8,
		},
	}
}
