package service

import (
	"context"

	"github.com/techx/portal/builder"
	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
)

type AdminService interface {
	UpdateCompanyDetails(ctx context.Context, params *domain.Company) (*domain.EmptyDomain, error)
	BulkUpdateUsers(ctx context.Context, from, to domain.UserProfile) (*domain.EmptyDomain, error)
}

type adminService struct {
	cfg      *config.Config
	registry *builder.Registry
}

func NewAdminService(cfg *config.Config, registry *builder.Registry) AdminService {
	return &adminService{
		cfg:      cfg,
		registry: registry,
	}
}

func (a adminService) UpdateCompanyDetails(ctx context.Context, params *domain.Company) (*domain.EmptyDomain, error) {
	err := a.registry.CompaniesRepository.UpdateCompany(ctx, params)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}

func (a adminService) BulkUpdateUsers(ctx context.Context, from, to domain.UserProfile) (*domain.EmptyDomain, error) {
	err := a.registry.UsersRepository.BulkUpdate(ctx, from, to)
	if err != nil {
		return nil, err
	}

	return &domain.EmptyDomain{}, nil
}
