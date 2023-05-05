import React from 'react'
import SelectedValueContext from './SelectedValueContext';

function RadioAlgo({selectedAlg, setSelectedAlg}) {
  const handleRadioChange = (event) => {
      setSelectedAlg(event.target.value);
  };

  return (
    <SelectedValueContext.Provider value={selectedAlg}>
      <div className='RadioAlgo'>
          <div className='buttons'>
              <div className='KMP-Button'>
                  <input type="radio" value="kmp" checked={selectedAlg === 'kmp'} onChange={handleRadioChange}/> KMP
              </div>
              <div className='BM-Button'>
                  <input type="radio" value="bm" checked={selectedAlg === 'bm'} onChange={handleRadioChange}/> BM
              </div>
          </div>
      </div>
    </SelectedValueContext.Provider>
  )
}

export default RadioAlgo