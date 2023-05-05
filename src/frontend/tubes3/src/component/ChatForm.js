import React from 'react'
import { useState } from 'react'
import { useContext } from 'react';
import SelectedValueContext from './SelectedValueContext';


function ChatForm() {
  const [message, setInputValue] = useState('');
  const handleChange = (event) => {
    setInputValue(event.target.value);
  }
  const selectedValue = useContext(SelectedValueContext);

  const handleSubmit = (event) => {
    event.preventDefault();
    // Send input (message) to database here
    // ....
    sessionStorage.getItem("session_id");
    sendGetRequestWithJsonBody("http://localhost:8080/question",
        {"question": event.target.value,
        "search_algorithm":{selectedValue},
        "session_id": "7bbf04e082"},
        x)
  };

  function x(json) {
    sessionStorage.setItem("session_id", "example");
    // var myObject = JSON.parse(json)
    console.log(json["session_id"])

  }

  function sendGetRequestWithJsonBody(url, data, callback) {
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
    <div className='ChatForm'>
      <form onSubmit={handleSubmit}>
      <input type='text' name="message" placeholder='Masukkan pertanyaan Anda (press enter to submit)' onChange={handleChange} onKeyDown={handleKeyDown}/>
      </form>
    </div>
  )
}

export default ChatForm