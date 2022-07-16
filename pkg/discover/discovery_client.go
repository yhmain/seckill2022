package discover

import (
	"log"
	"sync"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/longjoy/micro-go-book/ch13-seckill/pkg/common"
)

// 定义了一个接口来实现 服务与注册发现的客户端
type DiscoveryClient interface {
	/**
	 * 服务注册接口
	 * @param serviceName 服务名
	 * @param instanceId 服务实例Id
	 * @param instancePort 服务实例端口
	 * @param healthCheckUrl 健康检查地址
	 * @param weight 权重
	 * @param meta 服务实例元数据
	 */
	Register(instanceId, svcHost, healthCheckUrl, svcPort string, svcName string, weight int, meta map[string]string, tags []string, logger *log.Logger) bool

	/**
	 * 服务注销接口
	 * @param instanceId 服务实例Id
	 */
	DeRegister(instanceId string, logger *log.Logger) bool

	/**
	 * 发现服务实例接口
	 * @param serviceName 服务名
	 */
	DiscoverServices(serviceName string, logger *log.Logger) []*common.ServiceInstance
}

// 定义一个结构体，来实现接口
// 并存储一些额外的属性
type DiscoveryClientInstance struct {
	Host string //  Host
	Port int    //  Port
	// 连接 consul 的配置
	config *api.Config
	client consul.Client
	mutex  sync.Mutex
	// 服务实例（本地）缓存字段
	instancesMap sync.Map
}
