import {ENDPOINT} from "./App";
import {IManga, ICart} from "./models";
import axios from "axios";



export const getJsonMangas = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IManga[]>)
    return res
}

export const getJsonManga = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IManga>)
    return res
}

export const getJsonCart = async (url: string) => {
    const tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    access_token = access_token.replace(";", "")
    return axios.get(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (r) {
        return r.data
    }).catch((error)=>{
        window.location.replace("/error")
    })
}

export const deleteCart = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`, {method: "DELETE"})
    window.location.replace("/cart")
    return res
}

export const addToCart = async (url: string, uuid: string) => {
    const res = await fetch(`${ENDPOINT}/${url}` , {
        method: "POST",
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify({Manga: uuid})
    })
    return res
}

export function createUser(url: string, name: string, pass: string) {
        const body = {name: name, pass: pass}
        return axios.post(`${ENDPOINT}/${url}`, body).then(function (response) {
            console.log(response);
        })

}


export function loginUser (url: string, name: string, pass: string)  {
    const body = { login: name, password: pass }
    return axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true}).then(function (response) {
        console.log(response)
        window.location.replace("/manga")
    }).catch(function (reason) {
        window.location.replace("/login")
    })
}

export function logoutUser (url: string) {
    const tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    access_token = access_token.replace(";", "")
    console.log(access_token)
    return axios.get(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (r) {
            console.log(r.data)
        window.location.replace("/manga")
    })
}

export function checkToken() {
    let tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    access_token = access_token.replace(";", "")
    let showAddCartButton = true
    if (access_token == "") {
        showAddCartButton = false
    }
    return showAddCartButton
}