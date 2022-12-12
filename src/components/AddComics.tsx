import React, {useState} from "react"
import {Navbar} from "./Navbar";
import {AddingComics} from "../requests/AddComics";
import {Link} from "react-router-dom";

export function AddComics() {
    const [name, setName] = useState('Batman');
    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);
    };

    const [rate, setRate] = useState(7);
    const handleChangeRate = (event: { target: { value: any; }; }) => {
        setRate(Number(event.target.value));
    };

    const [year, setYear] = useState(1970);
    const handleChangeYear = (event: { target: { value: any; }; }) => {
        setYear(Number(event.target.value));
    };

    const [genre, setGenre] = useState('Детектив');
    const handleChangeGenre = (event: { target: { value: any; }; }) => {
        setGenre(event.target.value);
    };

    const [price, setPrice] = useState(2400);
    const handleChangePrice = (event: { target: { value: any; }; }) => {
        setPrice(Number(event.target.value));
    };

    const [episodes, setEpisodes] = useState(5);
    const handleChangeEpisodes = (event: { target: { value: any; }; }) => {
        setEpisodes(Number(event.target.value));
    };

    const [description, setDescription] = useState('Batman comics');
    const handleChangeDescription = (event: { target: { value: any; }; }) => {
        setDescription(event.target.value);
    };

    const [image, setImage] = useState('https://res.cloudinary.com/dl0tawm7w/image/upload/v1670865178/comics/91fCA7sK8mL_h2rpcq.jpg');
    const handleChangeImage = (event: { target: { value: any; }; }) => {
        setImage(event.target.value);
    };

    return(
        <>
            <Navbar/>
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/comics">Comics</Link> / AddComics
            </p>
            <div className="mt-10 sm:mt-0">
                <div className="md:gap-6">
                    <div className="px-4 ">
                        <h3 className="text-3xl mt-2 text-center font-medium leading-6 text-gray-900">Добавить комикс</h3>
                    </div>
                    <div className="mt-5  md:mt-0">
                        <form>
                            <div className="overflow-hidden shadow sm:rounded-md">
                                <div className="bg-white px-4 py-5 sm:p-6">
                                    <div className="grid grid-cols-5 grid-rows-4 gap-6">
                                        <div className="col-span-2">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Название
                                            </label>
                                            <input
                                                type="text"
                                                onChange={handleChangeName}
                                                value={name}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Рейтинг
                                            </label>
                                            <input
                                                type="number"
                                                onChange={handleChangeRate}
                                                value={rate}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Год
                                            </label>
                                            <input
                                                type="number"
                                                max="2022"
                                                onChange={handleChangeYear}
                                                value={year}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Количество эпизодов
                                            </label>
                                            <input
                                                type="number"
                                                onChange={handleChangeEpisodes}
                                                value={episodes}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Жанр
                                            </label>
                                            <select
                                                onChange={handleChangeGenre}
                                                value={genre}
                                                className="mt-1 block w-full rounded-md border border-gray-300 bg-white py-2 px-3 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-base"
                                            >
                                                <option>Детектив</option>
                                                <option>Ужасы</option>
                                                <option>Романтика</option>
                                                <option>Приключения</option>
                                                <option>Научная фантастика</option>
                                            </select>
                                        </div>

                                        <div className="">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Цена
                                            </label>
                                            <input
                                                type="number"
                                                onChange={handleChangePrice}
                                                value={price}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="col-span-2">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Ссылка на изображение
                                            </label>
                                            <input
                                                type="text"
                                                onChange={handleChangeImage}
                                                value={image}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>

                                        <div className="col-span-5">
                                            <label htmlFor="first-name" className="block text-base font-medium text-gray-700">
                                                Описание
                                            </label>
                                            <input
                                                type="text"
                                                onChange={handleChangeDescription}
                                                value={description}
                                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                                            />
                                        </div>




                                    </div>
                                </div>
                                <div className="bg-gray-50 px-4 py-3 text-center sm:px-6">
                                    {AddingComics(name, rate, year, genre, price, episodes, description, image)}
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </>
    )
}