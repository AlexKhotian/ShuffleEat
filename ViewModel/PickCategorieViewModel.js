function LoadCategorie(type) {
  $("#content").load("CategorieSearchResult.html");
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      var parsedData = JSON.parse(xhr.responseText);
      for (var i=0; i < parsedData._recipeTitles.length; i++) {
        var inputIng = document.createElement('input');
        inputIng.type = "submit";
        inputIng.setAttribute("name", "RecipeTitle");
        inputIng.value = parsedData._recipeTitles[i];
        var br = document.createElement("br");
        var container = document.getElementById("content");
        container.appendChild(inputIng);
        container.appendChild(br);
      }
    }
  }

  var data = JSON.stringify({"_categorie":type});
  console.log(data);
  var url = "PickCategorieRequest";
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-type", "application/json");
  xhr.send(data);
}
