package grpc

import (
	"2024_2_ThereWillBeName/internal/pkg/search"
	searchGen "2024_2_ThereWillBeName/internal/pkg/search/delivery/grpc/gen"
	"context"
	"log/slog"
)

type GrpcSearchHandler struct {
	searchGen.SearchServer
	uc     search.SearchUsecase
	logger *slog.Logger
}

func NewGrpcSearchHandler(uc search.SearchUsecase, logger *slog.Logger) *GrpcSearchHandler {
	return &GrpcSearchHandler{uc: uc, logger: logger}
}

func (h *GrpcSearchHandler) Search(ctx context.Context, in *searchGen.SearchRequest) (*searchGen.SearchResponse, error) {
	results, err := h.uc.Search(context.Background(), in.DecodedQuery)
	if err != nil {
		return nil, err
	}

	grpcResults := make([]*searchGen.SearchResult, len(results))
	for i, searchResult := range results {
		grpcResults[i] = &searchGen.SearchResult{
			Name: searchResult.Name,
			Id:   uint32(searchResult.Id),
			Type: searchResult.Type,
		}
	}

	return &searchGen.SearchResponse{SearchResult: grpcResults}, nil
}
