package model

type ListOperatorRequirementsResponse struct {
	OperatorRequirements []OperatorApplication `json:"operatorRequirements"`
}

type OperatorApplication struct {
	ApplicationName string      `json:"applicationName"`
	OperatorSetId   string      `json:"operatorSetId"`
	Description     string      `json:"description"`
	Components      []Component `json:"components"`
}

type Component struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	LatestArtifactId string `json:"latestArtifactId"`
	ReleaseTimestamp string `json:"releaseTimestamp"`
}
