// import './App.css';
import React, { useState } from 'react'
import Sidebar from './component/Sidebar';
import './AppStyle.scss';
import Chat from './component/Chat';

function App() {

  return (
    <div className="App">
      <div className='navbar'>
        <div className='navbar-container'>
        <div className='navbar-text'>
            ChatGPT
        </div>
        </div>
    </div>
    {/* <div className='TitleText'>
    Hi! I'm ChatGPT-chan~
      </div> */}
      <div className='container'>
      {/* <button onClick={testClick}>button</button> */}
        
         
        <Sidebar/>
        <Chat/>
      </div>
    </div>
  );
}

export default App;
