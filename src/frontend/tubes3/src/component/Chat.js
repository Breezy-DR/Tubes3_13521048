import React from 'react'
import ChatForm from './ChatForm';
import GPTMessage from './GPTMessage';
import UserMessage from './UserMessage';
import { useEffect, useState } from 'react';

function Chat() {
  
  const[data, setData] = useState([]);

  const API = 'https://my-json-server.typicode.com/Breezy-DR/chatTesting/Chats'

  const sendData = async () => {
    const url = "ec2-54-169-32-134.ap-southeast-1.compute.amazonaws.com:8080/question/";
    const data = {
      "question": "ping? namx 5 nama ikan? 2*9^(7-2); 7 nama ikan; sebutkan 7 nama ikan; 2023/11/12?pong",
      "search_algorithm": "bm",
      "session_id": "7bbf04e082"
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
        <ChatForm/>
    </div>
  )
}

export default Chat