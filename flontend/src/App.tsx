import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import TopPage from "./views/pages/TopPage";

function App() {
  return (
    <React.Fragment>
      <BrowserRouter>
        <Routes>
          <Route path="/hello" element={<h1>{"hello world"}</h1>} />
          <Route path="/" element={<TopPage />} />
        </Routes>
      </BrowserRouter>
    </React.Fragment>
  );
}

export default App;
