package main

import (
	"os"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8web/go/web/server"
	"github.com/saichler/l8bus/go/overlay/health"
	"github.com/saichler/l8bus/go/overlay/protocol"
	"github.com/saichler/l8bus/go/overlay/vnic"
	"github.com/saichler/netop/go/common"
	"github.com/saichler/netop/go/types"
	"github.com/saichler/l8reflect/go/reflect/introspecting"
)

func main() {
	resources := common.CreateResources("web-service" + os.Getenv("HOSTNAME"))
	resources.Logger().SetLogLevel(ifs.Info_Level)
	startWebServer(8443, "/data/netop")
}

func startWebServer(port int, cert string) {
	serverConfig := &server.RestServerConfig{
		Host:           protocol.MachineIP,
		Port:           port,
		Authentication: false,
		CertName:       cert,
		Prefix:         common.PREFIX,
	}
	svr, err := server.NewRestServer(serverConfig)
	if err != nil {
		panic(err)
	}

	resources := common.CreateResources("web-" + os.Getenv("HOSTNAME"))

	node, _ := resources.Introspector().Inspect(&types.Device{})
	introspecting.AddPrimaryKeyDecorator(node, "Id")

	nic := vnic.NewVirtualNetworkInterface(resources, nil)
	nic.Resources().SysConfig().KeepAliveIntervalSeconds = 60
	nic.Start()
	nic.WaitForConnection()

	nic.Resources().Registry().Register(&types.Device{})
	nic.Resources().Registry().Register(&types.DeviceList{})

	/*
		nic.Resources().Registry().Register(&types4.Pollaris{})
		nic.Resources().Registry().Register(&types4.Device{})
		nic.Resources().Registry().Register(&types4.DeviceList{})
		nic.Resources().Registry().Register(&types.NetworkDevice{})
		nic.Resources().Registry().Register(&types.NetworkDeviceList{})
		nic.Resources().Registry().Register(&types2.K8SCluster{})
		nic.Resources().Registry().Register(&types2.K8SClusterList{})
		nic.Resources().Registry().Register(&types3.Query{})
		nic.Resources().Registry().Register(&types3.Top{})
		nic.Resources().Registry().Register(&types3.Empty{})
		nic.Resources().Registry().Register(&types4.CJob{})
		nic.Resources().Registry().Register(&types2.NetworkTopology{})
	*/

	hs, ok := nic.Resources().Services().ServiceHandler(health.ServiceName, 0)
	if ok {
		ws := hs.WebService()
		svr.RegisterWebService(ws, nic)
	}

	//Activate the webpoints service
	nic.Resources().Services().RegisterServiceHandlerType(&server.WebService{})
	_, err = nic.Resources().Services().Activate(server.ServiceTypeName, ifs.WebService,
		0, nic.Resources(), nic, svr)

	nic.Resources().Logger().Info("Web Server Started!")

	svr.Start()
}
