package model

type ListReleasesResponse struct {
	Releases []Application `json:"releases"`
}

type Application struct {
	Name          string      `json:"name"`
	OperatorSetId string      `json:"operatorSetId"`
	Description   string      `json:"description"`
	Components    []Component `json:"components"`
}

type Component struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	LatestArtifactId string `json:"latestArtifactId"`
	ReleaseTimestamp string `json:"releaseTimestamp"`
}
