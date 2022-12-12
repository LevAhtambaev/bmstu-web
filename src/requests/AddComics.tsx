import {addComics} from "../modules";
import React from "react";


export function AddingComics(name: string, rate: number, year: number, genre: string, price: number, episodes: number, description: string, image: string) {

    const url = `comics`

    function Add() {
        addComics(url, name, rate, year, genre, price, episodes, description, image)
    }


    return (
        <>
            <button
                onClick={() => Add()}
    className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
        >
        Добавить
        </button>
        </>
);

}