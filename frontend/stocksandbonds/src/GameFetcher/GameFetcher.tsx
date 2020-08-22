import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import {List, ListItem, CircularProgress, Tab} from '@material-ui/core'
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';


const URL = "https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks"

interface Game {
    GameName: string;
    Players: number;
    Day: number;
    PlayerList: string[];
    Gamestate: Gamestate;
}

interface Gamestate {
    Prices: Map<string,number[]>;
    PlayerState: Map<string,Portfolio>;
}

interface Portfolio {
    Portfolio: Map<string,number[]>;
    Cash: number;
}

function GetGamelist() {
    let emptyGame: Game[] = [];
    const [gamelist, setGamelist] = useState(emptyGame)

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
<List>Gamelist: {createGiantTable(gamelist)}
</List>
    </div>
  );
}

async function fetchGameList() {
    const response = await fetch(URL, {
        method: 'GET'
    })
    const games = await response.json() 
    console.log(games)
    return games as Game[]
}

export default GetGamelist;


// omg material UI table implementations suck ass
function createGiantTable(list: Game[]) {
    return (
        <div className='GameTable'>
            <Paper>
    <TableContainer>
        <Table stickyHeader size="small">
            <TableHead>
                <TableRow>
                    <TableCell align='center'>Game Name</TableCell>
                    <TableCell align='right'>Game Creator</TableCell>
                    <TableCell align='right'>Day</TableCell>
                    <TableCell align='right'>Players</TableCell>
                </TableRow>
                </TableHead>
                <TableBody>
                    {list.map((row) => (
                        <TableRow hover key={row.GameName}>
                        <TableCell align='center'>{row.GameName ? row.GameName : ""}</TableCell>
                        <TableCell align='right'>{row.PlayerList ? row.PlayerList[0] : ""}</TableCell>
                        <TableCell align='right'>{row.Day ? row.Day : 0}</TableCell>
                        <TableCell align='right'>{row.PlayerList ? row.PlayerList : ""}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
        </Table>
    </TableContainer>
    </Paper>
    </div>
    )
}
