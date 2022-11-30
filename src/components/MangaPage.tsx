import {Manga} from "./Manga";
import {Banner} from "./Banner";
import {getJsonMangas} from "../modules";
import React, {useEffect, useState, useReducer} from "react";
import {createContext} from "react";
import {IManga} from "../models";
import {reducer, success} from "../reducer";
import {Filter} from "./Filter";
import {GetMangas} from "../requests/GetMangas";
import {manga_context} from "../context/context";
import Box from "@mui/material/Box";
import Slider from "@mui/material/Slider";
import {Navbar} from "./Navbar";


export const MyContext = createContext(manga_context);
const initialState = {mangas: []}

export function MangaPage() {
    const [state, dispatch] = useReducer(reducer, initialState)


    const url = `manga`

    useEffect(() => {
        getJsonMangas(url).then((result) => {
            dispatch({type: success, mangas: result})
        })
    }, [url])

    const mangas = GetMangas()

    const [name, setName] = useState('')

    const filteredMangas= mangas.filter((manga: { Name: string }) => {
        return manga.Name.toLowerCase().includes(name.toLowerCase())
    })

    const [price, setPrice] = useState<number[]>([0,2000]);

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
            value: 500,
            label: '500 Р',
        },
        {
            value: 1000,
            label: '1000 Р',
        },
        {
            value: 20000,
            label: '2000 Р',
        },
    ];

    function valuetext(price: number) {
        return `${price} Р`;
    }

    return (
        <>
            <Navbar/>
            <Banner/>
            <div className="flex pt-5 place-content-center">
                <form>
                    <input
                        type="text"
                        className="block w-full px-4 py-2 text-gray-500 text-2xl bg-white border rounded-full focus:border-gray-400 focus:ring-gray-400 focus:outline-none focus:ring focus:ring-opacity-40"
                        placeholder="Поиск..."
                        onChange={(event) => setName(event.target.value)}
                    />
                </form>
            </div>
            <div className="flex pt-5 place-content-center">
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
                        max={2000}
                    />
                </Box>
            </div>
            <div className="container grid grid-cols-3 gap-2">
                {filteredMangas.filter((manga: { Price: number; }) => manga.Price >= price[0] && manga.Price <= price[1]).map((manga: IManga)=> {
                    return (
                        <MyContext.Provider value={manga}>
                            <Manga/>
                        </MyContext.Provider>
                    )
                })}
            </div>
        </>
    )
}