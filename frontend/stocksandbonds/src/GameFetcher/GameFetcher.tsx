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

export interface Game {
    GameName: string;
    Players: number;
    Day: number;
    PlayerList: string[];
    Gamestate: Gamestate;
}

export const newEmptyGame = () => ({
    GameName: "",
    Players: 0,
    Day: 0,
    PlayerList: [],
    Gamestate: emptyGamestate(),
})

export interface Gamestate {
    Prices: Map<string,number[]>;
    PlayerState: Map<string,Portfolio>;
}

const emptyGamestate = () => ({
    Prices: new Map<string,number[]>(),
    PlayerState: new Map<string,Portfolio>()
})

export interface Portfolio {
    Portfolio: Map<string,number[]>;
    Cash: number;
}

export function GetGamelist(gamelist: Game[], setGamelist: Function, hoverFunc: Function) {

    useEffect(() => {
       fetchGameList(setGamelist)

    }, [])
  return (
      <div className="GameList">
    <Button 
    variant="contained" 
    color="default" 
    onClick={function() {
        setGamelist([])
        fetchGameList(setGamelist)
    }}>
     Fetch the gamelist
    </Button>
<List>Gamelist: {createGiantTable(gamelist, hoverFunc)}
</List>
    </div>
  );
}

async function fetchGameList(setGameList: Function) {
    const response = await fetch(URL, {
        method: 'GET'
    })
    const games = await response.json() as Game[]
    setGameList(games)
}

// omg material UI table implementations suck ass
function createGiantTable(list: Game[], hoverFunc: Function) {
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
                    <TableCell align='right'></TableCell>
                </TableRow>
                </TableHead>
                <TableBody>
                    {list.map((row) => (
                        <TableRow hover key={row.GameName}>
                        <TableCell align='center'>{row.GameName ? row.GameName : ""}</TableCell>
                        <TableCell align='right'>{row.PlayerList ? row.PlayerList[0] : ""}</TableCell>
                        <TableCell align='right'>{row.Day ? row.Day : 0}</TableCell>
                        <TableCell align='right'>{row.PlayerList ? row.PlayerList : ""}</TableCell>
                        <TableCell><Button size='small' onClick={(e) => {hoverFunc(row)}}>Select</Button></TableCell>
                        </TableRow>
                    ))}
                </TableBody>
        </Table>
    </TableContainer>
    </Paper>
    </div>
    )
}
