import React from 'react'
import { useState } from 'react'

function ChatForm({childToParent}) {

  const [message, setMessage] = useState('');

  const handleChange = (event) => {
    setMessage(event.target.value);
  }

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      setMessage(event.target.value);
      childToParent(message);
    }
  }

  return (
    <div className='ChatForm'>
        <input type='text' name="message" placeholder='Silakan masukan pesan Anda!' onChange={handleChange} onKeyDown={handleKeyDown}/>
    </div>
  )
}

export default ChatForm