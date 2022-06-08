import PlaylistRow from "./PlaylistRow";

const PlaylistsList = (props) => {
    return (
        <div className="col">
            <h5>Подборки</h5>
            {props.playlists.map((playlist) => (
                <PlaylistRow
                    key={playlist.ID}
                    playlist={playlist}
                />
            ))}
        </div>
    );
};

export default PlaylistsList;