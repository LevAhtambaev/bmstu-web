import {IManga, ICart, IOrder} from "../models";

export let manga_context : IManga = {
    UUID: "",
    Image: "",
    Year: 0,
    Name: "",
    Genre: "",
    Description: "",
    Episodes: 0,
    Rate: 1,
    Price: 0,
}

export let cart_context: ICart = {
    UUID: "",
    Manga: "",
}

export let orders_context: IOrder = {
    UUID: "",
    Mangas: [""],
    UserUUID: "",
    Date: "",
    Status: "",
}
