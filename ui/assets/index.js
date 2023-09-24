fetch("/api/views")
.then(r =>  r.json().then(data => ({status: r.status, body: data})))
.then(obj => console.log(obj));