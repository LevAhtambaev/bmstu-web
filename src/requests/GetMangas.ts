import {useEffect, useReducer} from "react";
import {getJsonMangas} from "../modules";

const initialState = {mangas: []}
const success = "Success"

function reducer(state: any, action: { type: any; mangas: any; }) {
    switch (action.type) {
        case success:
            return {
                mangas: action.mangas
            }
        default:
            return state
    }
}

export function GetMangas() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `mangas`

    useEffect(() => {
        getJsonMangas(url).then((result) => {
            dispatch({type: success, mangas: result})
        })
    }, [url])

    return state.mangas
}