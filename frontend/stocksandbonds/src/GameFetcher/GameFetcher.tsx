import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import {List, ListItem} from '@material-ui/core'


const URL = "https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks"


function GetGamelist() {
    const [gamelist, setGamelist] = useState([""])

    useEffect(() => {
        let games = fetchGameList()
        games.then(games => {
            setGamelist(games)})
    }, [])
  return (
      <div className="GameList">
    <Button 
    variant="outlined" 
    color="secondary" 
    onClick={() => {
        let games = fetchGameList()
        games.then(games => {
            setGamelist(games)})
    }}>
     Fetch the gamelist
    </Button>
<List>Gamelist: {toListItem(gamelist)}
</List>
    </div>
  );
}

async function fetchGameList() {
    const response = await fetch(URL, {
        method: 'GET'
    })
    const games = await response.json() as string[]
    return games
}

export default GetGamelist;

const toListItem = (list: string[]) => {
    let out = []
    for (var item of list) {
        out.push(
            <ListItem>{item}</ListItem>
        )
    }
    return out
}
