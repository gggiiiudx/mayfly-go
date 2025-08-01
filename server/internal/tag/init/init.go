package init

import (
	"context"
	"mayfly-go/initialize"
	"mayfly-go/internal/pkg/event"
	"mayfly-go/internal/tag/api"
	"mayfly-go/internal/tag/application"
	"mayfly-go/internal/tag/infra/persistence"
	"mayfly-go/pkg/eventbus"
	"mayfly-go/pkg/global"
)

func init() {
	initialize.AddInitIocFunc(func() {
		persistence.InitIoc()
		application.InitIoc()
		api.InitIoc()
	})

	initialize.AddInitFunc(Init)
}

func Init() {
	global.EventBus.SubscribeAsync(event.EventTopicResourceOp, "ResourceOpLogApp", func(ctx context.Context, event *eventbus.Event[any]) error {
		codePath := event.Val.(string)
		return application.GetResourceOpLogApp().AddResourceOpLog(ctx, codePath)
	}, false)
}
