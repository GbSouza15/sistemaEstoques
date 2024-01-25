package service

import (
	"encoding/json"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/google/uuid"
)

type DepositService struct {
	repo *repository.DepositRepository
}

func NewDepositService(repo *repository.DepositRepository) *DepositService {
	return &DepositService{repo}
}

func (ds *DepositService) CreateDeposit(body []byte) error {
	var deposit models.Deposit

	depositId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &deposit); err != nil {
		return err
	}

	if err := ds.repo.CreateDeposit(depositId, deposit.Name, deposit.Address, deposit.CompanyId); err != nil {
		return err
	}

	return nil
}

func (ds *DepositService) ListDeposits(companyId uuid.UUID) ([]models.Deposit, error) {
	deposits, err := ds.repo.ListDeposits(companyId)
	if err != nil {
		return nil, err
	}

	return deposits, nil
}

func (ds *DepositService) DeleteDeposit(id string) error {
	if err := ds.repo.DeleteDeposit(id); err != nil {
		return err
	}

	return nil
}

func (ds *DepositService) UpdateDeposit(body []byte) error {
	var depositUpdate models.DepositUpdate

	if err := json.Unmarshal(body, &depositUpdate); err != nil {
		return err
	}

	if err := ds.repo.UpdateDeposit(depositUpdate.Id, depositUpdate.Name, depositUpdate.Address); err != nil {
		return err
	}

	return nil
}
