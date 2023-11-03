import React from 'react';
import {
    BrowserRouter,
    Routes,
    Route,
} from "react-router-dom";
import './App.css';
import Albums from './pages/Albums';
import Login from './pages/Login';
import Register from "./pages/Register";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route path="/albums" element={<Albums />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;