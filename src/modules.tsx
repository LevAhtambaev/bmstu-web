import {ENDPOINT} from "./App";
import {IManga} from "./models";


export const getJson = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IManga[]>)
    return res
}