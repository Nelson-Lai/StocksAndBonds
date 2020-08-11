import React, { useState } from 'react';
import logo from './images/moneybag.svg';
import './App.css';
import GetGamelist from './GameFetcher/GameFetcher'
import NameEntry from './Layout/NameEntry'
import { AppBar } from '@material-ui/core';


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
      <header className="App-header">


      
      {NameEntry(updateName)}
      {name}
      </header>
    </div>
  );
}

export default App;
