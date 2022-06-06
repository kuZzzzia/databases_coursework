import {useState, useCallback, useEffect} from "react";

import Errors from "../Errors/Errors";
import FilmsList from "./FilmsList";
import PeopleList from "./PeopleList";
import PlaylistsList from "../Playlists/PlaylistsList"
import DiscussionList from "../Discussion/DiscussionList"

const Film = (props) => {
    const [errors, setErrors] = useState({});
    const [film, setFilm] = useState({});
    const [people, setPeople] = useState([]);
    const [alikeFilms, setAlikeFilms] = useState([]);
    const [playlists, setPlaylists] = useState([]);
    const [discussion, setDiscussion] = useState([]);

    const fetchFilmHandler = useCallback(async () => {
        setErrors({});
        try {
            const response = await fetch('/film/' + props.id);
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'No film found';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setFilm(data.film);
                setPeople(data.people);
                setAlikeFilms(data.alikeFilms);
                setPlaylists(data.playlists);
                setDiscussion(data.discussion);
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }, [props.id]);

    useEffect(() => {
        fetchFilmHandler();
    }, [fetchFilmHandler]);


    const peopleContent =
        people.length === 0 ?
            <p>Еще нет назначенных на роли актеров</p>
            :
            <PeopleList
                people={people}
            />;

    const alikeFilmsContent =
        alikeFilms.length === 0 ?
            <p>Нет похожих фильмов</p>
            :
            <FilmsList
                films={alikeFilms}
            />;

    const playlistsContent =
        playlists.length === 0 ?
            <p>Нет подборок с данным фильмом</p>
            :
            <PlaylistsList
                playlists={playlists}
            />;

    const discussionContent =
        discussion.length === 0 ?
            <p></p>
            :
            <DiscussionList
                discussion={discussion}
            />;


    const poster = '/' + film.Poster;
    const name = film.Name;
    const altName = film.AltName;
    const year = film.Year;

    const Content = Object.keys(errors).length === 0 ?
        <div>
            <div className="row " >
                <div className="col-lg-4">
                    <img src={poster}  alt={name} style={{width : '100%' }}/>
                </div>
                <div className="card-body">
                    <h5 className="card-title">{name}</h5>
                    <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                    <p className="card-text">Год производства: {year}</p>
                </div>
            </div>
            <div className="row " >
                {alikeFilmsContent}
                {playlistsContent}
            </div>
            {peopleContent}
            {discussionContent}
        </div>
        : Errors(errors);

    return (
        <section>
            {Content}
        </section>
    );
};

export default Film;