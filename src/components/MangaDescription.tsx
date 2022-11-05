import {useLocation, Link} from "react-router-dom";

export function MangaDescription() {
    return (
        <div>
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/manga">Mangas</Link>  / {useLocation().state.manga.Name}
            </p>

            <img src={`https://res.cloudinary.com/dsd9ne1xr/image/upload/${useLocation().state.manga.Image}/${useLocation().state.manga.UUID}.jpg`} width="23%" height="70%" className="mx-auto rounded-2xl" alt="manga"/>

            <p className=" font-bold text-4xl text-center">
                {useLocation().state.manga.Name}
            </p>

            <p className="mt-8 font-medium text-2xl text-center">
                {useLocation().state.manga.Price} рублей
                <p className="italic text-xl">
                    {useLocation().state.manga.Description}
                </p>
            </p>

            <p className="my-8 text-center">
                <Link to="/manga"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>
        </div>
    )
}