package rpc

import (
	"sync"

	"github.com/MosesHe/gomall/app/cart/conf"
	cartUtils "github.com/MosesHe/gomall/app/cart/utils"
	"github.com/MosesHe/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)

	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartUtils.MustHandleError(err)
}
