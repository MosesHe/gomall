package rpc

import (
	"sync"

	"github.com/MosesHe/gomall/app/cart/conf"
	cartUtils "github.com/MosesHe/gomall/app/cart/utils"
	"github.com/MosesHe/gomall/common/clientsuite"
	"github.com/MosesHe/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(&clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("productcatalogservice", opts...)
	cartUtils.MustHandleError(err)
}
