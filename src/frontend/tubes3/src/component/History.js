import React, {useEffect, useState} from 'react'
import HistoryTab from './HistoryTab'

function History({data, setData}) {
    const[history, setHistory] = useState([]);

    const HISTORY_API = "http://localhost:8080/history"

    function getHistory(url) {
        fetch(url, {
            method: 'POST',
        })
            .then(response => response.json())
            .then(json => handleHistory(json))
            .catch(error => console.error(error));
    }

    const handleHistory = (json) => {
        setHistory(json)
    }

    useEffect(() => {
        getHistory(HISTORY_API)
    }, []);

    return (
    <div className='history'>
        <div className='history-container'>
            {
                history && history.map((item) => {
                    console.log(item.SessionID)
                    return <HistoryTab sessionName={item.SessionName} sessionId={item.SessionID} data={data} setData={setData}/>
                })
            }
        </div>
    </div>
  )
}

export default History