package handlers

// func (sc *ScoreCardServiceHandler) APICreateScoreCard(c echo.Context) error {
// 	cs := new(services.ScoreCard)
//
// 	if err := c.Bind(cs); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
//
// 	scorecard := &services.ScoreCard{
// 		Tee:      cs.Tee,
// 		Par:      cs.Par,
// 		Rating:   cs.Rating,
// 		Slope:    cs.Slope,
// 		CourseID: cs.CourseID,
// 		Holes:    cs.Holes,
// 	}
//
// 	if err := sc.DB.Create(&scorecard).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
//
// 	response := map[string]interface{}{
// 		"data": &scorecard,
// 	}
//
// 	return c.JSON(http.StatusCreated, response)
// }
