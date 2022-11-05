import {IManga} from "../models";
import {Link} from "react-router-dom";


export interface MangaProps {
    manga: IManga
}

export function Manga(props: MangaProps) {
    return (
        <div
            className="py-3 px-0 rounded flex flex-col justify-self-center items-center mb-2 place-content-start"
        >
            <img src={`https://res.cloudinary.com/dsd9ne1xr/image/upload/${props.manga.Image}/${props.manga.UUID}.jpg`} className="w-1/5" alt={props.manga.Name}/>
            <p>{ props.manga.Name }</p>
            <p className="font-bold">{props.manga.Price} рублей</p>
            <Link to={`${props.manga.UUID}`}
                  className="w-1/10 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white py-1 px-3 place-self-center rounded-full text-xl font-bold"
                  state={{manga: props.manga}}
            >
                Подробнее
            </Link>
        </div>
    )
}