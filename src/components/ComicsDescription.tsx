import {useLocation, Link} from "react-router-dom";
import {Navbar} from "./Navbar";
import React from "react";

export function ComicsDescription() {
    return (
        <div>
            <Navbar/>
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/comics">Comics</Link> / {useLocation().state.comics.Name}
            </p>

            <img src={useLocation().state.comics.Image} width="23%" height="70%" className="mx-auto rounded-2xl" alt="comics"/>

            <p className=" font-bold text-4xl text-center">
                {useLocation().state.comics.Name}
            </p>

            <p className="mt-8 font-medium text-2xl text-center">
                {useLocation().state.comics.Price} рублей
                <p className="italic text-xl">
                    {useLocation().state.comics.Description}
                </p>
            </p>

            <p className="my-8 text-center">
                <Link to="/comics"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>
        </div>
    )
}