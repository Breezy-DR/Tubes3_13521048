import React from 'react'
import GPTMessage from './GPTMessage';
import UserMessage from './UserMessage';
import { useState } from 'react';

function Chat() {
  
  const[data, setData] = useState([]);
  const [inputValue, setInputValue] = useState('');


  const QUESTION_API = "http://ec2-54-169-32-134.ap-southeast-1.compute.amazonaws.com:8080/question"

  const handleSubmit = (event) => {
    event.preventDefault();
    data.push({from:'user', chat:event.target.value})
    ask(QUESTION_API,
        {"question": event.target.value, "search_algorithm":"kmp", "session_id": sessionStorage.getItem("session_id")},
        x)
        setInputValue('');
  };

  function x(json) {
    sessionStorage.setItem("session_id", json["session_id"]);
    let responseArr = json['response']
    let newData = [...data]
    responseArr.forEach(item => {
      newData.push({from:'server', chat:item})
      console.log(item)
    })
    setData(newData)
  }

  function ask(url, data, callback) {
    console.log(JSON.stringify(data))
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(json => callback(json))
        .catch(error => console.error(error));
  }

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      handleSubmit(event);
    }
  }

  return (
    <div className='chat'>
      <div className='conversation'>
      {
        data && data.map((item) => {
          if(item.from === "user"){
            return <GPTMessage dataToMessage={item.chat}/>
          }else{
            return <UserMessage dataToMessage={item.chat}/>
          }
        })
      }
        <div className='flex-container'></div>
      </div>
      <div className='ChatForm'>
        <form onSubmit={handleSubmit}>
          <input type='text' value={inputValue} placeholder='Masukkan pertanyaan Anda (press enter to submit)' onKeyDown={handleKeyDown}
          onChange={(event) => setInputValue(event.target.value)}
          />
        </form>
      </div>
    </div>
  )
}

export default Chat