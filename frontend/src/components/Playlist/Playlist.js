import {useState, useCallback, useEffect, useContext} from "react";

import Errors from "../Errors/Errors";
import Rate from "../Rate/Rate";
import AuthContext from '../../db/auth-context';
import FilmsList from "../Films/FilmsList";

const Playlist = (props) => {
    const authContext = useContext(AuthContext);

    const [errors, setErrors] = useState({});
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState({String: '', Valid: false});
    const [userName, setUserName] = useState({String: '', Valid: false});
    const [films, setFilms] = useState([]);
    const [likeAmount, setLikeAmount] = useState(0);
    const [dislikeAmount, setDislikeAmount] = useState(0);
    const [likeStatus, setLikeStatus] = useState(0);
    const [status, setStatus] = useState(false);

    const fetchPlaylistHandler = useCallback(async () => {
        setErrors({});
        try {
            const response = await fetch('/playlist/' + props.id);
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'No playlist found';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setTitle(data.playlist.Title);
                setDescription(data.playlist.Description);
                setUserName(data.playlist.UserName);
                setFilms(data.playlist.Films);
                setLikeAmount(data.playlist.LikeAmount);
                setDislikeAmount(data.playlist.DislikeAmount);
                if (authContext.loggedIn) {
                    try {
                        const response = await fetch('/auth/playlist/rateStatus/' + props.id,
                            {
                                method: "POST",
                                body: JSON.stringify({
                                    Src: 'playlist'
                                }),
                                headers: {
                                    'Content-Type': 'application/json',
                                    'Authorization': 'Bearer ' + authContext.token,
                                },
                            }
                        );
                        const data = await response.json();
                        if (response.ok) {
                            setLikeStatus(data.rate);
                        }
                    } catch (error) {
                        setErrors({"error": error.message});
                    }
                }
                setStatus(true);
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }, [props.id, authContext]);

    useEffect(() => {
        fetchPlaylistHandler();
    }, [fetchPlaylistHandler]);

    const filmsContent =
        films.length === 0 ?
            <p className="card-text">Нет фильмов в данной подборке</p>
            :
            <FilmsList
                films={films}
            />;

    const descriptionContent = description.Valid ?  'Описание: ' + description.String : '';
    const user = userName.Valid ? userName.String : 'Пользователь удалён';


    const Content = Object.keys(errors).length === 0 ?
        status ?
            <div className="card-body">
                <h5 className="card-title">Заголовок: {title}</h5>
                <h6 className="card-subtitle mb-2 text-muted">Автор подборки: {user}</h6>
                <p className="card-text">{descriptionContent}</p>
                {filmsContent}
                <Rate
                    key={props.id}
                    Like={likeAmount}
                    Dislike={dislikeAmount}
                    Status={likeStatus}
                    Addr={'playlist'}
                    ID={props.id}
                />
            </div>
            : <div>Processing...</div>
        : Errors(errors);

    return (
        <section>
            {Content}
        </section>
    );
};

export default Playlist;