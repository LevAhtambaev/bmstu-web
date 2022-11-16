
import {useEffect, useReducer} from "react";
import {getJsonManga} from "../modules";

const initialState = {manga: ""}
const success = "Success"

function reducer(state: any, action: { type: any; manga: any; }) {
    switch (action.type) {
        case success:
            return {
                manga: action.manga
            }
        default:
            return state
    }
}

export function GetManga(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `manga/${uuid}`

    useEffect(() => {
        getJsonManga(url).then((result) => {
            dispatch({type: success, manga: result})
        })
    }, [url])

    return state.manga
}