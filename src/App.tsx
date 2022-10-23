import React from 'react';
import {Manga} from './components/Manga'
import {Mangas} from './repository/Mangas'


function App() {
    return (
        <div className="container mx-auto max-w-2xl pt-5">
            {Mangas.map((manga, key) => {
                return <Manga manga={manga} key={key}/>
            })}
        </div>
    )
}

export default App;
