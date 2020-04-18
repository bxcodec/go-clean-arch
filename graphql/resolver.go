package graphql

import (
	"context"
	"github.com/bxcodec/go-clean-arch/article"

	"github.com/bxcodec/go-clean-arch/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	ArticleUsecase article.Usecase
}

func (r *Resolver) Article() ArticleResolver {
	return &articleResolver{r}
}
func (r *Resolver) Author() AuthorResolver {
	return &authorResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type articleResolver struct{ *Resolver }

func (r *articleResolver) ID(ctx context.Context, obj *models.Article) (string, error) {
	panic("not implemented")
}

type authorResolver struct{ *Resolver }

func (r *authorResolver) ID(ctx context.Context, obj *models.Author) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateArticle(ctx context.Context, input NewArticle) (*models.Article, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Articles(ctx context.Context) ([]*models.Article, error) {
	articles, _, err := r.ArticleUsecase.Fetch(ctx, "cursor", 1) // <--- example values because I don't want to change Usecase now
	return articles, err
}
func (r *queryResolver) Authors(ctx context.Context) ([]*models.Author, error) {
	panic("not implemented")
}
