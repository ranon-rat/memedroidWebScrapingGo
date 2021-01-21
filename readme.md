# webScrappingWithGolang

its a simple web scrapper who extract memes from the page memedroid 

you can use it in your webpage because its a simple api.


# the body of the api

```json
{
	"memes": [
		{
			"title": "Going all in ",
			"imageURL": "https://images3.memedroid.com/images/UPLOADED932/60089f759e36b.jpeg"
		},
		{
			"title": "Caveman be like ",
			"imageURL": "https://images3.memedroid.com/images/UPLOADED190/6008a00f7630b.jpeg"
		},
		{
			"title": "DAMN JACKAL SNIPERS ",
			"imageURL": "https://images7.memedroid.com/images/UPLOADED289/60089fd1ade02.jpeg"
		}
	]
}
```
# example to work with the api  with js
```ts
fetch(url/api",{})// made the request
.then(r=>r.json())// decode the request into a json
.then(d=>{
	
	d.memes.map(i=>{
	// then go through the array
		console.log(`
		${i.title}
		${i.imageURL}
		`);// and print the array
	})

})
.catch(e=>console.log(e))// if something happend
```
