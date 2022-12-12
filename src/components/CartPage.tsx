import {ICart} from "../models";
import {Cart} from "./Cart";
import {GetCart} from "../requests/GetCart";
import React, {createContext} from "react";
import {cart_context} from "../context/context";
import {Navbar} from "./Navbar";
import {AddOrder} from "../requests/AddOrder";


export const MyContext = createContext(cart_context);

export function CartPage() {
    let cart = GetCart()
    let showCart = true
    if (cart.length === 0) {
        showCart = false
    }
    let cars_uuid: string[] = new Array()
    cart.map((cart: ICart) => {
        cars_uuid.push(cart.Manga)
    })
    return (
        <>
            <Navbar/>
        <div className="bg-gray-100 min-h-screen">
            <div className=" flex flex-col gap-4 container">
                <p className="ml-4 text-2xl font-normal text-black">
                    Cars
                </p>
                {showCart && cart.map((cart: ICart) => {
                    return (
                        <MyContext.Provider value={cart}>
                            <Cart/>
                        </MyContext.Provider>
                    )
                })}
                {!showCart && <h1>Ваша корзина пуста!</h1>}
            </div>
            {showCart &&
                <form>
                    <p className="text-center">
                        {AddOrder(cars_uuid)}
                    </p>
                </form>}
        </div>
            </>
    )
}