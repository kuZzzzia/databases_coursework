import {Link} from "react-router-dom";

const PlaylistRow = (props) => {
    return (
        <div className="row">
            <div className="col-sm">
                {props.playlist.Title}
            </div>
            <div className="col-sm">
                {props.playlist.Rating === -1 ? 'Рейтинг: ' + props.playlist.Rating : 'Нет оценки'}
            </div>
            <div className="col-sm">
                <Link className="card-link-link" to={'/playlist/'+ props.playlist.ID}>View more</Link>
            </div>
        </div>
    )
};

export default PlaylistRow;