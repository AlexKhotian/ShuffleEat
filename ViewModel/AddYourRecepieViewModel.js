function ShowCategorieDropDown() {
  document.getElementById("")
}

function AddIngredient() {
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

  var inputQua = document.createElement('input');
  inputQua.type = "text";
  inputQua.setAttribute("name", "Quantity");

  var br = document.createElement("br");

  var container = document.getElementById("EditsArea");

  container.appendChild(br);
  container.appendChild(labelIng);
  container.appendChild(inputIng);
  container.appendChild(labelQua);
  container.appendChild(inputQua);
}

function CommitRecepie() {
  var form = $(document.myform);
  var elements = document.getElementsByTagName('input');
  var recepieTitle;
  var ingredients = [];
  var quantities = [];
  var description;
  var categorie;
  console.log(elements.length);
  for (var i=0; i<elements.length; i++){
    var name = elements[i].getAttribute("name");
    if (name == "recepieTitle"){
      recepieTitle = elements[i].value;
    } else if (name == "Ingredient") {
      ingredients.push(elements[i].value);
    } else if (name == "Quantity") {
      quantities.push(elements[i].value);
    } else if (name == "description") {
      description = elements[i].value;
    }
  }
  var sel = document.getElementById("categorieSelect");
  var categorie = sel.selectedIndex;

  var data = JSON.stringify({"_title":recepieTitle,
   "_ingredients":ingredients, "_quantities":quantities, "_description":description, "_categorie":categorie});
  console.log(data);
  var url = "url";
  xhr = new XMLHttpRequest();
  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-type", "application/json");
  xhr.send(data);
}
