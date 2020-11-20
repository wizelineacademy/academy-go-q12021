package repository

import "time"

type AdminRepository interface {
	FindBy(searchPattern string, startDate time.Time, endDate time.Time) ([]string, error)
}
