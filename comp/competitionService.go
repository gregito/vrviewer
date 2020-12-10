package comp

import (
	"fmt"
	"github.com/gregito/vrviewer/comp/log"
	"net/http"
	"sync"
	"time"

	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"github.com/gregito/vrviewer/webexec"
)

const (
	basePath                 string = "https://vr2.mhssz.hu/api/1.0.0/competitions/"
	firstValidYear                  = 2018
	valueToDisableYearFilter        = 0
)

var httpClient *http.Client

func init() {
	httpClient = webexec.GetClient()
}

func ListAllCompetitionsSimplified() ([]dto.Competition, time.Duration) {
	var ct model.ClimbingType
	return ListCompetitionsByYearAndKind(valueToDisableYearFilter, ct)
}

func ListAllCompetitionDetail() ([]model.CompetitionDetail, []time.Duration) {
	var execDurs []time.Duration
	var compDets []model.CompetitionDetail
	fmt.Println("")
	comps, dur := ListAllCompetitionsSimplified()
	execDurs = append(execDurs, dur)
	comDetChan := make(chan model.CompetitionDetail)
	compDetFetchDurChan := make(chan time.Duration)
	var wg sync.WaitGroup
	for _, cp := range comps {
		wg.Add(1)
		go func(id int64) {
			fetchCompetitionDetailByCompetitionID(httpClient, id, comDetChan, compDetFetchDurChan)
			wg.Done()
		}(cp.ID)
	}
	go func() {
		wg.Wait()
		close(comDetChan)
		close(compDetFetchDurChan)
	}()
	for detail := range comDetChan {
		compDets = append(compDets, detail)
		execDurs = append(execDurs, <-compDetFetchDurChan)
	}
	return compDets, execDurs
}

func fetchCompetitionDetailByCompetitionID(client *http.Client, id int64, comDetChan chan model.CompetitionDetail, compDetFetchDurChan chan time.Duration) {
	cpd, _, d := GetCompetitionResultsByCompetitionId(client, id)
	comDetChan <- cpd
	compDetFetchDurChan <- d
}

func GetCompetitionResultsByCompetitionId(client *http.Client, id int64) (model.CompetitionDetail, error, time.Duration) {
	resp, err, dur := webexec.MeasuredExecuteCallWithClient(client, fmt.Sprintf("%s%d/results", basePath, id), model.CompetitionDetail{})
	if err != nil {
		log.Println(err)
		return model.CompetitionDetail{}, err, dur
	}
	result := convertInterfaceToCompetitionDetail(resp)
	return result, nil, dur
}

func ListCompetitionsByYearAndKind(year int64, kind model.ClimbingType) ([]dto.Competition, time.Duration) {
	resp, err, dur := webexec.MeasuredExecuteCallWithClient(httpClient, basePath, []model.Competition{})
	if err != nil {
		log.Println(err)
		var empty []dto.Competition
		return empty, dur
	}
	iaep := convertInterfaceArrayToCompetitionArray(resp)
	cp := convertCompetitionArrayToCompetitionArray(iaep)
	result := filterCompetitions(cp, year, kind)
	return result, dur
}

func filterCompetitions(comps []dto.Competition, year int64, kind model.ClimbingType) []dto.Competition {
	result := collectCompetitionsByYear(comps, year)
	if kind != model.Boulder && kind != model.Lead {
		return result
	}
	return collectCompetitionsByType(result, kind)
}

func collectCompetitionsByYear(comps []dto.Competition, year int64) []dto.Competition {
	var result []dto.Competition
	if year != valueToDisableYearFilter {
		if year < firstValidYear {
			return result
		}
		for _, c := range comps {
			if c.Year == year {
				result = append(result, c)
			}
		}
		return result
	} else {
		for _, c := range comps {
			if c.Year >= firstValidYear {
				result = append(result, c)
			}
		}
		return result
	}
}

func collectCompetitionsByType(comps []dto.Competition, kind model.ClimbingType) []dto.Competition {
	var result []dto.Competition
	for _, c := range comps {
		if c.Type == kind {
			result = append(result, c)
		}
	}
	return result
}
