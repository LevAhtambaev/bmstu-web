import {IManga} from "../models";
import {Link} from "react-router-dom";
import {useContext} from "react";
import {MyContext} from "./CartPage";
import {GetManga} from "../requests/GetManga";
import {DeleteFromCart} from "../requests/DeleteFromCart";


export function Cart() {
    const ctx = useContext(MyContext)
    let Manga = GetManga(ctx.Manga)
    return (
        <div
            className="py-3 px-0 rounded flex flex-col justify-self-center items-center mb-2 place-content-start"
        >
        <img src={`https://res.cloudinary.com/dsd9ne1xr/image/upload/${Manga.Image}/${Manga.UUID}.jpg`} className="w-1/5" alt={Manga.Name}/>
    <p>{ Manga.Name }</p>
    <p className="font-bold">{Manga.Price} рублей</p>
    <Link to={`manga/${Manga.UUID}`}
    className="w-1/10 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white py-1 px-3 place-self-center rounded-full text-xl font-bold"
    state={{manga: Manga}}
>
    Подробнее
    </Link>

            <p className="place-self-center col-span-3 rounded-full bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded">
                {DeleteFromCart(Manga.UUID)}
            </p>


    </div>
)
}