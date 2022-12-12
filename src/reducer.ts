export const success = "Successs"

export function reducer(state : any, action: { type: any; comics: any; }) {
    switch (action.type) {
        case success:
            return {
                comics: action.comics
            }
        default:
            return state
    }
}