import {IComics, ICart, IOrder} from "../models";

export let comics_context : IComics = {
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
    Comics: "",
}

export let orders_context: IOrder = {
    UUID: "",
    Comics: [""],
    UserUUID: "",
    Date: "",
    Status: "",
}
