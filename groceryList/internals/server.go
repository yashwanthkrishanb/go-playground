package internals

import (
	"context"
	"fmt"
	ls "grocerylist/protos/listserver"
	itemRepo "grocerylist/repository"
	"log"
)

type Server struct {
	ls.UnimplementedListServiceServer
	list []ls.Item
}

func (s *Server) InsertListItem(ctx context.Context, in *ls.Item) (*ls.InsertResponse, error) {
	if in.Item == "" {
		return &ls.InsertResponse{Error: "Empty Item", Status: "error"}, nil
	}
	s.list = append(s.list, *in)
	repo := itemRepo.NewItemRepository()
	status, err := repo.CreateItem(in.Item)
	if status {
		return &ls.InsertResponse{Error: "", Status: "inserted"}, nil
	} else {
		return &ls.InsertResponse{Error: err.Error(), Status: "error"}, nil
	}

}

func (s *Server) FindListItem(ctx context.Context, in *ls.ItemReq) (*ls.Item, error) {
	if in.ItemName == "" {
		return &ls.Item{}, fmt.Errorf("Not Found")
	}
	for _, v := range s.list {
		if v.Item == in.ItemName {
			log.Println(v)
			return &v, nil
		}
	}
	repo := itemRepo.NewItemRepository()
	response := repo.GetItem(in.ItemName)
	return &ls.Item{Item: response.Name}, fmt.Errorf("Not Found")
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

// func (s *Server) GetList(*ls.Empty,) error {
// 	return nil
// }

func (s *Server) UpdateListItem(ctx context.Context, in *ls.UpdateReq) (*ls.Item, error) {
	if in.ItemName == "" {
		return &ls.Item{}, fmt.Errorf("request is empty")
	}
	var search ls.Item
	log.Println("hello from upli")
	var idx int
	for i, v := range s.list {
		if v.Item == in.ItemName {
			log.Println(v)
			search = v
			idx = i
		}
	}
	if search.Item != "" {
		search.Item = in.NewName
		newItem := ls.Item{Item: in.NewName}
		s.list[idx] = newItem
		log.Println(newItem)
		log.Println(s.list)
		return &newItem, nil
	}

	return &ls.Item{}, fmt.Errorf("Not Found")
}
