import Person from '../components/People/Person';
import {useParams} from "react-router-dom";

const PersonPage = (props) => {
    let { id } = useParams();
    return <Person id = {id}/>;
};

export default PersonPage;