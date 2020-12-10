package model

type SectionResult struct {
	Section      int64         `json:"section"`
	Points       int64         `json:"points"`
	Tops         int64         `json:"tops"`
	Zones        int64         `json:"zones"`
	TopTries     int64         `json:"topTries"`
	ZoneTries    int64         `json:"zoneTries"`
	RouteResults []RouteResult `json:"routeResults"`
}

func (sr SectionResult) HasValidLeadResult() bool {
	return sr.Points > 0
}

func (sr SectionResult) HasValidBoulderResult() bool {
	return !(sr.Tops == 0 && sr.TopTries == 0 && sr.Zones == 0 && sr.ZoneTries == 0)
}
