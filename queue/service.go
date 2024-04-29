package queue

import (
	"context"
	"github.com/violetpay-org/queuemanager/config"
	"github.com/violetpay-org/queuemanager/item"
	"sync"
)

type Service interface {
	GetQueueName() config.QueueName

	PushToTheQueue(
		item item.Universal,
	) error
}

// ILowLevelQueueOperator allows for some low-level interactions with the queue itself.
// Through this interface, it is possible to run or stop some crucial operations.
type ILowLevelQueueOperator interface {
	StopQueue(
		callback StopCallback,
	) error

	StartQueue(
		onConsume ConsumeCallback,
		waitGroup *sync.WaitGroup,
		ctx *context.Context,
	) error

	// This behavior is repeated IQueueService.PushToTheQueue
	InsertMessage(item item.Universal) error
}