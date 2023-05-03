import React from 'react'

function RadioAlgorithm() {
  return (
    <div className='RadioAlgorithm'>
        <hr style="color:gray;background-color:gray"></hr>
        Pilih Algoritma yang diinginkan
        <div className='algorithm-radio'>
            <div className='kmp-radio'>
                <input type="radio" value="KMP"/> KMP
            </div>
            <div className='bm-radio'>
                <input type="radio" value="BM" /> BM
            </div>
        </div>
    </div>
  )
}

export default RadioAlgorithm