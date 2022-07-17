package client

import (
	"context"

	"github.com/longjoy/micro-go-book/ch13-seckill/pb"
	"github.com/longjoy/micro-go-book/ch13-seckill/pkg/discover"
	"github.com/longjoy/micro-go-book/ch13-seckill/pkg/loadbalance"
	"github.com/opentracing/opentracing-go"
)

type OAuthClient interface {
	// 校验用户Token
	CheckToken(ctx context.Context, tracer opentracing.Tracer, request *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error)
}

// 定义了OAuthClientImpl结构体
// 里面的属性都是可以配置的
type OAuthClientImpl struct {
	manager     ClientManager           // 客户端管理器
	serviceName string                  // 服务名称
	loadBalance loadbalance.LoadBalance // 负载均衡策略
	tracer      opentracing.Tracer      // 链路追踪系统
}

// OAuthClientImpl实现了CheckToken方法
// 对于使用该RPC客户端的业务服务就可以直接初始化OAuthClientImpl实例，然后调用CheckToken进行用户Token校验
func (impl *OAuthClientImpl) CheckToken(ctx context.Context, tracer opentracing.Tracer, request *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	response := new(pb.CheckTokenResponse)
	// 方法内部进行了RPC调用
	if err := impl.manager.DecoratorInvoke("/pb.OAuthService/CheckToken", "token_check", tracer, ctx, request, response); err == nil {
		return response, nil
	} else {
		return nil, err
	}
}

// 生成OAuthClient的工厂方法
// 会初始化OAuthClientImpl实例
func NewOAuthClient(serviceName string, lb loadbalance.LoadBalance, tracer opentracing.Tracer) (OAuthClient, error) {
	if serviceName == "" {
		serviceName = "oauth"
	}
	if lb == nil {
		lb = defaultLoadBalance
	}

	return &OAuthClientImpl{
		manager: &DefaultClientManager{
			serviceName:     serviceName,
			loadBalance:     lb,
			discoveryClient: discover.ConsulService,
			logger:          discover.Logger,
		},
		serviceName: serviceName,
		loadBalance: lb,
		tracer:      tracer,
	}, nil

}
