package service

import (
	"github.com/saichler/l8services/go/services/dcache"
	"github.com/saichler/l8srlz/go/serialize/object"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types/l8api"
	"github.com/saichler/l8types/go/types/l8web"
	"github.com/saichler/l8utils/go/utils/web"
	"github.com/saichler/netop/go/types"
	"github.com/saichler/reflect/go/reflect/introspecting"
)

const (
	ServiceType = "DeviceInventoryService"
	ServiceName = "dev-inv"
	ServiceArea = byte(0)
)

type DeviceInventoryService struct {
	cache ifs.IDistributedCache
}

func (this *DeviceInventoryService) Activate(serviceName string, serviceArea byte, r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.Device{})
	r.Registry().Register(&types.DeviceList{})
	node, _ := r.Introspector().Inspect(&types.Device{})
	introspecting.AddPrimaryKeyDecorator(node, "Id")
	this.cache = dcache.NewDistributedCache(serviceName, serviceArea, &types.Device{}, nil, l, r)
	return nil
}

func (this *DeviceInventoryService) DeActivate() error { return nil }
func (this *DeviceInventoryService) Post(elems ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	for _, element := range elems.Elements() {
		if !elems.Notification() {
			exist, _ := this.cache.Get(element)
			if exist != nil {
				return object.NewError("Element Already Exist")
			}
		}
		_, err := this.cache.Post(element, elems.Notification())
		if err != nil {
			return object.NewError(err.Error())
		}
	}
	return object.New(nil, &l8web.L8Empty{})
}
func (this *DeviceInventoryService) Put(elems ifs.IElements, nic ifs.IVNic) ifs.IElements {
	for _, element := range elems.Elements() {
		if !elems.Notification() {
			exist, _ := this.cache.Get(element)
			if exist == nil {
				return object.NewError("Element Does not Exist")
			}
		}
		_, err := this.cache.Put(element, elems.Notification())
		if err != nil {
			return object.NewError(err.Error())
		}
	}
	return object.New(nil, &l8web.L8Empty{})
}
func (this *DeviceInventoryService) Patch(elems ifs.IElements, nic ifs.IVNic) ifs.IElements {
	for _, element := range elems.Elements() {
		if !elems.Notification() {
			exist, _ := this.cache.Get(element)
			if exist == nil {
				return object.NewError("Element does not Exist")
			}
		}
		_, err := this.cache.Patch(element, elems.Notification())
		if err != nil {
			return object.NewError(err.Error())
		}
	}
	return object.New(nil, &l8web.L8Empty{})
}
func (this *DeviceInventoryService) Delete(elems ifs.IElements, nic ifs.IVNic) ifs.IElements {
	for _, element := range elems.Elements() {
		_, err := this.cache.Delete(element, elems.Notification())
		if err != nil {
			return object.NewError(err.Error())
		}
	}
	return object.New(nil, &l8web.L8Empty{})
}
func (this *DeviceInventoryService) Get(elems ifs.IElements, nic ifs.IVNic) ifs.IElements {
	if elems.IsFilterMode() {
		elem, err := this.cache.Get(elems.Element())
		return object.New(err, elem)
	}
	q, e := elems.Query(nic.Resources())
	if e != nil {
		return object.NewError(e.Error())
	}
	return object.New(nil, this.cache.Fetch(int(q.Page()), int(q.Limit())))
}
func (this *DeviceInventoryService) Failed(elems ifs.IElements, vnic ifs.IVNic, msg *ifs.Message) ifs.IElements {
	return nil
}
func (this *DeviceInventoryService) TransactionConfig() ifs.ITransactionConfig {
	return nil
}
func (this *DeviceInventoryService) WebService() ifs.IWebService {
	ws := web.New(ServiceName, ServiceArea, nil,
		nil, nil, nil, nil, nil, nil, nil,
		&l8api.L8Query{}, &types.DeviceList{})
	return ws
}
