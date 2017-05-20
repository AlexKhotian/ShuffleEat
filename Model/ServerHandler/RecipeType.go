package ServerHandler

// RecipeForm contains in about recipe from html form
type RecipeForm struct {
	RecipeTitle string   `json:"_title"`
	Ingredients []string `json:"_ingredients"`
	Quantities  []string `json:"_quantities"`
	Description string   `json:"_description"`
	Categorie   uint16   `json:"_categorie"`
}

const dbName = "Recipes"
const dbCollection = "RecipeCollection"
