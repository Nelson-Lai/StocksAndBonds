import React from 'react';
import logo from './logo.svg';
import './App.css';
import GetGamelist from './GameFetcher/GameFetcher'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        {GetGamelist()}
      </header>
    </div>
  );
}

export default App;
