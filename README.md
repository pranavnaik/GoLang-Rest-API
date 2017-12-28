# GoLang-Rest-API
Rest API in GoLang

Author: Pranav Naik

#Instructions:
==============
- The API uses the external package from `https://github.com/gorilla/mux/` for routing and parsing the incoming request.
- The API is congiured to run on port 8081
- The API has only one end point i.e `/check-inventory-levels/" and takes three parameters /{ItemName}/{ItemId}/{Quantity}/`
- example Local API address:  `http://localhost:8081/check-inventory-levels/coffee/2/601/` 
- The quantity to check againt is hardcoded to 100units for all items
- The response will be in JSON format telling whether the queried quanity is present or not..
- example API response => `[{"Id":2,"Name":"coffee","quantity":0,"IsAvailable":1}] `
- where if `IsAvailable = 1` then we have that qanity in our stock.
  
#Usage:
=======
- Download or clone the repository to your GoWorkspace directory.
- Open and run the http server present in RestAPI folder
- Open the browser and fire the appropriate link for example "http://localhost:8081/check-inventory-levels/coffee/2/601/ "
- You should get the response in JSON format like `[{"Id":2,"Name":"coffee","quantity":0,"IsAvailable":1}]`
