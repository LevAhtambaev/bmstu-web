import React from "react"
import {Link} from "react-router-dom";



export function Navbar() {
    return (
        <nav className="relative shadow-xl flex flex-wrap items-center px-2 py-3 bg-gray-600 mb-3">
            <p className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 uppercase text-white">
                <Link to="/homepage">Home Page</Link>
            </p>

            <p className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 uppercase text-white">
                <Link to="/manga">All Manga</Link>
            </p>

            <p className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 uppercase text-white">
                <Link to="/info">About Us</Link>
            </p>
        </nav>
    );
}