import React, { useState } from 'react';
import { greeterClient as GreeterClient } from './generated/HelloServiceClientPb.ts';
import { HelloRequest } from './generated/hello_pb';

const client = new GreeterClient('http://localhost:10000', null, null);

function App() {
  const [response, setResponse] = useState('');

  const callSayHello = () => {
    const request = new HelloRequest();
    request.setName('World');

    client.sayHello(request, {}, (err, response) => {
      if (err) {
        console.error(err);
      } else {
        setResponse(response.getMessage());
      }
    });
  };

  return (
    <div className="App">
      <button onClick={callSayHello}>Say Hello</button>
      <div>{response}</div>
    </div>
  );
}

export default App;
