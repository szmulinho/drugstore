fetch(`http://localhost:8081/drugs/`)
.then(res => res.json())
.then(data => console.log(data))