import {Comics} from "./Comics";
import {getJsonAllComics} from "../modules";
import React, {useEffect, useState, useReducer} from "react";
import {createContext} from "react";
import {IComics} from "../models";
import {reducer, success} from "../reducer";
import {GetAllComics} from "../requests/GetAllComics";
import {comics_context} from "../context/context";
import Box from "@mui/material/Box";
import Slider from "@mui/material/Slider";
import {Navbar} from "./Navbar";
import {Link, useLocation} from "react-router-dom";


export const MyContext = createContext(comics_context);

export function ComicsPage() {
    const comics = GetAllComics()
    console.log(comics)
    const [name, setName] = useState('')

    const filteredComics= comics.filter((comics: { Name: string }) => {
        return comics.Name.toLowerCase().includes(name.toLowerCase())
    })

    const [price, setPrice] = useState<number[]>([0,5000]);

    const minDistance = 100;

    const handleChange = (event: Event, newValue: number | number[], activeThumb: number) => {
        if (!Array.isArray(newValue)) {
            return;
        }

        if (activeThumb === 0) {
            setPrice([Math.min(newValue[0], price[1] - minDistance), price[1]]);
        } else {
            setPrice([price[0], Math.max(newValue[1], price[0] + minDistance)]);
        }
    };

    const marks = [
        {
            value: 0,
            label: '0 Р',
        },
        {
            value: 1000,
            label: '1000 Р',
        },
        {
            value: 2000,
            label: '2000 Р',
        },
        {
            value: 3000,
            label: '3000 Р',
        },
        {
            value: 4000,
            label: '4000 Р',
        },
        {
            value: 5000,
            label: '5000 Р',
        },
    ];

    function valuetext(price: number) {
        return `${price} Р`;
    }

    return (
        <>
            <Navbar/>
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/comics">Comics</Link>
            </p>
            <div className="flex pt-5 sm:pt-5 place-content-center">
                <form>
                    <input
                        type="text"
                        className="block w-full px-4 py-2 text-gray-500 text-2xl bg-white border rounded-full focus:border-gray-400 focus:ring-gray-400 focus:outline-none focus:ring focus:ring-opacity-40"
                        placeholder="Поиск..."
                        onChange={(event) => setName(event.target.value)}
                    />
                </form>
            </div>
            <div className="flex pt-5 sm:pt-5 place-content-center">
                <Box sx={{ width: 300 }}>
                    <Slider
                        getAriaLabel={() => 'Price filter'}
                        valueLabelDisplay="auto"
                        getAriaValueText={valuetext}
                        value={price}
                        marks={marks}
                        onChange={handleChange}
                        disableSwap
                        step={100}
                        min={0}
                        max={5000}
                    />
                </Box>
            </div>
            <div className="container mt-6 grid grid-cols-2 sm:grid-cols-3 gap-2 mx-auto">
                {filteredComics.filter((comics: { Price: number; }) => comics.Price >= price[0] && comics.Price <= price[1]).map((comics: IComics, key: any)=> {
                    return (
                        <MyContext.Provider value={comics} key={key}>
                            <Comics/>
                        </MyContext.Provider>
                    )
                })}
            </div>
        </>
    )
}