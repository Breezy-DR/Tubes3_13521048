import React from 'react'
import { useState } from 'react'

function ChatForm({childToParent}) {
  const [message, setInputValue] = useState('');
  const handleChange = (event) => {
    setInputValue(event.target.value);
  }

  const handleSubmit = (event) => {
    event.preventDefault();
    // Send input (message) to database here
    // ....
    console.log(message);
  };

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      handleSubmit(event);
    }
  }

  return (
    <div className='ChatForm'>
      <form onSubmit={handleSubmit}>
      <input type='text' name="message" placeholder='Silakan masukkan pertanyaan Anda...' onChange={handleChange} onKeyDown={handleKeyDown}/>
      </form>
    </div>
  )
}

export default ChatForm