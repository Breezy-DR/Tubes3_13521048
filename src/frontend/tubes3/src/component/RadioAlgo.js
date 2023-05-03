import React from 'react'

function RadioAlgo() {
  return (
    <div className='RadioAlgo'>
        <hr></hr>
        <div className='buttons'>
            <div className='KMP-Button'>
                <input type="radio" value="KMP"/> KMP
            </div>
            <div className='BM-Button'>
                <input type="radio" value="BM" /> BM
            </div>
        </div>
    </div>
  )
}

export default RadioAlgo