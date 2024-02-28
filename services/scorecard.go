package services

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScoreCard struct {
	gorm.Model
	Tee      string  `json:"tee"`
	Par      int     `json:"par"`
	Rating   float64 `json:"rating"`
	Slope    int     `json:"slope"`
	Holes    []Hole  `json:"holes"`
	CourseID uint    `json:"course_id"`
}

type Hole struct {
	gorm.Model
	Number      int `json:"number"`
	Par         int `json:"par"`
	StrokeIndex int `json:"stroke_index"`
	ScoreCardID uint
}

type ServiceScoreCard struct {
	ScoreCard ScoreCard
	DB        *gorm.DB
}

func NewServiceScoreCard(scoreCard ScoreCard, db *gorm.DB) *ServiceScoreCard {
	return &ServiceScoreCard{
		ScoreCard: scoreCard,
		DB:        db,
	}
}

func (sc *ServiceScoreCard) CreateScoreCard(scoreCard ScoreCard) error {
	card := ScoreCard{
		Tee:      scoreCard.Tee,
		Par:      scoreCard.Par,
		Rating:   scoreCard.Rating,
		Slope:    scoreCard.Slope,
		Holes:    scoreCard.Holes,
		CourseID: scoreCard.CourseID,
	}
	fmt.Println(card)

	if err := sc.DB.Create(&card).Error; err != nil {
		return err
	}

	return nil
}

func (sc *ServiceScoreCard) APICreateScoreCard(scoreCard *ScoreCard) (*ScoreCard, error) {
	if err := sc.DB.Create(&scoreCard).Error; err != nil {
		return scoreCard, err
	}

	return scoreCard, nil
}

func (sc *ServiceScoreCard) GetAllScoreCards() ([]*ScoreCard, error) {
	var scoreCards []*ScoreCard

	if res := sc.DB.Find(&scoreCards); res.Error != nil {
		return nil, res.Error
	}

	fmt.Println(&scoreCards)

	return scoreCards, nil
}

func (sc *ServiceScoreCard) GetScoreCardById(id int) (ScoreCard, error) {
	var scoreCard []*ScoreCard

	if res := sc.DB.Preload(clause.Associations).Find(&scoreCard, id); res.Error != nil {
		return ScoreCard{}, res.Error
	}

	return *scoreCard[0], nil
}
