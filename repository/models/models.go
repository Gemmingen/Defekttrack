package models

import (
	"time"

	"defekttrack/api"
)

// GORM-Modelle
type LaptopModel struct {
	ID     int        `gorm:"primaryKey;autoIncrement"`
	Marke  string     `gorm:"not null"`
	Name   string     `gorm:"not null"`
	Os     string     `gorm:"not null"`
	Fehler string     `gorm:"not null"`
	Logs   []LogModel `gorm:"foreignKey:LaptopID;constraint:OnDelete:CASCADE"`
}

func (LaptopModel) TableName() string { return "laptops" }

type LogModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	LaptopID   int       `gorm:"not null;index"`
	Bearbeiter string    `gorm:"not null"`
	Notiz      string    `gorm:"not null"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
}

func (LogModel) TableName() string { return "logs" }

// DB -> API
func (m LaptopModel) ToAPI() api.Laptop {
	var logsPtr *[]api.LogEintrag
	if len(m.Logs) > 0 {
		apiLogs := make([]api.LogEintrag, len(m.Logs))
		for i, l := range m.Logs {
			apiLogs[i] = l.ToAPI()
		}
		logsPtr = &apiLogs
	}

	return api.Laptop{
		Id:     m.ID,
		Marke:  m.Marke,
		Name:   m.Name,
		Os:     m.Os,
		Fehler: api.FehlerKategorie(m.Fehler),
		ItLogs: logsPtr,
	}
}

func (m LogModel) ToAPI() api.LogEintrag {
	return api.LogEintrag{
		Id:         m.ID,
		Bearbeiter: m.Bearbeiter,
		Notiz:      m.Notiz,
		Timestamp:  m.Timestamp,
	}
}

// API -> DB
func ToLaptopModel(input api.LaptopInput) LaptopModel {
	return LaptopModel{
		Marke:  input.Marke,
		Name:   input.Name,
		Os:     input.Os,
		Fehler: string(input.Fehler),
	}
}

func ToLogModel(input api.LogEintragInput, laptopID int) LogModel {
	return LogModel{
		LaptopID:   laptopID,
		Bearbeiter: input.Bearbeiter,
		Notiz:      input.Notiz,
	}
}
