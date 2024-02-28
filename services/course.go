package services

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewServiceCourse(c Course, db *gorm.DB) *ServiceCourse {
	return &ServiceCourse{c, db}
}

type Course struct {
	gorm.Model
	Name       string `json:"name"`
	Par        int    `json:"par"`
	ScoreCards []ScoreCard
}

type ServiceCourse struct {
	Course Course
	DB     *gorm.DB
}

func (cs *ServiceCourse) CreateCourse(course Course) error {
	c := Course{
		Name: course.Name,
		Par:  course.Par,
	}

	if err := cs.DB.Create(&c).Error; err != nil {
		return err
	}

	return nil
}

func (cs *ServiceCourse) GetAllCourses() ([]*Course, error) {
	var courses []*Course

	if res := cs.DB.Find(&courses); res.Error != nil {
		return nil, res.Error
	}

	fmt.Println(&courses[0])

	return courses, nil
}

func (cs *ServiceCourse) GetCourseById(id int) (Course, error) {
	var courses []*Course

	if res := cs.DB.Preload(clause.Associations).Find(&courses, id); res.Error != nil {
		return Course{}, res.Error
	}

	fmt.Println(*courses[0])

	return *courses[0], nil
}

func (cs *ServiceCourse) UpdateCourse(course Course) error {
	existingCourse := new(Course)

	if res := cs.DB.Find(&existingCourse, "id = ?", course.ID); res.Error != nil {
		return res.Error
	}

	if course.Name != "" {
		existingCourse.Name = course.Name
	}

	if course.Par != 0 {
		existingCourse.Par = course.Par
	}

	if err := cs.DB.Save(&existingCourse).Error; err != nil {
		return err
	}

	return nil
}
