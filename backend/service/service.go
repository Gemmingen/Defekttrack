package service

import (
	"errors"
	"fmt"
	"strings"

	"defekttrack/api"
	"defekttrack/repository"
)

var ErlaubteMarken = []string{"HP", "Lenovo", "Dell", "Apple", "Asus", "Acer"}

type LaptopService interface {
	GetLaptops(filter string) ([]api.Laptop, error)
	GetLaptopByID(id int) (api.Laptop, error)
	CreateLaptop(input api.LaptopInput) (api.Laptop, error)
	UpdateLaptop(id int, input api.LaptopInput) (api.Laptop, error)
	DeleteLaptop(id int) error
	CreateLog(laptopID int, input api.LogEintragInput) (api.LogEintrag, error)
	UpdateLog(laptopID int, logID int, input api.LogEintragInput) (api.LogEintrag, error)
	DeleteLog(laptopID int, logID int) error
}

type laptopService struct {
	repo repository.Repository
}

func New(repo repository.Repository) LaptopService {
	return &laptopService{repo: repo}
}

func validateBrand(marke string) error {
	for _, m := range ErlaubteMarken {
		if strings.EqualFold(m, marke) {
			return nil
		}
	}
	return fmt.Errorf("ungültige Marke '%s'. Erlaubt: %s", marke, strings.Join(ErlaubteMarken, ", "))
}

func validateNote(notiz string) error {
	if len(strings.TrimSpace(notiz)) < 5 {
		return errors.New("die Notiz muss mindestens 5 Zeichen lang sein")
	}
	return nil
}

func (s *laptopService) GetLaptops(filter string) ([]api.Laptop, error) {
	return s.repo.FindAll(filter)
}

func (s *laptopService) GetLaptopByID(id int) (api.Laptop, error) {
	return s.repo.FindByID(id)
}

func (s *laptopService) CreateLaptop(input api.LaptopInput) (api.Laptop, error) {
	if err := validateBrand(input.Marke); err != nil {
		return api.Laptop{}, err
	}
	return s.repo.Create(input)
}

func (s *laptopService) UpdateLaptop(id int, input api.LaptopInput) (api.Laptop, error) {
	if err := validateBrand(input.Marke); err != nil {
		return api.Laptop{}, err
	}
	return s.repo.Update(id, input)
}

func (s *laptopService) DeleteLaptop(id int) error {
	return s.repo.Delete(id)
}

func (s *laptopService) CreateLog(laptopID int, input api.LogEintragInput) (api.LogEintrag, error) {
	if err := validateNote(input.Notiz); err != nil {
		return api.LogEintrag{}, err
	}
	if _, err := s.repo.FindByID(laptopID); err != nil {
		return api.LogEintrag{}, errors.New("laptop existiert nicht")
	}
	return s.repo.CreateLog(laptopID, input)
}

func (s *laptopService) UpdateLog(laptopID int, logID int, input api.LogEintragInput) (api.LogEintrag, error) {
	if err := validateNote(input.Notiz); err != nil {
		return api.LogEintrag{}, err
	}
	return s.repo.UpdateLog(laptopID, logID, input)
}

func (s *laptopService) DeleteLog(laptopID int, logID int) error {
	return s.repo.DeleteLog(logID, laptopID)
}
