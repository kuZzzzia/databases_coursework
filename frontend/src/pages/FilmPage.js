import Film from '../components/Films/Film';
import {useParams} from "react-router-dom";

const FilmPage = () => {
    let { id } = useParams();
    return <Film id = {id}/>;
};

export default FilmPage;