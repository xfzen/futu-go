package models

import (
	"futuq/models/akmodels"
)

type (
	TMonthlyIndexData struct {
		akmodels.TIndexMonthly
	}

	TYearlyData struct {
		akmodels.TModelYearlyInfo
	}

	TBoardInfoData struct {
		akmodels.TModelBoardInfo
	}
)
