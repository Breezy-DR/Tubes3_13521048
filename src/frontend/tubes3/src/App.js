// import './App.css';
import React, { useState} from 'react'
import Sidebar from './component/Sidebar';
import './AppStyle.scss';
import Chat from './component/Chat';

function App() {
  const[data, setData] = useState([]);
  const[selectedAlg, setSelectedAlg] = useState('kmp');


  return (
    <div className="App">
      <div className='navbar'>
        ChatGPT
      </div>
    {/* <div className='TitleText'>
    Hi! I'm ChatGPT-chan~
      </div> */}
      <div className='container'>
        <Sidebar data={data} setData={setData} selectedAlg={selectedAlg} setSelectedAlg={setSelectedAlg}/>
        <Chat data={data} setData={setData} selectedAlg={selectedAlg}/>
      </div>
    </div>
  );
}

export default App;
