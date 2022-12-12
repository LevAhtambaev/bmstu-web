import {ENDPOINT} from "./App";
import {IComics} from "./models";
import axios from "axios";



export const getJsonAllComics = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IComics[]>)
    return res
}

export const getJsonComics = async (url: string) => {
    const res = await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IComics>)
    return res
}

export function getToken() {
    let tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    return access_token.replace(";", "")
}

export function getRole(token: string) {
    return axios.get(`${ENDPOINT}/role`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${token}`
        }}).then(r => r.data)
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
    let access_token = getToken()
    return axios.delete(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
}

export const addToCart = async (url: string, uuid: string) => {
    const body = {Comics: uuid}
    let access_token = getToken()
    return  axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (response) {
        console.log(response);
    })
}

export function addComics(url: string, name: string, rate: number, year: number, genre: string, price: number, episodes: number, description: string, image: string) {
    const body = {
        Name: name,
        Rate: rate,
        Year: year,
        Genre: genre,
        Price: price,
        Volumes: episodes,
        Description: description,
        Image: image,
    }
    let access_token = getToken()
    console.log(body)
    return axios.post(`${ENDPOINT}/${url}`, body, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(function (response) {
        console.log(response);
    })
}

export function changeComics(uuid:string, url: string, name: string, rate: number, year: number, genre: string, price: number, episodes: number, description: string, image: string) {
    const body = {
        Name: name,
        Rate: rate,
        Year: year,
        Genre: genre,
        Price: price,
        Volumes: episodes,
        Description: description,
        Image: image,
    }
    let access_token = getToken()
    console.log(body)
    return axios.put(`${ENDPOINT}/${url}/${uuid}`, body, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(function (response) {
        console.log(response);
    })
}

export function deleteComics (url: string, uuid: string) {
    let access_token = getToken()
    return axios.delete(`${ENDPOINT}/${url}/${uuid}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
}

export function addOrder (url: string, comics_uuid: string[])  {
    const body = { Comics: comics_uuid }
    let access_token = getToken()
    return  axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (response) {
        console.log(response);
    })

}

export function updateStatus(token: string, uuid: string, status: string) {
    const body = { Status: status }
    return axios.put(`${ENDPOINT}/orders/${uuid}`, body,{withCredentials: true, headers: {
            "Authorization": `Bearer ${token}`
        }}).then(r => r.data)
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
        window.location.replace("/comics")
    }).catch(function (reason) {
        window.location.replace("/login")
    })
}

export function logoutUser (url: string) {
    let access_token = getToken()
    return axios.get(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(function (r) {
            console.log(r.data)
        window.location.replace("/login")
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