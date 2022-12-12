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
import {AddManga} from "./components/AddManga"
import {OrderPage} from "./components/OrderPage";
import {ChangeManga} from "./components/ChangeManga";




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
                <Route path="/add" element={<AddManga/>}/>
                <Route path="/profile" element={<ProfilePage/>}></Route>
                <Route path="/orders" element={<OrderPage/>}></Route>
                <Route path="/error" element={<NotFound/>}/>
                <Route path="/change" element={<ChangeManga/>}/>
            </Routes>
        </div>
    )
}

export default App;
