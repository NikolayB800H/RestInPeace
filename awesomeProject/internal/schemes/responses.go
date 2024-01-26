package schemes

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/role"
	"time"
)

type AllDataTypesResponse struct {
	DataTypes []ds.DataTypes `json:"data_types"`
}

type ForecastApplicationsShort struct {
	ApplicationId string `json:"application_id"`
	DataTypeCount int    `json:"data_type_count"`
}

// GetAllDataTypesResponse model info
// @Description Ответ с черновикомм заявки на прогноз и со всеми типами данных
type GetAllDataTypesResponse struct {
	DraftForecastApplications *ForecastApplicationsShort `json:"draft_application"`
	DataTypes                 []ds.DataTypes             `json:"data_types"`
}

type AllForecastApplicationssResponse struct {
	ForecastApplications []ForecastApplicationsOutput `json:"applications"`
}

type ForecastApplicationsResponse struct {
	ForecastApplications ForecastApplicationsOutput       `json:"application"`
	DataTypes            []ds.ConnectorAppsTypesDataTypes `json:"data_types"`
} //!!!

type UpdateForecastApplicationsResponse struct {
	ForecastApplications ForecastApplicationsOutput `json:"application"`
}

type ForecastApplicationsOutput struct {
	ApplicationId             string  `json:"application_id"`
	ApplicationStatus         string  `json:"application_status"` // Replace with Enum
	CalculateStatus           *string `json:"calculate_status"`   // Replace with Enum
	ApplicationCreationDate   string  `json:"application_creation_date"`
	ApplicationFormationDate  *string `json:"application_formation_date"`
	ApplicationCompletionDate *string `json:"application_completion_date"`
	InputStartDate            *string `json:"input_start_date"`
	Creator                   string  `json:"creator"`
	Moderator                 *string `json:"moderator"`

	Calculated *int64 `json:"calculated"`
	Total      *int64 `json:"total"`
}

type Stats struct {
	Calculated int64
	Total      int64
}

func ConvertForecastApplications(application *ds.ForecastApplications, stats ...*Stats) ForecastApplicationsOutput {
	output := ForecastApplicationsOutput{
		ApplicationId:           application.ApplicationId,
		ApplicationStatus:       application.ApplicationStatus,
		CalculateStatus:         application.CalculateStatus,
		ApplicationCreationDate: application.ApplicationCreationDate.Format("2006-01-02 15:04:05"),
		Creator:                 application.Creator.Login,
	}

	if application.InputStartDate != nil {
		inputStartDate := application.InputStartDate.Format("2006-01-02")
		output.InputStartDate = &inputStartDate
	}

	if application.ApplicationFormationDate != nil {
		formationDate := application.ApplicationFormationDate.Format("2006-01-02 15:04:05")
		output.ApplicationFormationDate = &formationDate
	}

	if application.ApplicationCompletionDate != nil {
		completionDate := application.ApplicationCompletionDate.Format("2006-01-02 15:04:05")
		output.ApplicationCompletionDate = &completionDate
	}

	if application.Moderator != nil {
		output.Moderator = &application.Moderator.Login
	}

	for _, stat := range stats {
		output.Calculated = &stat.Calculated
		output.Total = &stat.Total
	}

	return output
}

type AuthResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	Role        role.Role     `json:"role"`
	Login       string        `json:"login"`
	//TokenType   string        `json:"token_type"`
}

type SwaggerLoginResp struct {
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Role        int    `json:"role"`
	TokenType   string `json:"token_type"`
}
