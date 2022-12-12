import React from 'react';
import {Routes, Route} from 'react-router-dom'
import {Info} from "./components/Info";
import {NotFound} from "./components/NotFound";
import {ComicsDescription} from "./components/ComicsDescription";
import {MainPage} from "./components/MainPage";
import {ComicsPage} from "./components/ComicsPage";
import {CartPage} from "./components/CartPage";
import {Registration} from "./components/RegisterPage";
import {LoginPage} from "./components/LoginPage";
import {ProfilePage} from "./components/ProfilePage";
import {AddComics} from "./components/AddComics"
import {OrderPage} from "./components/OrderPage";
import {ChangeComics} from "./components/ChangeComics";




export const ENDPOINT = "http://localhost:8080"

function App() {
    return (
        <div>
            <Routes>
                <Route path="/homepage" element={<MainPage/>}> </Route>
                <Route path="/comics" element={<ComicsPage/>}></Route>
                <Route path="/comics/:id" element={<ComicsDescription/>}></Route>
                <Route path="/cart"element={<CartPage/>}></Route>
                <Route path="/info" element={<Info/>}/>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/registration" element={<Registration/>}/>
                <Route path="/add" element={<AddComics/>}/>
                <Route path="/profile" element={<ProfilePage/>}></Route>
                <Route path="/orders" element={<OrderPage/>}></Route>
                <Route path="/error" element={<NotFound/>}/>
                <Route path="/change" element={<ChangeComics/>}/>
            </Routes>
        </div>
    )
}

export default App;
