package services

import (
	"context"
	"errors"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/models"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/repositories"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrNotFound = errors.New("unit not found")

type UnitService struct {
	repo *repositories.UnitRepository
}

func NewUnitService(repo *repositories.UnitRepository) *UnitService {
	return &UnitService{repo: repo}
}

func (s *UnitService) CreateUnit(ctx context.Context, b *models.UnitBoundary) (*models.Unit, error) {
	if b.Name == "" || b.Description == "" {
		return nil, errors.New("name and description are required")
	}
	u := &models.Unit{
		UnitID:      b.UnitID,
		Name:        b.Name,
		Description: b.Description,
	}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UnitService) GetUnitByID(ctx context.Context, id string) (*models.Unit, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNotFound
	}
	return u, err
}

func (s *UnitService) ListUnits(ctx context.Context, page, size int64) ([]models.Unit, error) {
	return s.repo.List(ctx, page, size)
}

func (s *UnitService) UpdateUnit(ctx context.Context, id string, b *models.UnitBoundary) error {
	upd := make(map[string]interface{})
	if b.Name != "" {
		upd["name"] = b.Name
	}
	if b.Description != "" {
		upd["description"] = b.Description
	}
	return s.repo.Update(ctx, id, upd)
}

func (s *UnitService) DeleteAllUnits(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
