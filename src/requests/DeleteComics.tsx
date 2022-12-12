import {deleteComics} from "../modules";


export function DeleteComics(uuid: string) {

    const url = `comics`

    function Delete() {
        deleteComics(url, uuid)
    }


    return (
        <form>
            <button onClick={() => Delete()}>Удалить</button>
        </form>
    );

}