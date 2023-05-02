import React from 'react';
import { useState } from 'react';
import ChatForm from './ChatForm';

function Chat() {
  
  const[data, setData] = useState('');

  const API = 'https://my-json-server.typicode.com/Breezy-DR/chatTesting/Chats'

  const parentToChild = () => {
    setData()
  }

  const childToParent = (childData) => {
    setData(childData);
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

  return (
    <div className='Chat'>
        {/* <Conversation parentToChild={data}/> */}
    <div className='conversation'>
      {
        data && data.map((item) => {
          if(item.from === "user"){
            // return <MessageBubbleRight dataToMessage={item.chat}/>
            <div className='usermessage'>
                <p>{dataToMessage}</p>
            </div>
          }else{
            // return <MessageBubbleLeft dataToMessage={item.chat}/>
            <div className='messagebubble owner'>
                <p>{dataToMessage}</p>
            </div>
          }
        })
      }
        {/* <div className='flex-container'></div> */}
        <div className='message-layout'></div>
    </div>
        <ChatForm childToParent={childToParent}/>
    </div>
  )
}

export default Chat
