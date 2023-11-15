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
import Users from "./pages/Users";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route path="/albums" element={<Albums />} />
                <Route path="/users" element={<Users />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;