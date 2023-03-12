package models

import (
	"futuq/models/akmodels"
)

type (
	MonthlyIndexData struct {
		akmodels.TIndexMonthly
	}

	TYearlyData struct {
		akmodels.TModelYearlyInfo
	}
)
