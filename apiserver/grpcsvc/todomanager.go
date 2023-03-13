package grpcsvc

import (
	"context"
	"fmt"
	"log"

	"galaxywave.com/go-todo/apiserver/models"
	"galaxywave.com/go-todo/apiserver/services"
	pb "galaxywave.com/go-todo/todoapi"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TodoManager struct {
	pb.UnimplementedTodoManagerServer
}

// SayHello implements helloworld.GreeterServer
func (s *TodoManager) GetTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())
	var todo models.Todo

	if err := models.DB.First(&todo, &in).Error; err != nil {
		fmt.Println("error: ", err)
		return &pb.TodoReply{}, err
	}
	fmt.Println(todo, in)
	return &pb.TodoReply{
		Id:    todo.ID,
		Title: todo.Title,
	}, nil
}

func (s *TodoManager) WatchNewTodo(_ *emptypb.Empty, src pb.TodoManager_WatchNewTodoServer) error {
	for {
		todo := <-services.TODOChan
		fmt.Println(todo)
		src.Send(&pb.TodoReply{Id: todo.ID, Title: todo.Title})
	}
	return nil
}
