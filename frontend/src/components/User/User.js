import {useState, useCallback, useEffect, useContext} from "react";

import Errors from "../Errors/Errors";
import PlaylistsList from "../Playlist/PlaylistsList"
import {Link, useNavigate} from "react-router-dom";
import AuthContext from '../../db/auth-context';

const User = (props) => {
    const authContext = useContext(AuthContext);

    const [errors, setErrors] = useState({});

    const [playlists, setPlaylists] = useState([]);
    const [username, setUsername] = useState('');
    const [status, setStatus] = useState(false);

    const fetchUserHandler = useCallback(async () => {
        setErrors({});
        try {
            const response = await fetch('/auth/profile/',
                {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'Error, while fetching info about playlists';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setPlaylists(data.playlists);
                setUsername(data.username);
                setStatus(true);
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }, [authContext]);

    useEffect(() => {
        fetchUserHandler().then();
    }, [fetchUserHandler]);

    async function deletePlaylistHandler(playlistID) {
        try {
            const method = 'POST';
            const response = await fetch('/auth/delete/playlist/' + playlistID,
                {
                    method: method,
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'Failed to delete playlist.';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setErrors({});
                setPlaylists((prevState) => {
                    return prevState.filter((i) => i.ID !== playlistID);
                });
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }

    const navigate = useNavigate();
    const logoutHandler = () => {
        authContext.logout();
        navigate('/');
    };

    async function deleteUserHandler() {
        try {
            const method = 'POST';
            const response = await fetch('/auth/delete/user',
                {
                    method: method,
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'Failed to add user.';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setErrors({});
                logoutHandler();
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }

    const playlistsContent =
        playlists.length === 0 ?
            <div className="row">
                <p className="col">У Вас пока что нет подборок</p>
                <Link className="card-link-link" to='/profile/playlist'>Создать</Link>
            </div>
            :
            <PlaylistsList
                playlists={playlists}
                onDeletePlaylist={deletePlaylistHandler}
            />;

    const usernameContent = 'Имя пользователя: ' + username;

    const Content = Object.keys(errors).length === 0 ?
        status ?
            <div>
                <div className="row " >
                    <div className="pr-5">{usernameContent}</div>
                    <button className="btn btn-primary" onClick={deleteUserHandler}>Удалить</button>
                </div>
                {playlistsContent}
            </div>
            : <div>Обработка...</div>
        : Errors(errors);

    return (
        <section>
            {Content}
        </section>
    );
};

export default User;