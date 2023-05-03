import React from 'react'
import ChatForm from './ChatForm';
import GPTMessage from './GPTMessage';
import UserMessage from './UserMessage';
import { useEffect, useState } from 'react';

function Chat() {
  
  const[data, setData] = useState('');

  const API = 'https://my-json-server.typicode.com/Breezy-DR/chatTesting/Chats'

  const parentToChild = () => {
    setData()
  }

  const childToParent = (childData) => {
    setData(childData);
    //alert(childData);
  }

  const fetchChats = () => {
    fetch(API).then((res) => res.json()).then((res) => {
      console.log(res);
      setData(res);
    })
  }

  useEffect(() => {
    fetchChats()
  }, []);

  // const dataToMessage = () => {
  //   setData({parentToChild});
  // }

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
        <ChatForm childToParent={childToParent}/>
    </div>
  )
}

export default Chat