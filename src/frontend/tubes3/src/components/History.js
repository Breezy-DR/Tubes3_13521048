import React from 'react'
import io from 'socket.io-client'
const socket = io.connect("https://localhost:3001");

function History() {
  return (
    <div className='History'>
        <div className='title'>Riwayat</div>
        <div className='content'>
        <div className='HistoryTab'>
            
        </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            <div className='HistoryTab'>
        
    </div>
            
        </div>
    </div>
  )
}

export default History