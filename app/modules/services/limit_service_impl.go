package services

import (
	"fmt"
	"strconv"

	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
)

type LimitServiceImpl struct {
	limitRepository repositories.LimitRepository
}

func NewLimitService(limitRepository repositories.LimitRepository) LimitService {
	return &LimitServiceImpl{limitRepository}
}

func (l *LimitServiceImpl) FindAll() ([]models.Limit, error) {
	return l.limitRepository.FindAll()
}

func (l *LimitServiceImpl) FindById(idStr string) (models.Limit, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Errorf("invalid id")		
	}
	return l.limitRepository.FindById(id)
}

func (l *LimitServiceImpl) Create(req dto.CreateLimitRequest) (*models.Limit, error) {
	amount, err := strconv.Atoi(req.LimitAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid limit amount")
	}
	exist, _ := l.limitRepository.FindByKonsumenTenor(req.KonsumenID, req.Tenor)
	if exist.LimitAmount != 0 {
		return nil, fmt.Errorf("limit already exist")
	}
	limit := &models.Limit{
		KonsumenID:  req.KonsumenID,
		Tenor:       req.Tenor,
		LimitAmount: int64(amount),
	}
	return limit, nil
}

func (l *LimitServiceImpl) Update(idStr string, req dto.UpdateLimitRequest) (*models.Limit, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID")
	}
	amount, err := strconv.Atoi(req.LimitAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid limit amount")
	}
	limit, err := l.limitRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	limit.LimitAmount = int64(amount)
	return &limit, nil
	
}
