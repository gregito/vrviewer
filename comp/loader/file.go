package loader

import (
	"encoding/json"
	"fmt"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
	"github.com/gregito/vrviewer/comp/model"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	vrvTmpDirName        = "vrvTmp/"
	manualTempDirPathKey = "VRVTD"
)

func CompetitionsFromFile() ([]dto.Competition, error) {
	var result []dto.Competition
	fileName := "cmps"
	log.Println("Reading CompetitionDetail from: " + getTempFolderPath() + fileName)
	file, err := ioutil.ReadFile(getTempFolderPath() + fileName)
	if err != nil {
		log.Printf("Unable to read file to obtain stored competition(s): %s", err)
		return result, err
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		log.Println("Unable to unmarshall competition(s): ", err)
		return result, err
	}
	return result, nil
}

func CompetitionDetailFromFile(competitionId int64) (model.CompetitionDetail, error) {
	var result model.CompetitionDetail
	fileName := fmt.Sprintf("%scd%s", getTempFolderPath(), strconv.Itoa(int(competitionId)))
	log.Println("Reading CompetitionDetail from: " + getTempFolderPath() + fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Unable to read competition detail from file by name \"%s\" because: %s\n", fileName, err)
		return result, err
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		log.Println("Unable to unmarshall competition detail: ", err)
		return result, err
	}
	return result, nil
}

func WriteCompetitionsIntoFile(comps []dto.Competition) error {
	var err error
	content, err := json.Marshal(comps)
	if err != nil {
		return err
	}
	file := getTempFolderPath() + "cmps"
	log.Println("Writing file: " + file)
	err = ioutil.WriteFile(file, content, 0644)
	return err
}

func WriteCompetitionDetailIntoFile(competitionId int64, cd model.CompetitionDetail) error {
	var err error
	content, err := json.Marshal(cd)
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%scd%s", getTempFolderPath(), strconv.Itoa(int(competitionId)))
	log.Println("Writing file: " + file)
	err = ioutil.WriteFile(file, content, 0644)
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
