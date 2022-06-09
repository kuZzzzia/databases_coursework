import PlaylistRow from "./PlaylistRow";

const PlaylistsList = (props) => {
    return (
        <div>
            <h5>Подборки</h5>
            <div className="col " style={{overflowY: 'scroll', maxHeight: '20rem'}}>
                {props.playlists.map((playlist) => (
                    <PlaylistRow
                        key={playlist.ID}
                        playlist={playlist}
                        onDeletePlaylist={props.onDeletePlaylist}
                    />
                ))}
            </div>
        </div>
    );
};

export default PlaylistsList;