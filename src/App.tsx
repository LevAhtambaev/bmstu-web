import React from 'react';
import {Routes, Route} from 'react-router-dom'
import {Info} from "./components/Info";
import {NotFound} from "./components/NotFound";
import {MangaDescription} from "./components/MangaDescription";
import {MainPage} from "./components/MainPage";
import {MangaPage} from "./components/MangaPage";
import {CartPage} from "./components/CartPage";
import {Registration} from "./components/RegisterPage";
import {LoginPage} from "./components/LoginPage";
import {ProfilePage} from "./components/ProfilePage";




export const ENDPOINT = "http://localhost:8080"

function App() {
    return (
        <div>
            <Routes>
                <Route path="/homepage" element={<MainPage/>}> </Route>
                <Route path="/manga" element={<MangaPage/>}></Route>
                <Route path="/manga/:id" element={<MangaDescription/>}></Route>
                <Route path="/cart"element={<CartPage/>}></Route>
                <Route path="/info" element={<Info/>}/>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/registration" element={<Registration/>}/>
                <Route path="/profile/:id" element={<ProfilePage/>}></Route>
                <Route path="/error" element={<NotFound/>}/>
            </Routes>
        </div>
    )
}

export default App;
