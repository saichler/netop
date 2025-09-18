package main

import (
	"os"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/layer8/go/overlay/vnic"
	"github.com/saichler/netop/go/common"
	"github.com/saichler/netop/go/device-inv/service"
)

func createPod() {
	res := common.CreateResources("device-inv-" + os.Getenv("HOSTNAME"))
	res.Logger().SetLogLevel(ifs.Info_Level)
	res.Logger().Info("Starting Device Inventory")
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()
}

func main() {
	res := common.CreateResources("device-inv-" + os.Getenv("HOSTNAME"))
	res.Logger().SetLogLevel(ifs.Info_Level)
	res.Logger().Info("Starting Device Inventory")
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()

	res.Logger().Info("Registering device service")

	res.Services().RegisterServiceHandlerType(&service.DeviceInventoryService{})
	nic.Resources().Services().Activate(service.ServiceType, service.ServiceName, service.ServiceArea, nic.Resources(), nic)
	common.WaitForSignal(nic.Resources())
}
