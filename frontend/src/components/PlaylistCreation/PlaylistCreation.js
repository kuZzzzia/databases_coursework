import { useRef, useState, useContext} from "react";
import { useNavigate } from 'react-router-dom';

import Errors from "../Errors/Errors";
import FilmSearch from "../FilmSearch/FilmSearch";
import FilmsList from "../Film/FilmsList";
import AuthContext from '../../db/auth-context';

const PlayListCreation = () => {
    const navigate = useNavigate();
    const authContext = useContext(AuthContext);

    const [films, setFilms] = useState([]);
    const [errors, setErrors] = useState({});

    const titleRef = useRef();
    const descriptionRef = useRef();

    async function submitHandler(event) {
        event.preventDefault();
        setErrors({});

        const titleValue = titleRef.current.value;
        const descriptionValue = descriptionRef.current.value;

        try {
            const response = await fetch("/auth/playlist/create",
                {
                    method: 'POST',
                    body: JSON.stringify({
                        Title: titleValue,
                        Description: { String: descriptionValue},
                        Films: films
                    }),
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'не удалось создать подборку';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({'unknown': data['error']})
                } else {
                    setErrors(data['error']);
                }
            } else {
                navigate('/profile');
            }
        } catch (error) {
            setErrors({'error': error.message});
        }
    }

    const addFilmHandler = (filmData) => {
        setFilms((prevState) => {
            const dup = prevState.find(item => item.ID === filmData.ID)
            if (dup) {
                setErrors({'error': 'Film ' + filmData.Name + ' is already added'})
                return prevState
            } else {
                return [...prevState, filmData]
            }
        });
    }

    const header = 'Подборка';
    const mainButtonText = 'Создать';
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);


    return (
        <section>
            <h3 className="text-center">{header}</h3>

            <div className="container w-75">
                <form onSubmit={submitHandler}>
                    <div className="form-group pb-3">
                        <label htmlFor="title">Заголовок</label>
                        <input id="title" type="text" className="form-control" placeholder={"Найти фильм..."} required ref={titleRef} ></input>
                    </div>
                    <div className="form-group pb-2">
                        <label htmlFor="description">Описание</label>
                        <textarea id="description" className="form-control" placeholder={"Описание..."} rows="5" ref={descriptionRef} ></textarea>
                    </div>
                    <FilmsList films={films}/>
                    <div className="pb-2 d-flex justify-content-between">
                        <button type="submit" className="btn btn-success">{mainButtonText}</button>
                    </div>
                </form>
            </div>
            {errorContent}
            <FilmSearch onAddFilm={addFilmHandler}/>
        </section>
    );
};

export default PlayListCreation;