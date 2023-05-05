import React from 'react'
import RadioAlgo from './RadioAlgo'
import History from './History'

function Sidebar({data, setData, selectedAlg, setSelectedAlg}) {
  return (
    <div className='sidebar'>
        <History data={data} setData={setData}/>
        <div className='flex-container'></div>
        <RadioAlgo selectedAlg={selectedAlg} setSelectedAlg={setSelectedAlg}/>
    </div>
  )
}

export default Sidebar