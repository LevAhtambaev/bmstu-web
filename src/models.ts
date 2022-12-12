export interface IComics {
    UUID: string
    Name: string
    Rate: number
    Year: number
    Genre: string
    Price: number
    Episodes: number
    Description: string
    Image: string
}

export interface ICart {
    UUID: string
    Comics: string
}

export interface IOrder {
    UUID: string
    Comics: string[]
    UserUUID: string
    Date: string
    Status: string
}
