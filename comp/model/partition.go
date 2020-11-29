package model

type Partition struct {
	Tags         string          `json:"tags"`
	Gender       PartitionGender `json:"gender"`
	AgeGroup     string          `json:"ageGroup"`
	ClimbingType ClimbingType    `json:"climbingType"`
	Sections     []int64         `json:"sections"`
	Results      []Result        `json:"results"`
}
