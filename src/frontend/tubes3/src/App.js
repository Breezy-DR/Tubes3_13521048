import React, { useState } from 'react'
import Sidebar from './component/Sidebar';
import './Styles.css';
import Chat from './component/Chat';
import io from 'socket.io-client'

const socket = io.connect("http://localhost:3001");

function App() {
  const current = new Date();
  const time = current.toLocaleTimeString("en-US");
  return (
    <div className="App">
      {/* <div className='container'></div> */}
      <div className='main-layout'>
      {/* <button onClick={testClick}>button</button> */}

        <Sidebar/>
        <Chat/>
      </div>
    </div>
  );
}

export default App;
