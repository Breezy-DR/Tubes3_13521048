import React from 'react'
import RadioAlgo from './RadioAlgo'
import History from './History'

function Sidebar() {
  return (
    <div className='sidebar'>
        <History/>
        <div className='flex-container'></div>
        <RadioAlgo/>
    </div>
  )
}

export default Sidebar