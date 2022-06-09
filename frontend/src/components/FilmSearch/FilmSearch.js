import { useRef, useState, useCallback, useEffect } from "react";

import Errors from "../Errors/Errors";
import FilmSearchContainer from "./FilmSearchContainer";

const FilmSearch = (props) => {
    const [searchStatus, setSearchStatus] = useState(false);
    const [films, setFilms] = useState([]);
    const [errors, setErrors] = useState({});
    const [genres, setGenres] = useState([]);
    const [countries, setCountries] = useState([]);

    const filmRef = useRef();
    const countryRef = useRef();
    const genreRef = useRef();

    async function fetchRequest(filmValue, genreValue, countryValue) {
        try {
            const response = await fetch("/api/films",
                {
                    method: 'POST',
                    body: JSON.stringify({
                        Pattern: filmValue,
                        Country: countryValue,
                        Genre: genreValue
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

    const fetchFilmSearchHandler = useCallback(async () => {
        setErrors({});
        try {
            const response = await fetch("/api/categories",
                {
                    method: 'POST',
                }
            );
            const data = await response.json();
            if (response.ok) {
                setGenres(data.genres);
                setCountries(data.countries);
            }
        } catch (error) {
            setErrors({"error": error.message});
        }
        await fetchRequest()
    }, []);

    useEffect(() => {
        fetchFilmSearchHandler().then();
    }, [fetchFilmSearchHandler]);

    async function submitHandler(event) {
        event.preventDefault();
        setErrors({});

        const filmValue = filmRef.current.value;
        const genreValue = genreRef.current.value;
        const countryValue = countryRef.current.value;
        await fetchRequest(filmValue, genreValue, countryValue);
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
                    <div className="form-row pb-2">
                        <div className="ml-3 pb-2 col-10 d-flex justify-content-center">
                            <input id="username" type="text" className="form-control" placeholder={"Найти фильм..."} ref={filmRef} ></input>
                        </div>
                        <div className="col d-flex justify-content-center">
                            <button type="submit" className="btn btn-success mb-2">{mainButtonText}</button>
                        </div>
                        <div className="form-group ml-3 col-md-4">
                            <label htmlFor="inputState">Жанр</label>
                            <select id="inputState" className="form-control" defaultValue="" ref={genreRef}>
                                <option value="" disabled>Выбрать...</option>
                                {genres.map((genre) => (
                                    <option value={genre}>{genre}</option>
                                ))}
                            </select>
                        </div>
                        <div className="form-group ml-3 col-md-4">
                            <label htmlFor="inputState">Страна</label>
                            <select id="inputState" className="form-control" defaultValue="" ref={countryRef}>
                                <option value="" disabled>Выбрать...</option>
                                {countries.map((country) => (
                                    <option value={country}>{country}</option>
                                ))}
                            </select>
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