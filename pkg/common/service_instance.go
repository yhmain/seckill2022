package common

// common包在pkg包下，作为项目中的基础组件
// common是共同依赖，定义了一些基础结构

// 定义了 服务实例 结构体
type ServiceInstance struct {
	Host      string //  Host 主机ip
	Port      int    //  Port HTTP网络服务的端口号
	Weight    int    // 负载权重
	CurWeight int    // 当前权重

	GrpcPort int // RPC服务的端口号
}
