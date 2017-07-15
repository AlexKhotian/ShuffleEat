package ServerHandler

// RecipeForm contains in about recipe from html form
type RecipeForm struct {
	RecipeTitle string   `json:"_title"`
	Ingredients []string `json:"_ingredients"`
	Quantities  []string `json:"_quantities"`
	Description string   `json:"_description"`
	Categorie   uint16   `json:"_categorie"`
}

// CategorieSearchRequest send by html form, when user asks for seatch by categories
type CategorieSearchRequest struct {
	CategorieType uint16 `json:"_categorie"`
}

// CategorieSearchResponse send by CategorieSearchHandler
type CategorieSearchResponse struct {
	RecipeTitles []string `json:"_recipeTitles"`
}

const dbName = "Recipes_v2_test"
const dbCollection = "RecipesCollection_v2_test"
