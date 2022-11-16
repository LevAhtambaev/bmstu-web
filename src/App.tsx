import React from 'react';
import {Routes, Route} from 'react-router-dom'
import {Navbar} from "./components/Navbar";
import {Info} from "./components/Info";
import {NotFound} from "./components/NotFound";
import {MangaDescription} from "./components/MangaDescription";
import {MainPage} from "./components/MainPage";
import {MangaPage} from "./components/MangaPage";
import {CartPage} from "./components/CartPage";


export const ENDPOINT = "http://localhost:8080"

function App() {
    return (
        <div>
            <Navbar/>
            <Routes>
                <Route path="/homepage" element={<MainPage/>}> </Route>
                <Route path="/manga" element={<MangaPage/>}></Route>
                <Route path="/manga/:id" element={<MangaDescription/>}></Route>
                <Route path="/cart"element={<CartPage/>}></Route>
                <Route path="/info" element={<Info/>}/>
                <Route path="*" element={<NotFound/>}></Route>
                <Route path="/cart" element={<CartPage/>}/>
            </Routes>
        </div>
    )
}

export default App;
