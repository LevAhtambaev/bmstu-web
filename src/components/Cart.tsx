import {IComics} from "../models";
import {Link} from "react-router-dom";
import {useContext} from "react";
import {MyContext} from "./CartPage";
import {GetComics} from "../requests/GetComics";
import {DeleteFromCart} from "../requests/DeleteFromCart";


export function Cart() {
    const ctx = useContext(MyContext)
    let Comics = GetComics(ctx.Comics)
    return (
        <div
            className="py-3 px-0 rounded flex flex-col justify-self-center items-center mb-2 place-content-start"
        >
        <img src={Comics.Image} className="w-1/2 sm:w-1/5" alt={Comics.Name}/>
    <p>{ Comics.Name }</p>
    <p className="font-bold">{Comics.Price} рублей</p>
    <Link to={`Comics/${Comics.UUID}`}
    className="border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white py-1 sm:px-3 place-self-center rounded-full text-xl font-bold"
    state={{Comics: Comics}}
>Подробнее
    </Link>

            <p className="mt-2 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white sm:px-3 place-self-auto rounded-full text-xl font-bold">
                {DeleteFromCart(Comics.UUID)}
            </p>


    </div>
)
}