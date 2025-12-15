package service

import (
	"bwanews/config"
	"bwanews/internal/adapter/cloudflare"
	"bwanews/internal/adapter/repository"
	"bwanews/internal/core/domain/entity"
	"context"

	"github.com/gofiber/fiber/v2/log"
)

var err error

type ContentService interface {
	GetContents(ctx context.Context) ([]entity.ContentEntity, error)
	GetContent(ctx context.Context, id int64) (*entity.ContentEntity, error)
	CreateContent(ctx context.Context, req entity.ContentEntity) error
	EditContentByID(ctx context.Context, req entity.ContentEntity) error
	DeleteContent(ctx context.Context, id int64) error
	UploadImageR2(ctx context.Context, req entity.FileUploadEntity) (string, error)
}

type contentService struct {
	contentRepository repository.ContentRepository
	cfg               *config.Config
	r2                cloudflare.CloudflareR2Adapter
}

// CreateContent implements ContentService.
func (c *contentService) CreateContent(ctx context.Context, req entity.ContentEntity) error {
	panic("unimplemented")
}

// DeleteContent implements ContentService.
func (c *contentService) DeleteContent(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// EditContentByID implements ContentService.
func (c *contentService) EditContentByID(ctx context.Context, req entity.ContentEntity) error {
	panic("unimplemented")
}

// GetContent implements ContentService.
func (c *contentService) GetContent(ctx context.Context, id int64) (*entity.ContentEntity, error) {
	panic("unimplemented")
}

// GetContents implements ContentService.
func (c *contentService) GetContents(ctx context.Context) ([]entity.ContentEntity, error) {
	results, err := c.contentRepository.GetContents(ctx)

	if err != nil {
		code = "[SERVICE] GetContents - 1"
		
		log.Errorw(code, err)

		return nil, err
	}

	return results, nil
}

// UploadImageR2 implements ContentService.
func (c *contentService) UploadImageR2(ctx context.Context, req entity.FileUploadEntity) (string, error) {
	panic("unimplemented")
}

func NewContentService(contentRepository repository.ContentRepository, cfg *config.Config, r2 cloudflare.CloudflareR2Adapter) ContentService {
	return &contentService{
		contentRepository: contentRepository,
		cfg:               cfg,
		r2:                r2,
	}
}
