import {addOrder} from "../modules";

export function AddOrder(comics_uuid: string[]) {

    const url = `orders`

    function Add() {
        addOrder(url, comics_uuid)
    }


    return (
        <>
            <button className="mt-2 border-4 border-indigo-700 text-indigo-700 hover:bg-blue-700 hover:text-white sm:px-3 place-content-between rounded-full text-xl font-bold" onClick={() => Add()}>Приобрести</button>
        </>
    );

}