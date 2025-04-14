package entity

type HealthCheckTable struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Status string `gorm:"unique"`
}

func (HealthCheckTable) TableName() string {
	return "your_schema_name.healthcheck"
}
