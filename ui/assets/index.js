var responseData = [];

var dataset = [];

async function getObjects(){
    const response = await fetch("/api/views");
    const data = await response.json();
    responseData = data.Posts;
    responseData.forEach((element) =>{
        dataset.push(element.Description);
    })
    for(let i = dataset.length-1; i > 0; i--){
        const newIndex = Math.floor(Math.random()*(i+1))
        const oldValue = dataset[newIndex];
        dataset[newIndex] = dataset[i];
        dataset[i] = oldValue;
    }
}



// async function shuffle(){
//     await addData();
//     for(let i = dataset.length-1; i > 0; i--){
//         const newIndex = Math.floor(Math.random()*(i+1))
//         const oldValue = dataset[newIndex];
//         dataset[newIndex] = dataset[i];
//         dataset[i] = oldValue;
//     }
// }
// shuffle();


const mainContainer = document.getElementById('#main');

async function addToHTML(){
    await getObjects();
    for(let i = 0; i < dataset.length; i++)
    {
        const value = marked.parse(dataset[i])
        main.innerHTML += `<div class="preview-area">${value}</div>`
    }
}

addToHTML();


