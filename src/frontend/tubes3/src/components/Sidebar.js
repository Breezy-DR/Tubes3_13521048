import React from 'react'
import RadioAlgorithm from './RadioAlgorithm'
import History from './History'

function Sidebar() {
    return (
      <div className='sidebar'>
          <History/>
          <div className='message-layout'></div>
          <RadioAlgorithm/>
      </div>
    )
  }
  
  export default Sidebar