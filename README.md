APIS:

POST `/players`<br/>
To insert a new player
```json
{
    "name":"player3",
    "country":"UK",
    "score":1
}
```

GET `/players`<br/>
Displays the list of all players in descending order


DELETE `/players/:id`<br/>
Deletes the player entry


PUT `/players/:id`<br/>
Updates the player attributes. Only name and
score can be updated
```json
{
    "name":"Updated Name",
    "score":1
}
```


GET `/players/rank/:val`<br/>
Fetches the player ranked “val”


GET `/players/random`<br/>
Fetches a random player
