package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	serror "github.com/panda8z/shorturl/errors"
	"github.com/panda8z/shorturl/graph/generated"
	"github.com/panda8z/shorturl/graph/model"
	m "github.com/panda8z/shorturl/model"
	"github.com/panda8z/shorturl/x"
)

func (r *mutationResolver) CreateURL(ctx context.Context, input model.NewURL) (*model.URL, error) {
	long, errCode := m.CheckUrlOrigin(input.Origin)
	if errCode == serror.SUCCESS {
		url := x.CreateUrl(input.Origin)
		if url != nil {
			return &model.URL{
				ID:     strconv.FormatInt(int64(url.ID), 10),
				Origin: url.Origin,
				Short:  url.Short,
			}, nil
		}
	}

	if errCode == serror.Exist && long != nil {
		return &model.URL{
			ID:     strconv.FormatInt(int64(long.ID), 10),
			Origin: long.Origin,
			Short:  long.Short,
		}, nil
	}
	return nil, nil
}

func (r *queryResolver) Origin(ctx context.Context, origin string) (*model.URL, error) {
	long, errCode := m.CheckUrlOrigin(origin)
	if errCode == serror.SUCCESS {
		url := x.CreateUrl(long.Origin)
		return &model.URL{
			ID:     strconv.FormatInt(int64(url.ID), 10),
			Origin: url.Origin,
			Short:  url.Short,
		}, nil
	}

	if errCode == serror.Exist && long != nil {
		return &model.URL{
			ID:     strconv.FormatInt(int64(long.ID), 10),
			Origin: long.Origin,
			Short:  long.Short,
		}, nil
	}
	return nil, nil
}

func (r *queryResolver) Short(ctx context.Context, short string) (*model.URL, error) {
	url, errCode := m.CheckUrlShort(short)
	if errCode == serror.Exist {
		return &model.URL{
			ID:     strconv.FormatInt(int64(url.ID), 10),
			Origin: url.Origin,
			Short:  url.Short,
		}, nil
	} else if errCode == serror.UnExist {
		return nil, &serror.SurlErr{Msg: serror.CodeMap[errCode]}
	}
	return nil, nil
}

func (r *queryResolver) Check(ctx context.Context, id string) (*model.URL, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
