package loader

import (
	"encoding/json"
	"fmt"
	"github.com/gregito/vrviewer/comp/common"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
	"github.com/gregito/vrviewer/comp/model"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	vrvTmpDirName        = "vrvTmp/"
	manualTempDirPathKey = "VRV_TMPDIR"
)

var appVersion string

func init() {
	appVersion = "0.1.0"
}

func CompetitionsFromFile() ([]dto.Competition, error, bool) {
	var result []dto.Competition
	var cacheFile dto.CompetitionFile
	fileName := "cmps"
	log.Println("Reading CompetitionDetail from: " + getTempFolderPath() + fileName)
	file, err := ioutil.ReadFile(getTempFolderPath() + fileName)
	if err != nil {
		log.Printf("Unable to read file to obtain stored competition(s): %s", err)
		return result, err, true
	}
	err = json.Unmarshal(file, &cacheFile)
	if err != nil {
		log.Println("Unable to unmarshall competition(s): ", err)
		return result, err, true
	}
	return cacheFile.Competitions, nil, isRenewalNeeded(cacheFile.CacheProperty)
}

func CompetitionDetailsFromFile() (map[int64]model.CompetitionDetail, error, bool) {
	var result map[int64]model.CompetitionDetail
	var cacheFile dto.CompetitionDetailFile
	fileName := fmt.Sprintf("%scds", getTempFolderPath())
	log.Println("Reading CompetitionDetail cache from: " + getTempFolderPath())
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Unable to read competition details cache from file by name \"%s\" because: %s\n", fileName, err)
		return result, err, true
	}
	err = json.Unmarshal(file, &cacheFile)
	if err != nil {
		log.Println("Unable to unmarshall competition details cache: ", err)
		return result, err, true
	}
	return unload(cacheFile.CompetitionDetails), nil, isRenewalNeeded(cacheFile.CacheProperty)
}

func WriteCompetitionsIntoFile(comps []dto.Competition) error {
	var err error
	cacheFile := &dto.CompetitionFile{
		CacheProperty: createCacheFile(),
		Competitions:  comps,
	}
	content, err := json.Marshal(cacheFile)
	if err != nil {
		return err
	}
	file := getTempFolderPath() + "cmps"
	log.Println("Writing file: " + file)
	err = ioutil.WriteFile(file, content, 0644)
	return err
}

func WriteCompetitionDetailsIntoFile(compIdCompDetailMap map[int64]model.CompetitionDetail) error {
	var err error
	var fseg []dto.CompetitionDetailFileSegment
	for k, v := range compIdCompDetailMap {
		fseg = append(fseg, dto.CompetitionDetailFileSegment{
			CompetitionId:     k,
			CompetitionDetail: v,
		})
	}
	file := dto.CompetitionDetailFile{
		CacheProperty:      createCacheFile(),
		CompetitionDetails: fseg,
	}
	path := getTempFolderPath() + "cds"
	content, err := json.Marshal(file)
	if err != nil {
		return err
	}
	log.Println("Writing file: " + path)
	err = ioutil.WriteFile(path, content, 0644)
	return err
}

func CreateTempFolderIfNotExists() {
	if isTempFolderDoesNotExists() {
		log.Printf("Temp folder does not exists, we are about to create it here: %s\n", getTempFolderPath())
		err := os.MkdirAll(getTempFolderPath(), 0700)
		if err != nil {
			log.Printf("Unable to create temp folder on path: %s, because: %s", getTempFolderPath(), err)
		}
	}
}

func unload(cdfs []dto.CompetitionDetailFileSegment) map[int64]model.CompetitionDetail {
	result := make(map[int64]model.CompetitionDetail)
	for _, v := range cdfs {
		result[v.CompetitionId] = v.CompetitionDetail
	}
	return result
}

func isRenewalNeeded(cacheProperty dto.CacheFile) bool {
	if common.IsStructEmpty(cacheProperty) || cacheProperty.IsOutdated() {
		return true
	}
	return false
}

func getTempFolderPath() string {
	tempDirEnvVar := os.Getenv(manualTempDirPathKey)
	if len(tempDirEnvVar) > 0 {
		if tempDirEnvVar[len(tempDirEnvVar)-1:] != string(os.PathSeparator) {
			return tempDirEnvVar + string(os.PathSeparator)
		}
		return tempDirEnvVar
	}
	return os.TempDir() + vrvTmpDirName
}

func isTempFolderDoesNotExists() bool {
	_, err := os.Stat(getTempFolderPath())
	return os.IsNotExist(err)
}

func createCacheFile() dto.CacheFile {
	return dto.CacheFile{
		Created: strconv.FormatInt(time.Now().Unix(), 10),
		Version: appVersion,
	}
}
