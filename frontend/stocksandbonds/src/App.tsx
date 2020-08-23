import React, { useState } from 'react';
import logo from './images/moneybag.svg';
import './App.css';
import {GetGamelist, Game, newEmptyGame} from './GameFetcher/GameFetcher'
import NameEntry from './Layout/NameEntry'
import { AppBar, Button, CircularProgress, TextField } from '@material-ui/core';
import CheckIcon from '@material-ui/icons/Check';
import ClearIcon from '@material-ui/icons/Clear';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const gameEndpoint = "https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks"

export var CompanyList: string[] = [
	"Central City",
	"Growth Corporation",
	"Metro Properties",
	"Pioneer Mutual",
	"Shady Brooks",
	"Stryker Drilling",
	"Tri-City Transport",
	"United Auto",
	"Uranium Enterprises",
	"Valley Power",
]

function App() {

  const [playerName, updateName] = useState("example user name")
  let emptyGameList: Game[] = [];
  const [gamelist, setGamelist] = useState(emptyGameList)
  let emptyGame: Game = newEmptyGame()
  const [currentSelectedGame, setCurrentSelectedGame] = useState(emptyGame)

  return (
    <div className="App">
      <AppBar color='inherit'>
      <div className="TopBar">
        <img src={logo} className="App-logo" alt="logo" />
        Stocks And Bonds
        <div className='credits'>
        A game NOT by Peet the Cheat and Brian "Prank God" Krasnopolsky
        </div>
      </div>
      </AppBar>
      {GetGamelist(gamelist, setGamelist, setCurrentSelectedGame)}{renderCurrentGameTable(currentSelectedGame)}
      {CreateGame(playerName)}
      <div className="nameentry">
      {NameEntry(updateName)}
      {"Your Name: "}
      {playerName}
      </div>
    </div>
  );
}

export default App;

function CreateGame(playerName: string) {

  const [gameCreated, updateGameCreated] = useState(0)
  const [gameNameEntry, updateGameNameEntry] = useState("")

  return (
    <div className="creategame">
       <TextField placeholder="New Game Name" id="outlined-basic" label="New Game Name" variant="outlined" 
      onChange={(e) => {
        updateGameNameEntry(e.target.value)
      }}
    />
      <div>
        <Button
    variant="outlined" 
    color="primary" 
    onClick={() => {
      createGamePOST(gameNameEntry, playerName, updateGameCreated)
      }}>
      Create Game
      </Button>
      {checkGameCreation(gameCreated)}
      </div>
    </div>

  )
}

/*
-2  Input error
-1  Request Failure / Name Taken
0   Blank / Default
1   Loading
2   Success
*/

function checkGameCreation(gameStatus: number) {
  if (gameStatus === 1) {
    return <div><CircularProgress/></div>
  }
  if (gameStatus === -1) {
    return <div><ClearIcon/>Game name already taken?</div>
  }
  if (gameStatus === 2) {
    return <div><CheckIcon/></div>
  }
  if (gameStatus === -2) {
    return <div><ClearIcon/>Max Game Name Length is 20</div>
  }
  return
}

async function createGamePOST(gamename: string, playerName: string, updateGameStatus: Function) {
  updateGameStatus(1)
  if (gamename.length > 20) {
    updateGameStatus(-2)
    return
  }
  console.log(gamename)
  if (gamename === "" || playerName == "") {
    updateGameStatus(0)
    return
  }


  const requestBody = {
    requestType: "creategame",
    requester: playerName,
    gameState: {
      gameName: gamename
    }
  }

  const response = await fetch(gameEndpoint, {
    method: 'POST',
    body: JSON.stringify(requestBody),
    headers: {"Content-Type": "application/json"},
  })

  if (response.status === 201) {
    updateGameStatus(2)
  } else {
    updateGameStatus(-1)
  }

}

function renderCurrentGameTable(game: Game) {
  console.log(game)
  return (
    <div className='CurrentGameTable'>
        <Paper>
<TableContainer>
    <Table stickyHeader size="small">
        <TableHead>
            <TableRow>
              {CompanyList.map((company:string) => (
                <TableCell align='left'>{company}</TableCell>
              ))}
            </TableRow>
            </TableHead>
            <TableBody>
              <TableRow>
                {game.Gamestate.Prices && game.GameName !== "" ?
                CompanyList.map((company: string) => (
                  <TableCell align='center'>{game.Gamestate?.Prices.get(company)}</TableCell>
                )) : 
                <TableCell align='center'>Not Loaded</TableCell>}

              </TableRow>
            </TableBody>
    </Table>
</TableContainer>
</Paper>
</div>
)
}