package main

import (
	"os"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/layer8/go/overlay/vnet"
	"github.com/saichler/netop/go/common"
)

func main() {
	resources := common.CreateResources("vnet-" + os.Getenv("HOSTNAME"))
	resources.Logger().SetLogLevel(ifs.Info_Level)
	net := vnet.NewVNet(resources)
	net.Start()
	resources.Logger().Info("vnet started!")
	resources.Logger().SetLogLevel(ifs.Error_Level)
	common.WaitForSignal(resources)
}
