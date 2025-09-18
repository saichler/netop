package main

import (
	"os"

	"github.com/saichler/l8inventory/go/inv/service"
	"github.com/saichler/l8pollaris/go/types/l8poll"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/layer8/go/overlay/vnic"
	"github.com/saichler/netop/go/common"
	netosTypes "github.com/saichler/netop/go/types"
	"github.com/saichler/reflect/go/reflect/introspecting"
)

func main() {
	res := common.CreateResources("device-inv-" + os.Getenv("HOSTNAME"))
	res.Logger().SetLogLevel(ifs.Info_Level)
	res.Logger().Info("Starting Device Inventory")
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()

	res.Logger().Info("Registering box service")
	//Add the inventory model and mark the Id field as key
	inventoryNode, _ := nic.Resources().Introspector().Inspect(&netosTypes.Device{})
	introspecting.AddPrimaryKeyDecorator(inventoryNode, "Id")
	nic.Resources().Registry().Register(&netosTypes.DeviceList{})

	//Activate the box inventory service with the primary key & sample model instance
	res.Services().RegisterServiceHandlerType(&inventory.InventoryService{})
	_, err := nic.Resources().Services().Activate(inventory.ServiceType, common.INVENTORY_SERVICE_BOX, common.INVENTORY_AREA_BOX,
		nic.Resources(), nic, "Id", &netosTypes.Device{}, &l8poll.L8ServiceInfo{ServiceName: common.ORM_SERVICE, ServiceArea: 0})

	invCenter := inventory.Inventory(res, common.INVENTORY_SERVICE_BOX, common.INVENTORY_AREA_BOX)
	invCenter.AddStats("Total", Total)
	//invCenter.AddStats("Online", Online)

	if err != nil {
		res.Logger().Error(err)
	}

	res.Logger().SetLogLevel(ifs.Error_Level)
	common.WaitForSignal(nic.Resources())
}

func Total(any interface{}) bool {
	if any == nil {
		return false
	}
	return true
}

/*
func Online(any interface{}) bool {
	if any == nil {
		return false
	}
	nd := any.(*types2.Device)
	if nd.Equipmentinfo == nil {
		return false
	}
	if nd.Equipmentinfo.DeviceStatus == types2.DeviceStatus_DEVICE_STATUS_ONLINE {
		return true
	}
	return false
}*/
