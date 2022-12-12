export interface IManga {
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
    Manga: string
}

export interface IOrder {
    UUID: string
    Mangas: string[]
    UserUUID: string
    Date: string
    Status: string
}
