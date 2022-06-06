import {Link} from "react-router-dom";

const PlaylistsList = (props) => {
    return (
        <div>
            <h5>Подборки</h5>
            {props.playlists.map((playlist) => (
                <div className="row">
                    <div className="col-sm">
                        {playlist.Name}
                    </div>
                    <div className="col-sm">
                        Год производства: {playlist.Year}
                    </div>
                    <div className="col-sm">
                        <Link className="card-link-link" to={'/person/'+playlist.ID}>View more</Link>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default PlaylistsList;