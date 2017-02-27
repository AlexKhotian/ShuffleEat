package DataBase

type MeasurementMetric int16

// in which metric type our Ingredient
const (
	Gramm MeasurementMetric = 0
	Liter MeasurementMetric = 1
)

// Ingredient is food type (eg. tomato)
type Ingredient struct {
	name            string
	quantity        uint64
	measurementType MeasurementMetric
}

type Recipe struct {
	recipeId uint64
}
