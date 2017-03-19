package DataBase

type MeasurementMetric int16

// in which metric type our Ingredient
const (
	Gramm MeasurementMetric = 0
	Liter MeasurementMetric = 1
)

// Ingredient is food type (eg. tomato)
type Ingredient struct {
	TypeId          uint64            `bson:"_id" json:"typeId"`
	Name            string            `bson:"_name" json:"name"`
	Quantity        uint64            `bson:"_quantity" json:"quantity"`
	MeasurementType MeasurementMetric `bson:"_measurementType" json:"measurementType"`
}

type Recipe struct {
	RecipeId    uint64       `bson:"_id" json:"RecipeId"`
	Ingredients []Ingredient `bson:"_ingredients" json:"_ingredients"`
	Description []string     `bson:"_description" json:"_description"`
}
