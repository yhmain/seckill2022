package loadbalance

import (
	"errors"
	"math/rand"

	"github.com/longjoy/micro-go-book/ch13-seckill/pkg/common"
)

// 负载均衡器
// 接口：根据一定的负载均衡策略从服务实例列表中选择一个服务实例返回
type LoadBalance interface {
	SelectService(service []*common.ServiceInstance) (*common.ServiceInstance, error)
}

// 完全随机策略
type RandomLoadBalance struct {
}

// 随机负载均衡
func (loadBalance *RandomLoadBalance) SelectService(services []*common.ServiceInstance) (*common.ServiceInstance, error) {

	if services == nil || len(services) == 0 {
		return nil, errors.New("service instances are not exist")
	}

	return services[rand.Intn(len(services))], nil
}

// 带权重的平滑轮询策略
type WeightRoundRobinLoadBalance struct {
}

// 权重平滑负载均衡
func (loadBalance *WeightRoundRobinLoadBalance) SelectService(services []*common.ServiceInstance) (best *common.ServiceInstance, err error) {

	if services == nil || len(services) == 0 {
		return nil, errors.New("service instances are not exist")
	}

	total := 0
	for i := 0; i < len(services); i++ {
		w := services[i]
		if w == nil {
			continue
		}
		// 对于每个服务实例，增加它的Weight值到CurWeight
		w.CurWeight += w.Weight
		// 增加所有服务实例的Weight
		total += w.Weight
		// 选择具有最大CurWeight值的服务实例
		if best == nil || w.CurWeight > best.CurWeight {
			best = w
		}
	}

	if best == nil {
		return nil, nil
	}

	best.CurWeight -= total
	return best, nil
}
