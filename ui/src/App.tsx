import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';
import { StatusResponse } from '../../autogen/ts/example'

function App() {
  const [version, setVersion] = useState<any>({});
  const [status, setStatus] = useState<StatusResponse>({} as StatusResponse);

  useEffect(() => {
    const init = async () => {
      try {
        const result = await axios.get('/version', { headers: {'Accept': 'application/json'} });
        setVersion(result.data)
      } catch (e) {
        console.error(e)
      }
      try {
        const result = await axios.get('/api/v1/status', { headers: {'Accept': 'application/json'} });
        setStatus(result.data)
      } catch (e) {
        console.error(e)
      }
    };
    init();
  }, []);

  return (
    <div className="App">
      Hello World, if you want to see api documentation it should be somewhere <a href="/api/docs">here</a><br/>
      
      Version:
      <pre>
        {JSON.stringify(version, null, 2)}
      </pre>
      Status:
      <pre>
        {JSON.stringify(status, null, 2)}
      </pre>
    </div>
  );
}

export default App;
