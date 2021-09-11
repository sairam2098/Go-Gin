
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