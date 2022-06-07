import {useState, useCallback, useEffect} from "react";

import Errors from "../Errors/Errors";
import PeopleList from "./PeopleList";
import PlaylistsList from "../Playlists/PlaylistsList"
import DiscussionList from "../Discussion/DiscussionList"
import {Link} from "react-router-dom";

const Film = (props) => {
    const [errors, setErrors] = useState({});
    const [film, setFilm] = useState({
        AltName: {String: "", Valid: false},
        Description: {String: "", Valid: false},
        Director: {String: "", Valid: false},
        Year: {Int16: 0, Valid: false},
        Duration: {Int16: 0, Valid: false},
        DirectorID: {Int16: 0, Valid: false},
    });
    const [people, setPeople] = useState([]);
    const [playlists, setPlaylists] = useState([]);
    const [discussion, setDiscussion] = useState([]);
    const [status, setStatus] = useState(false);

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
                setPeople(data.cast);
                setPlaylists(data.playlists);
                setDiscussion(data.discussion);
                setStatus(true);
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
            <p className="col">Еще нет назначенных на роли актеров</p>
            :
            <PeopleList
                people={people}
            />;

    const playlistsContent =
        playlists.length === 0 ?
            <p className="col">Нет подборок с данным фильмом</p>
            :
            <PlaylistsList
                playlists={playlists}
            />;
    //
    // const discussionContent =
    //     discussion.length === 0 ?
    //         <p></p>
    //         :
    //         <DiscussionList
    //             discussion={discussion}
    //         />;


    const poster = '/' + film.Poster;
    const name = film.Name;
    const altName = film.AltName.Valid ? film.AltName.String : '';
    const description = film.Description.Valid ? film.Description.String : '';
    const director = film.Director.Valid ? 'Режиссер' + film.Director.String : '';
    const duration = film.Duration.Valid ? 'Продолжительность: ' + film.Duration.Int16 + ' минут' : '';
    const year = film.Year.Valid ? 'Год производства: ' + film.Year.Int16 : '';
    const directorContent = film.DirectorID.Valid ?
        <div className="row">
            <p className="card-text">{director}</p>
            <Link className="card-link-link" to={'/person/' + film.DirectorID.Int16}>View more</Link>
        </div>
        : <div></div>;


    const Content = Object.keys(errors).length === 0 ?
        status ?
            <div>
                <div className="row " >
                    <div className="col-lg-4">
                        <img src={poster}  alt={name} style={{width : '100%' }}/>
                    </div>
                    <div className="card-body">
                        <h5 className="card-title">{name}</h5>
                        <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                        <p className="card-text">{year}</p>
                        <p className="card-text">{duration}</p>
                        <p className="card-text">{description}</p>
                        {directorContent}
                    </div>
                </div>
                <div className="row" >
                    {playlistsContent}
                    {peopleContent}
                </div>
                {/*{discussionContent}*/}
            </div>
            : <div>Processing...</div>
        : Errors(errors);

    return (
        <section>
            {Content}
        </section>
    );
};

export default Film;