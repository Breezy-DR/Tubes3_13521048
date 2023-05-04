import React from 'react'
import { useState } from 'react'

function ChatForm() {
  const [message, setInputValue] = useState('');
  const handleChange = (event) => {
    setInputValue(event.target.value);
  }

  const handleSubmit = (event) => {
    event.preventDefault();
    // Send input (message) to database here
    // ....
  };

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      handleSubmit(event);
    }
  }

  const sendData = async () => {
    const url = "ec2-54-169-32-134.ap-southeast-1.compute.amazonaws.com:8080/question/";
    const data = {
      question: "ping? namx 5 nama ikan? 2*9^(7-2); 7 nama ikan; sebutkan 7 nama ikan; 2023/11/12?pong",
      search_algorithm: "bm",
      session_id: "7bbf04e082"
    };
    const options = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
    const response = await fetch(url, options);
    const jsonResponse = await response.json();
    console.log(jsonResponse);
  };

  return (
    <div className='ChatForm'>
      <form onSubmit={sendData}>
      <input type='text' name="message" placeholder='Masukkan pertanyaan Anda (press enter to submit)' onChange={handleChange} onKeyDown={handleKeyDown}/>
      </form>
    </div>
  )
}

export default ChatForm