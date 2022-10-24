import {IManga} from '../models'
import {useState} from "react";

interface MangaProps {
    manga: IManga
}

export function Manga(props: MangaProps) {
    const [showDetails, setShowDetails] = useState<boolean>(false)
    return (
        <div
        className="border w-1/2 py-0 px-0 rounded flex flex-col justify-between items-center mb-2 place-content-start"
        >
            <img src={process.env.PUBLIC_URL + props.manga.image} className="w-1/5" alt={props.manga.title}/>
            <p>{ props.manga.title }</p>
            <p className="font-bold">{props.manga.price} рублей</p>
            <button
                className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
                onClick={() => setShowDetails((prevState) => !prevState)}
            >
                {!showDetails ? <div>Show Details</div> : <div>Hide Details</div>}
            </button>

            {showDetails && <div>
                <p>{props.manga.description}</p>
            </div>}
        </div>
    )
}