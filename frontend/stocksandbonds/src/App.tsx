import React, { useState } from 'react';
import logo from './images/moneybag.svg';
import './App.css';
import GetGamelist from './GameFetcher/GameFetcher'
import NameEntry from './Layout/NameEntry'
import { AppBar, Button, CircularProgress, TextField } from '@material-ui/core';
import CheckIcon from '@material-ui/icons/Check';
import ClearIcon from '@material-ui/icons/Clear';

const gameEndpoint = "https://eyu6c6iiy3.execute-api.us-east-2.amazonaws.com/development/stocks"


function App() {

  const [name, updateName] = useState("example user name")

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
      {GetGamelist()}
      {CreateGame()}
      <div className="nameentry">
      {NameEntry(updateName)}
      {name}
      </div>
    </div>
  );
}

export default App;

function CreateGame() {

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
    onClick={() => {createGamePOST(gameNameEntry, updateGameCreated)}}>
      Create Game
      </Button>{checkGameCreation(gameCreated)}
      </div>
    </div>

  )
}

function checkGameCreation(gameStatus: number) {
  if (gameStatus === 1) {
    return <div><CircularProgress/></div>
  }
  if (gameStatus === -1) {
    return <div><ClearIcon/></div>
  }
  if (gameStatus === 2) {
    return <div><CheckIcon/></div>
  }
  if (gameStatus === -2) {
    return <div><ClearIcon/>Max Game Name Length is 15</div>
  }
  return
}

async function createGamePOST(gamename: string, updateGameStatus: Function) {
  updateGameStatus(1)
  if (gamename.length > 15) {
    updateGameStatus(-2)
    return
  }
  console.log(gamename)
  if (gamename === "") {
    updateGameStatus(0)
    return
  }


  const requestBody = {
    requestType: "creategame",
    requester: "memer420",
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