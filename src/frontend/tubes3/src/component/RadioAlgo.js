import React from 'react'
import { useState } from 'react';
import SelectedValueContext from './SelectedValueContext';

function RadioAlgo() {
  const [selectedValue, setSelectedValue] = useState('');

  const handleRadioChange = (event) => {
    setSelectedValue(event.target.value);
  };

  return (
    <SelectedValueContext.Provider value={selectedValue}>
      <div className='RadioAlgo'>
          <div className='buttons'>
              <div className='KMP-Button'>
                  <input type="radio" value="KMP" checked={selectedValue === 'KMP'} onChange={handleRadioChange}/> KMP
              </div>
              <div className='BM-Button'>
                  <input type="radio" value="BM" checked={selectedValue === 'BM'} onChange={handleRadioChange}/> BM
              </div>
          </div>
      </div>
    </SelectedValueContext.Provider>
  )
}

export default RadioAlgo