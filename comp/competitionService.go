package comp

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gregito/vrviewer/comp/log"

	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"github.com/gregito/vrviewer/webexec"
)

const (
	ValueToDisableYearFilter = 0
	basePath                 = "https://vr.mhssz.hu/api/1.0.0/competitions/"
	firstValidYear           = 2018
)

var httpClient *http.Client

func init() {
	httpClient = webexec.GetClient()
}

func ListAllCompetitionDetail(name string, year int64, kind model.ClimbingType) ([]model.CompetitionDetail, []time.Duration) {
	var execDurs []time.Duration
	var compDets []model.CompetitionDetail
	fmt.Println("")
	comps, dur := ListCompetitionsByYearAndNameAndKind(name, year, kind)
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

func ListCompetitionsByYearAndNameAndKind(name string, year int64, kind model.ClimbingType) ([]dto.Competition, time.Duration) {
	resp, err, dur := webexec.MeasuredExecuteCallWithClient(httpClient, basePath, []model.Competition{})
	if err != nil {
		log.Println(err)
		var empty []dto.Competition
		return empty, dur
	}
	iaep := convertInterfaceArrayToCompetitionArray(resp)
	cp := convertCompetitionArrayToCompetitionArray(iaep)
	result := filterCompetitions(cp, name, year, kind)
	return result, dur
}

func collectCompetitionByName(comps []dto.Competition, name string) []dto.Competition {
	var result []dto.Competition
	for _, comp := range comps {
		if comp.Name == name {
			result = append(result, comp)
		}
	}
	return result
}

func filterCompetitions(comps []dto.Competition, name string, year int64, kind model.ClimbingType) []dto.Competition {
	result := collectCompetitionsByYear(comps, year)
	if len(name) > 0 {
		result = collectCompetitionByName(comps, name)
	}
	if kind != model.Boulder && kind != model.Lead {
		return result
	}
	return collectCompetitionsByType(result, kind)
}

func collectCompetitionsByYear(comps []dto.Competition, year int64) []dto.Competition {
	var result []dto.Competition
	if year != ValueToDisableYearFilter {
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
