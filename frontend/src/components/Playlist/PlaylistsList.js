import {Link} from "react-router-dom";

const PlaylistsList = (props) => {
    return (
        <div className="col">
            <h5>Подборки</h5>
            {props.playlists.map((playlist) => (
                <div className="row">
                    <div className="col-sm">
                        {playlist.Title}
                    </div>
                    <div className="col-sm">
                        {playlist.Rating === -1 ? 'Рейтинг: ' + playlist.Rating : 'Нет оценки'}
                    </div>
                    <div className="col-sm">
                        <Link className="card-link-link" to={'/playlist/'+playlist.ID}>View more</Link>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default PlaylistsList;