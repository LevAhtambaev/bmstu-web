import {deleteManga} from "../modules";


export function DeleteManga(uuid: string) {

    const url = `manga`

    function Delete() {
        deleteManga(url, uuid)
    }


    return (
        <form>
            <button onClick={() => Delete()}>Удалить</button>
        </form>
    );

}