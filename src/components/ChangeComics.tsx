import React, {useState} from "react"
import {Navbar} from "./Navbar";
import {ChangingComics} from "../requests/ChangeComics";
import {Link, useLocation} from "react-router-dom";

export function ChangeComics() {
    const [name, setName] = useState(useLocation().state.Name);
    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);
    };

    const [rate, setRate] = useState(useLocation().state.Rate);
    const handleChangeRate = (event: { target: { value: any; }; }) => {
        setRate(Number(event.target.value));
    };

    const [year, setYear] = useState(useLocation().state.Year);
    const handleChangeYear = (event: { target: { value: any; }; }) => {
        setYear(Number(event.target.value));
    };

    const [genre, setGenre] = useState(useLocation().state.Genre);
    const handleChangeGenre = (event: { target: { value: any; }; }) => {
        setGenre(event.target.value);
    };

    const [price, setPrice] = useState(useLocation().state.Price);
    const handleChangePrice = (event: { target: { value: any; }; }) => {
        setPrice(Number(event.target.value));
    };

    const [episodes, setEpisodes] = useState(useLocation().state.Episodes);
    const handleChangeEpisodes = (event: { target: { value: any; }; }) => {
        setEpisodes(Number(event.target.value));
    };

    const [description, setDescription] = useState(useLocation().state.Description);
    const handleChangeDescription = (event: { target: { value: any; }; }) => {
        setDescription(event.target.value);
    };

    const [image, setImage] = useState(useLocation().state.Image);
    const handleChangeImage = (event: { target: { value: any; }; }) => {
        setImage(event.target.value);
    };

    return(
        <>
            <Navbar/>
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/comics">Comics</Link> / ChangeComics
            </p>
            <div className="mt-10 sm:mt-0">
                <div className="md:gap-6">
                    <div className="px-4 ">
                        <h3 className="text-3xl mt-2 text-center font-medium leading-6 text-gray-900">Измененить комикс</h3>
                    </div>
                    <div className="mt-5  md:mt-0">
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
                                                <option>Сэнен</option>
                                                <option>Приключение</option>
                                                <option>Боевик</option>
                                                <option>Боевик</option>
                                                <option>Романтика</option>
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
                                    {ChangingComics(useLocation().state.UUID, name, rate, year, genre, price, episodes, description, image)}
                                </div>
                            </div>
                    </div>
                </div>
            </div>
        </>
    )
}