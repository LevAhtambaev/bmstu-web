import {Manga} from "./Manga";
import {Banner} from "./Banner";
import {getJson} from "../modules";
import {useEffect, useState} from "react";
import {IManga} from "../models";



export function MangaPage() {
    const [mangas, setManga] = useState<IManga[]>([])

    const getAllManga = async () => {
        const result = await getJson("manga/")
        await setManga(result)
    }

    useEffect(() => {getAllManga()}, [])
    return (
        <>
            <Banner/>
            <div className="container grid grid-cols-3 gap-2">
                {mangas.map((manga, key)=> {
                    return <Manga manga={manga} key={key}/>
                })}
            </div>
        </>
    )
}