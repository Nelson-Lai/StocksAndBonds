import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import {List, ListItem, CircularProgress} from '@material-ui/core'


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
    variant="contained" 
    color="default" 
    onClick={async function() {
        setGamelist([])
        let games = await fetchGameList()
        setGamelist(games)
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

function toListItem(list: string[]) {
    if (list.length === 0) {
        return <CircularProgress />
    }
    let out = []
    for (var item of list) {
        out.push(
            <ListItem>{item}</ListItem>
        )
    }
    return out
}
