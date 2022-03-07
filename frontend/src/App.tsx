import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import TopPage from "./views/pages/TopPage";
import World from "./views/pages/World";

function App() {
  return (
    <React.Fragment>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<TopPage />} />
          <Route path="/hello" element={<h1>{"hello world"}</h1>} />
          <Route path="/world" element={<World />} />
        </Routes>
      </BrowserRouter>
    </React.Fragment>
  );
}

export default App;
