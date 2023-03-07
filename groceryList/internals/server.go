package internals

import (
	"context"
	"fmt"
	"grocerylist/models"
	ls "grocerylist/protos/listserver"
	itemRepo "grocerylist/repository"
	"log"
)

type Server struct {
	ls.UnimplementedListServiceServer
}

func (s *Server) InsertListItem(ctx context.Context, in *ls.Item) (*ls.InsertResponse, error) {
	if in.Item == "" {
		return &ls.InsertResponse{Error: "Empty Item", Status: "error"}, fmt.Errorf("empty item passed")
	}
	repo := itemRepo.NewItemRepository()
	status, err := repo.CreateItem(in.Item)
	if status {
		return &ls.InsertResponse{Error: "", Status: "inserted"}, nil
	} else {
		return &ls.InsertResponse{Error: err.Error(), Status: "error"}, fmt.Errorf("[db]:%v", err.Error())
	}

}

func (s *Server) FindListItem(ctx context.Context, in *ls.ItemReq) (*ls.Item, error) {
	if in.ItemName == "" {
		return &ls.Item{}, fmt.Errorf("not found")
	}

	repo := itemRepo.NewItemRepository()
	response := repo.GetItem(in.ItemName)
	return &ls.Item{Item: response.Name}, fmt.Errorf("not found")
}

func (s *Server) DeleteListItem(ctx context.Context, in *ls.ItemReq) (*ls.DeleteResponse, error) {
	if in.ItemName == "" {
		return &ls.DeleteResponse{Error: "Item Request is empty"}, nil
	}

	repo := itemRepo.NewItemRepository()
	status, err := repo.DeleteItem(in.ItemName)
	if status {
		return &ls.DeleteResponse{Error: "", Status: "Deleted"}, nil
	} else {
		return &ls.DeleteResponse{Error: err.Error(), Status: "error"}, nil
	}
}

func (s *Server) GetList(in *ls.Empty, srv ls.ListService_GetListServer) error {
	repo := itemRepo.NewItemRepository()
	var list []models.Item
	repo.GetList(&list)
	log.Println(list)
	for _, v := range list {
		resp := ls.Item{Item: v.Name}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
	return nil
}

func (s *Server) UpdateListItem(ctx context.Context, in *ls.UpdateReq) (*ls.Item, error) {
	if in.ItemName == "" {
		return &ls.Item{}, fmt.Errorf("request is empty")
	}
	var search ls.Item
	log.Println("hello from upli")

	if search.Item != "" {
		search.Item = in.NewName
		newItem := ls.Item{Item: in.NewName}
		return &newItem, nil
	}

	return &ls.Item{}, fmt.Errorf("not found")
}
