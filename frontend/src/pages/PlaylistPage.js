import Playlist from '../components/Playlist/Playlist';
import {useParams} from "react-router-dom";

const PlaylistPage = () => {
    let { id } = useParams();
    return <Playlist id = {id}/>;
};

export default PlaylistPage;