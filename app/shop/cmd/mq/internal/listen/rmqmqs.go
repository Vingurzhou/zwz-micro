package listen

import (
	"context"
	"github.com/zeromicro/go-zero/core/service"
	"looklook/app/shop/cmd/mq/internal/config"
	"looklook/app/shop/cmd/mq/internal/svc"
	"looklook/common/rabbitmq"
)

func RabbitMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		rabbitmq.NewRabbitMQSimple(),
	}

}
