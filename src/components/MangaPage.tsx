import {Manga} from "./Manga";
import {Banner} from "./Banner";
import {getJsonMangas} from "../modules";
import {useEffect, useState, useReducer} from "react";
import {createContext} from "react";
import {IManga} from "../models";
import {reducer, success} from "../reducer";
import {Filter} from "./Filter";
import {GetMangas} from "../requests/GetMangas";
import {manga_context} from "../context/context";


export const MyContext = createContext(manga_context);
const initialState = {mangas: []}

export function MangaPage() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const [showFilter, setShowFilter] = useState(false)


    const url = `manga`

    useEffect(() => {
        getJsonMangas(url).then((result) => {
            dispatch({type: success, mangas: result})
        })
    }, [url])

    return (
        <>
            <Banner/>
            <form className="form--manga">
                <label>
                    <input className="shadow ml-2 appearance-none border rounded w-1/8 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="search" type="text" placeholder="Search"/>
                    <button className="ml-2">Find</button>
                </label>
            </form>
            <button className="bg-white ml-1 hover:bg-gray-100 text-gray-800 ml-72 py-1.5 px-2 border rounded shadow" onClick={()=>{
                setShowFilter(!showFilter)
            }}>Filter</button>
            {showFilter && <Filter/>}
            <div className="container grid grid-cols-3 gap-2">
                {GetMangas().map((manga: IManga)=> {
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