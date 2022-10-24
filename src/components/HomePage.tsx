import {Mangas} from "../repository/Mangas";
import {Manga} from "./Manga";
import React from "react";

export function HomePage() {
    return (
        <div className="container mx-auto max-w-5xl pt-5 flex justify-between">
            {Mangas.map((manga, key) => {
                return <Manga manga={manga} key={key}/>
            })}
        </div>
    )
}