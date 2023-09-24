fetch(`/api/views`)
   .then(response => response.json())
   .then(json => console.log(JSON.stringify(json)))

