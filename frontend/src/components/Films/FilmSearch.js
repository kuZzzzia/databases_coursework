import { useRef, useState } from "react";

import Errors from "../Errors/Errors";
import FilmSearchContainer from "./FilmSearchContainer";

const FilmSearch = (props) => {
    const [searchStatus, setSearchStatus] = useState(false);
    const [films, setFilms] = useState([]);
    const [errors, setErrors] = useState({});

    const filmRef = useRef();

    async function submitHandler(event) {
        event.preventDefault();
        setErrors({});

        const filmValue = filmRef.current.value;

        try {
            const response = await fetch("/api/films",
                {
                    method: 'POST',
                    body: JSON.stringify({
                        Pattern: filmValue,
                    }),
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'No actors found';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({'unknown': data['error']})
                } else {
                    setErrors(data['error']);
                }
            } else {
                setSearchStatus(true);
                setFilms(data.data);
            }
        } catch (error) {
            setErrors({"error": error.message});
        }
    }

    const filmsContent = searchStatus ?
        films.length === 0 ?
            <p>No films found</p>
            :
            props.onAddFilm
                ? <FilmSearchContainer
                films={films}
                onAddFilm={props.onAddFilm}
                />
                : <FilmSearchContainer
                    films={films}
                />
        : <p></p>;

    const header = 'Films';
    const mainButtonText = 'Search';
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <section>
            <h1 className="text-center">{header}</h1>
            <div className="container w-75 pb-5">
                <form onSubmit={submitHandler}>
                    <div className="form-row">
                        <div className="ml-3 col-10 d-flex justify-content-center">
                            <input id="username" type="text" className="form-control" placeholder={"Найти фильм..."} required ref={filmRef} ></input>
                        </div>
                        <div className="col d-flex justify-content-center">
                            <button type="submit" className="btn btn-success mb-2">{mainButtonText}</button>
                        </div>
                    </div>
                </form>
            </div>
            {errorContent}
            {filmsContent}
        </section>
    );
};

export default FilmSearch;