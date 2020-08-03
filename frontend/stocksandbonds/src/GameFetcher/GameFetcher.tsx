import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import {List, ListItem} from '@material-ui/core'


const URL = "https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks"


function GetGamelist() {
    const [gamelist, setGamelist] = useState([""])
  return (
      <div>
    <Button 
    variant="outlined" 
    color="secondary" 
    onClick={() => {
        let games = fetchGameList()
        games.then(games => {
            console.log(games)
            setGamelist(games)})
    }}>
     Fetch the gamelist
    </Button>
<List>Gamelist: {gamelist}
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
