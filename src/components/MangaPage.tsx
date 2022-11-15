import {Manga} from "./Manga";
import {Banner} from "./Banner";
import {getJson} from "../modules";
import {useEffect, useState, useReducer} from "react";
import {createContext} from "react";
import {IManga} from "../models";
import reducer from "../reducer";

let dv : IManga = {
    UUID: "",
    Image: "",
    Year: 0,
    Name: "",
    Genre: "",
    Description: "",
    Episodes: 0,
    Rate: 1,
    Price: 0,
}

export const MyContext = createContext(dv);

export function MangaPage() {
    const [state, dispatch] = useReducer(reducer, [] as IManga[])

    const getAllManga = async () => {
        const result = await getJson("manga/")
        await dispatch(result)
    }

    useEffect(() => {getAllManga()}, [])
    return (
        <>
            <Banner/>
            <div className="container grid grid-cols-3 gap-2">
                {state.map((manga)=> {
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