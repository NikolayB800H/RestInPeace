package schemes

import (
	"awesomeProject/internal/app/ds"
)

type AllDataTypesResponse struct {
	DataTypes []ds.DataTypes `json:"data_types"`
}

type ForecastApplicationsShort struct {
	ApplicationId string `json:"application_id"`
	DataTypeCount int    `json:"data_type_count"`
}

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
	ApplicationCreationDate   string  `json:"application_creation_date"`
	ApplicationFormationDate  *string `json:"application_formation_date"`
	ApplicationCompletionDate *string `json:"application_completion_date"`
	InputStartDate            string  `json:"input_start_date"`
	Creator                   string  `json:"creator"`
	Moderator                 *string `json:"moderator"`
}

func ConvertForecastApplications(application *ds.ForecastApplications) ForecastApplicationsOutput {
	output := ForecastApplicationsOutput{
		ApplicationId:           application.ApplicationId,
		ApplicationStatus:       application.ApplicationStatus,
		ApplicationCreationDate: application.ApplicationCreationDate.Format("2006-01-02 15:04:05"),
		InputStartDate:          application.InputStartDate.Format("2006-01-02 15:04:05"),
		Creator:                 application.Creator.Login,
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

	return output
}
