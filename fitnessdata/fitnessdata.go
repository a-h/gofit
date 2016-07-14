package fitnessdata

import (
	"errors"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/fitness/v1"
)

// Get uses the Google API to extract fitness data from the REST endpoint.
func Get(conf *oauth2.Config, token *oauth2.Token) (*FitnessData, error) {
	client := conf.Client(oauth2.NoContext, token)

	svc, err := fitness.New(client)
	if err != nil {
		return nil, errors.New("Unable to create Fitness service: " + err.Error())
	}

	d := &FitnessData{}

	now := time.Now()

	d.StartDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, -7)
	d.Days = 7
	d.DayNames = getDayNames(d.StartDate, 7)

	_, ceoa, err := aggregateData(svc, d.StartDate, d.Days, "derived:com.google.calories.expended:com.google.android.gms:from_activities")

	if err != nil {
		return d, errors.New("failed to retrieve the calories expended on activities: " + err.Error())
	}

	d.CaloriesExpendedOnActivities = ceoa

	_, ceobmr, err := aggregateData(svc, d.StartDate, d.Days, "derived:com.google.calories.expended:com.google.android.gms:from_bmr")

	if err != nil {
		return d, errors.New("Failed to retrieve the calories expended on bmr: " + err.Error())
	}

	d.CaloriesExpendedOnBMR = ceobmr

	_, h, err := aggregateData(svc, d.StartDate, d.Days, "derived:com.google.height:com.google.android.gms:merge_height")

	if err != nil {
		return d, errors.New("Failed to retrieve height data: " + err.Error())
	}

	d.Height = h

	steps, _, err := aggregateData(svc, d.StartDate, d.Days, "derived:com.google.step_count.delta:com.google.android.gms:estimated_steps")

	if err != nil {
		return d, errors.New("Failed to retrieve step data: " + err.Error())
	}

	d.Steps = steps

	_, weight, err := aggregateData(svc, d.StartDate, d.Days, "derived:com.google.weight:com.google.android.gms:merge_weight")

	if err != nil {
		return d, errors.New("Failed to retrieve weight data: " + err.Error())
	}

	d.Weight = weight

	return d, nil
}

func getDayNames(start time.Time, days int) []string {
	dayNames := []string{}
	for d := 0; d < days; d++ {
		day := start.AddDate(0, 0, d)
		dayNames = append(dayNames, day.Weekday().String())
	}
	return dayNames
}

func aggregateData(svc *fitness.Service, startDate time.Time, days int, dataSourceID string) (intdata []int64, floatdata []float64, err error) {
	ar := &fitness.AggregateRequest{
		StartTimeMillis: unixMilliseconds(startDate),
		EndTimeMillis:   unixMilliseconds(startDate.AddDate(0, 0, days)),
		AggregateBy: []*fitness.AggregateBy{&fitness.AggregateBy{
			DataSourceId: dataSourceID,
		}},
		BucketByTime: &fitness.BucketByTime{
			Period: &fitness.BucketByTimePeriod{
				Type:       "day", // Other values: "day", "month", "week"
				Value:      1,
				TimeZoneId: "UTC",
			},
		},
	}

	results, err := svc.Users.Dataset.Aggregate("me", ar).Do()

	if err != nil {
		return nil, nil, err
	}

	// Because we're bucketing by time, each bucket contains a single dataset corresponding to
	// the day, with a single point and value in each dataset.
	for _, b := range results.Bucket {
		for _, ds := range b.Dataset {
			for _, p := range ds.Point {
				for _, v := range p.Value {
					floatdata = append(floatdata, v.FpVal)
					intdata = append(intdata, v.IntVal)
				}
			}
		}
	}

	return intdata, floatdata, nil
}

func fromUnixMilliseconds(milliseconds int64) time.Time {
	return time.Unix(milliseconds/1000, 0)
}

func unixMilliseconds(time time.Time) int64 {
	return time.UnixNano() / 1000000
}

// FitnessData represents fitness data extracted from Google Fit from a time period.
type FitnessData struct {
	StartDate                    time.Time `json:"startDate"`
	Days                         int       `json:"days"`
	DayNames                     []string  `json:"dayNames"`
	Steps                        []int64   `json:"steps"`
	Weight                       []float64 `json:"weight"`
	Height                       []float64 `json:"height"`
	CaloriesExpendedOnActivities []float64 `json:"caloriesExpendedOnActivities"`
	CaloriesExpendedOnBMR        []float64 `json:"caloriesExpendedOnBMR"`
}
