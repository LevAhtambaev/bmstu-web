import {useEffect, useReducer} from "react";
import {getJsonAllComics} from "../modules";

const initialState = {comics: []}
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

export function GetAllComics() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `comics`

    useEffect(() => {
        getJsonAllComics(url).then((result) => {
            dispatch({type: success, comics: result})
        })
    }, [url])

    return state.comics
}