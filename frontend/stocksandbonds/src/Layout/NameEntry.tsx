import React from 'react';
import TextField from '@material-ui/core/TextField';

export default function NameEntry(nameChange:Function) {

  var name = ""

  return (

      <TextField placeholder="Your Name" id="outlined-basic" label="New Name" variant="outlined" 
      onChange={(e) => {
        name = e.target.value
      }}
      onKeyPress={(e) => {
          if (e.key === 'Enter') {
          nameChange(name)
          }
      }}      />

  );
}