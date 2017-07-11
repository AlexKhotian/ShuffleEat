function LoadRecipe() {
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
      var parsedData = JSON.parse(xmlHttp.responseText);
      var desription = document.getElementsByTagName('description');
      desription.value = parsedData._description;
      var categorie = document.getElementsByTagName('categorie');
      categorie.value = parsedData._categorie
      var recipeTitle = document.getElementsByTagName('recipeTitle');
      recipeTitle.value = parsedData._title
      for (var i=0; i < parsedData._ingredients.length; i++) {
        var labelIng = document.createElement('label');
        var textIng = document.createTextNode("Ingredient:");
        labelIng.setAttribute("for", "Ingredient:");
        labelIng.appendChild(textIng);

        var labelQua = document.createElement('label');
        var textQua = document.createTextNode(" Quantity:");
        labelQua.setAttribute("for", " Quantity:");
        labelQua.appendChild(textQua);

        var inputIng = document.createElement('input');
        inputIng.type = "text";
        inputIng.setAttribute("name", "Ingredient");
        inputIng.value = parsedData._ingredients[i];
        inputIng.readonly = true;

        var inputQua = document.createElement('input');
        inputQua.type = "text";
        inputQua.setAttribute("name", "Quantity");
        inputQua.value = parsedData._quantities[i];
        inputQua.readonly = true;

        var br = document.createElement("br");

        var container = document.getElementById("EditsArea");

        container.appendChild(br);
        container.appendChild(labelIng);
        container.appendChild(inputIng);
        container.appendChild(labelQua);
        container.appendChild(inputQua);
      }
    }
  }
  var randomMealUrl = "RandomMeal"
  xmlHttp.open("GET", randomMealUrl, true); // true for asynchronous
  xmlHttp.send(null);
}
$(document).ready(LoadRecipe)
