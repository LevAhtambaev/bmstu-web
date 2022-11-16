import {ENDPOINT} from "./App";
import {IManga, ICart} from "./models";


export const getJsonMangas = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IManga[]>)
    return res
}

export const getJsonManga = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IManga>)
    return res
}

export const getJsonCart = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<ICart[]>)
    return res
}

export const deleteCart = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`, {method: "DELETE"})
    return res
}

export const addToCart = async (url: string, uuid: string) => {
    const res = await fetch(`${ENDPOINT}/${url}` , {
        method: "POST",
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify({Car: uuid})
    })
    return res
}