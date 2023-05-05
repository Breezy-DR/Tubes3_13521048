import React, {createRef} from 'react'
import GPTMessage from './GPTMessage';
import UserMessage from './UserMessage';

function Chat({data, setData, selectedAlg}) {

  const QUESTION_API = "http://localhost:8080/question"

  const chatForm = createRef();
  const chatContainer = createRef();

  const handleSubmit = (event) => {
    event.preventDefault();
    let input = event.target.value
    chatForm.current.value = ''
    data.push({from:'user', chat:input})

    if (!sessionStorage.getItem("session_id")) {
      sessionStorage.setItem("session_id", "")
    }

    ask(QUESTION_API,
        {"question": input, "search_algorithm":selectedAlg, "session_id": sessionStorage.getItem("session_id")},
        x)
  };

  function x(json) {
    sessionStorage.setItem("session_id", json["session_id"]);
    let responseArr = json['response']
    let newData = [...data]
    responseArr.forEach(item => {
      newData.push({from:'server', chat:item})
    })
    setData(newData)
  }

  function ask(url, data, callback) {
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
        <div ref={chatContainer} className='conversation' style={{maxHeight:'80%'}}>
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
            <input ref={chatForm} type='text' name="message" placeholder='Masukkan pertanyaan Anda (press enter to submit)' onKeyDown={handleKeyDown}/>
          </form>
        </div>
      </div>
  )
}

export default Chat