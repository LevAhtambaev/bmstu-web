import {useEffect, useReducer} from "react";
import {getJsonComics} from "../modules";

const initialState = {comics: ""}
const success = "Success"

function reducer(state: any, action: { type: any; comics: any; }) {
    switch (action.type) {
        case success:
            return {
                comics: action.comics
            }
        default:
            return state
    }
}

export function GetComics(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `comics/${uuid}`

    useEffect(() => {
        getJsonComics(url).then((result) => {
            dispatch({type: success, comics: result})
        })
    }, [url])

    return state.comics
}