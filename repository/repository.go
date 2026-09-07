package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"defekttrack/api"
	"defekttrack/repository/models" // Unter-Package importieren
)

type Repository interface {
	FindAll(fehlerFilter string) ([]api.Laptop, error)
	FindByID(id int) (api.Laptop, error)
	Create(input api.LaptopInput) (api.Laptop, error)
	Update(id int, input api.LaptopInput) (api.Laptop, error)
	Delete(id int) error
	CreateLog(laptopID int, input api.LogEintragInput) (api.LogEintrag, error)
	UpdateLog(laptopID int, logID int, input api.LogEintragInput) (api.LogEintrag, error)
	DeleteLog(id int, laptopID int) error
}

type gormRepository struct {
	db *gorm.DB
}

func ConnectAndMigrate(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Import aus dem models-Package
	if err := db.AutoMigrate(&models.LaptopModel{}, &models.LogModel{}); err != nil {
		return nil, err
	}
	return db, nil
}

func New(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) FindAll(fehlerFilter string) ([]api.Laptop, error) {
	var dbModels []models.LaptopModel
	query := r.db.Preload("Logs")
	if fehlerFilter != "" {
		query = query.Where("fehler = ?", fehlerFilter)
	}
	if err := query.Find(&dbModels).Error; err != nil {
		return nil, err
	}

	result := make([]api.Laptop, 0, len(dbModels))
	for _, m := range dbModels {
		result = append(result, m.ToAPI())
	}
	return result, nil
}

func (r *gormRepository) FindByID(id int) (api.Laptop, error) {
	var m models.LaptopModel
	if err := r.db.Preload("Logs").First(&m, id).Error; err != nil {
		return api.Laptop{}, err
	}
	return m.ToAPI(), nil
}

func (r *gormRepository) Create(input api.LaptopInput) (api.Laptop, error) {
	dbModel := models.ToLaptopModel(input)
	if err := r.db.Create(&dbModel).Error; err != nil {
		return api.Laptop{}, err
	}
	return dbModel.ToAPI(), nil
}

func (r *gormRepository) Update(id int, input api.LaptopInput) (api.Laptop, error) {
	var m models.LaptopModel
	if err := r.db.First(&m, id).Error; err != nil {
		return api.Laptop{}, err
	}
	m.Marke = input.Marke
	m.Name = input.Name
	m.Os = input.Os
	m.Fehler = string(input.Fehler)

	if err := r.db.Save(&m).Error; err != nil {
		return api.Laptop{}, err
	}
	return m.ToAPI(), nil
}

func (r *gormRepository) Delete(id int) error {
	res := r.db.Delete(&models.LaptopModel{}, id)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (r *gormRepository) CreateLog(laptopID int, input api.LogEintragInput) (api.LogEintrag, error) {
	logModel := models.ToLogModel(input, laptopID)
	if err := r.db.Create(&logModel).Error; err != nil {
		return api.LogEintrag{}, err
	}
	return logModel.ToAPI(), nil
}

func (r *gormRepository) UpdateLog(laptopID int, logID int, input api.LogEintragInput) (api.LogEintrag, error) {
	var logModel models.LogModel
	if err := r.db.Where("id = ? AND laptop_id = ?", logID, laptopID).First(&logModel).Error; err != nil {
		return api.LogEintrag{}, err
	}
	logModel.Bearbeiter = input.Bearbeiter
	logModel.Notiz = input.Notiz

	if err := r.db.Save(&logModel).Error; err != nil {
		return api.LogEintrag{}, err
	}
	return logModel.ToAPI(), nil
}

func (r *gormRepository) DeleteLog(id int, laptopID int) error {
	res := r.db.Where("id = ? AND laptop_id = ?", id, laptopID).Delete(&models.LogModel{})
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
