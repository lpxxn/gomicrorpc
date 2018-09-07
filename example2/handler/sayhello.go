package handler

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example2/proto/model"
	"github.com/lpxxn/gomicrorpc/example2/proto/rpcapi"
	"time"
)

type Say struct {}

var _ rpcapi.SayHandler = (*Say)(nil)

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND

	return nil
}

func (s *Say) Stream(ctx context.Context, req *model.SRequest, stream rpcapi.Say_StreamStream) error {

	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&model.SResponse{Count: int64(i)}); err != nil {
			return err
		}
		// 模拟
		time.Sleep(time.Microsecond * 20)
	}
	return nil

	return nil
}