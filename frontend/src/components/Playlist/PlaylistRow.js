import {Link} from "react-router-dom";

const PlaylistRow = (props) => {
    function submitHandler() {
        props.onDeletePlaylist(props.playlist.ID)
    }
    return (
        <div className="row pb-2">
            <div className="col-sm">
                {props.playlist.Title}
            </div>
            <div className="col-md-auto">
                <Link className="card-link-link" to={'/playlist/'+ props.playlist.ID}>View more</Link>
            </div>
            <div className="col-sm">
                {props.playlist.Rating !== -1 ? 'Рейтинг: ' + props.playlist.Rating + '%' : 'Нет оценки'}
            </div>
            {props.onDeletePlaylist && (
                <button className="btn btn-danger" onClick={submitHandler}>Удалить</button>
            )}
        </div>
    )
};

export default PlaylistRow;