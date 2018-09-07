package subscriber

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example2/proto/model"
)

func Handler(ctx context.Context, msg *model.SayParam) error {
	fmt.Printf("Received message: ", msg.Msg)
	return nil
}