import React, { useEffect, useState } from "react";
import axios from "axios";

const World = () => {
  const path = "/api";
  const [message, setMessage] = useState("");
  useEffect(() => {
    const setResponse = async () => {
      const response = await axios.get(path);
      setMessage(response.data.hello);
      console.log(response.data.hello);
    };
    setResponse();
  }, []);
  return (
    <React.Fragment>
      <h1>{"Hello: "}</h1>
      <h1>{message}</h1>
    </React.Fragment>
  );
};

export default World;
