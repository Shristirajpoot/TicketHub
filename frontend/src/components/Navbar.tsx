import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { removeToken } from "../utils/auth";

const Navbar: React.FC = () => {
    const navigate = useNavigate();

    const handleLogout = () => {
        removeToken();
        navigate("/login");
    };

    return (
        <nav>
            <Link to="/">Dashboard</Link>
            <Link to="/create-event">Create Event</Link>
            <button onClick={handleLogout}>Logout</button>
        </nav>
    );
};

export default Navbar;
