package service

import (
	"TheFirstKrotas/api/helloworld/v1"
	"context"
	"fmt"
)

type DemoService struct {
	v1.UnimplementedDemoServer
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *v1.CreateDemoRequest) (*v1.CreateDemoReply, error) {

	return &v1.CreateDemoReply{Status: 1}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *v1.UpdateDemoRequest) (*v1.UpdateDemoReply, error) {
	return &v1.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *v1.DeleteDemoRequest) (*v1.DeleteDemoReply, error) {
	return &v1.DeleteDemoReply{Status: 0}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *v1.GetDemoRequest) (*v1.GetDemoReply, error) {
	fmt.Println("Hello 123")
	return &v1.GetDemoReply{Res: "Hello World !!!"}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *v1.ListDemoRequest) (*v1.ListDemoReply, error) {
	return &v1.ListDemoReply{}, nil
}

type Message struct {
	Head    string
	Context string
	Size    int64
}
