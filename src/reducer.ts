export const success = "Successs"

export function reducer(state : any, action: { type: any; mangas: any; }) {
    switch (action.type) {
        case success:
            return {
                mangas: action.mangas
            }
        default:
            return state
    }
}