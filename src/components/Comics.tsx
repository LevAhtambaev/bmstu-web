import {IComics} from "../models";
import {Link} from "react-router-dom";
import {useContext, useState} from "react";
import {MyContext} from "./ComicsPage";
import {AddToCart} from "../requests/AddToCart";
import {getRole, getToken} from "../modules";
import {DeleteComics} from "../requests/DeleteComics";


export function Comics() {
    const ctx = useContext(MyContext)
    let access_token = getToken()
    let showAddCartButton = true
    if (access_token == "") {
        showAddCartButton = false
    }
    const [roles, setRole] = useState()
    const role = getRole(access_token)
    role.then((result) => {
        setRole(result)
    })
    let showManagerButton = false
    if (roles === 1) {
        showManagerButton = true
    }
    return (
        <div
            className="py-3 px-0 rounded flex flex-col justify-self-center items-center mb-2 place-content-start"
        >
            <img src={ctx.Image} className="w-3/5 sm:w-2/5" alt={ctx.Name}/>
            <p className="text-2xl">{ ctx.Name }</p>
            <p className="text-xl font-bold">{ctx.Price} рублей</p>
            <Link to={`${ctx.UUID}`}
                  className="w-30 sm:w-1/10 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white py-1 sm:px-3 place-self-center rounded-full text-xl font-bold"
                  state={{comics: ctx}}
            >Подробнее
            </Link>
            {showAddCartButton && <p className="mt-2 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white sm:px-3 place-content-between rounded-full text-xl font-bold">
                {AddToCart(ctx.UUID)}
            </p>}
            {showManagerButton && <p className="mt-2 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white sm:px-3 place-content-between rounded-full text-xl font-bold">
                {DeleteComics(ctx.UUID)}
            </p>}
            {showManagerButton && <Link to="/change" className="mt-2 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white sm:px-3 place-content-between rounded-full text-xl font-bold"
                                        state={{UUID: ctx.UUID, Name: ctx.Name, Rate: ctx.Rate, Year: ctx.Year, Genre: ctx.Genre, Price: ctx.Price, Episodes: ctx.Episodes, Description: ctx.Description, Image: ctx.Image}}
            >
                Изменить
            </Link>}
        </div>
    )
}