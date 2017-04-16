package DataBase

// MeasurementMetric defines type of metric for Ingredient
type MeasurementMetric int16

// in which metric type our Ingredient
const (
	Gramm MeasurementMetric = 0
	Liter MeasurementMetric = 1
)

// Categorie of dish
type Categorie int32

// To which Categorie belongs your dish
const (
	LowKall   = 0
	Breakfest = 1
	Lunch     = 3
	Dinner    = 4
	Salat     = 5
	Meat      = 6
)

// Ingredient is food type (eg. tomato)
type Ingredient struct {
	TypeID          uint64            `bson:"_id" json:"typeID"`
	Name            string            `bson:"_name" json:"name"`
	Quantity        uint64            `bson:"_quantity" json:"quantity"`
	MeasurementType MeasurementMetric `bson:"_measurementType" json:"measurementType"`
}

// Recipe contains in about recipe
type Recipe struct {
	RecipeID      uint64       `bson:"_id" json:"_recipeID"`
	Ingredients   []Ingredient `bson:"_ingredients" json:"_ingredients"`
	Description   []string     `bson:"_description" json:"_description"`
	CategorieType Categorie    `bson:"_categorieType" json:"categorieType"`
}
