import { useState } from "react";

const Counter = () => {
  const [counter, setCounter] = useState(0);
  return (
    <div>
      {/* <header>hola, {props.name}</header> */}
      <h2>{counter}</h2>
      <button onClick={() => setCounter(counter - 1)}>Decrement</button>
      <button onClick={() => setCounter(counter + 1)}>Increment</button>
    </div>
  );
};

export default Counter;
