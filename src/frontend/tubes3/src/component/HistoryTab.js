import React from 'react'

function HistoryTab({sessionName, sessionId, data, setData }) {

  const HIST_SESSION_ID = "http://localhost:8080/history"

  function getSHistory() {
    let x = "/"+sessionId
    fetch(HIST_SESSION_ID + x, {
      method: 'POST'
    })
        .then(response => response.json())
        .then(json => populateFromHist(json))
        .catch(error => console.error(error));
  }

  function populateFromHist(json) {
    let newData = [...data]
    json.forEach(item => {
      newData.push({from:'user', chat:item.UserEntry})
      newData.push({from:'server', chat:item.Answer})
    })
    setData(newData)
  }

  return (
    <div className='history-tab' onClick={getSHistory}>
      {sessionName}
    </div>
  )
}

export default HistoryTab