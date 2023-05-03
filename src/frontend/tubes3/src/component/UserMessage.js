import React from 'react'

function UserMessage({dataToMessage}) {
    return (
      <div className='UserMessage'>
          <p>{dataToMessage}</p>
      </div>
    )
}


export default UserMessage