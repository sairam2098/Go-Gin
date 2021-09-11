
function getByCategory(){
    document.getElementById('ListToDo').innerHTML = "<tr><th>Name</th><th>Price</th><th>Quantity</th></tr>";
    var category = document.getElementById('ToDoId').value;
    const url = 'http://localhost:7999/items/by-category?category=' + category ;
    fetch(url)
        .then(resp => resp.json())
        .then((todosA)=>{
            document.getElementById('Description').innerHTML = "<b>List of the " + category + " :</b>";
            todosA.forEach((todo)=> {
                list = "<tr><th>" + todo.Name + "</th><th>" + todo.Price + "</th><th>" + todo.Quantity + "</th></tr>";
                document.getElementById("ListToDo").innerHTML += list;
            })
        }).catch(function(error){
            document.getElementById('ListToDo').innerHTML = "<tr><th>Name</th><th>Price</th><th>Quantity</th></tr>";
            alert("Internal Error")
        });
}

function postItem(){
    var categoryObj = document.getElementById('NewCategory').value;
    var nameObj = document.getElementById('NewName').value;
    var priceObj = document.getElementById('NewPrice').value;
    var quantityObj = document.getElementById('NewQuantity').value;
    
    const data = { category: categoryObj, name: nameObj, price: priceObj, quantity: quantityObj };

    fetch('http://localhost:7999/newitems', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    }).then((resp) => {
        if(resp.status != 200) {
            alert("Internal Error");
        }
   });
}