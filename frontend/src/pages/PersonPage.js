import Person from '../components/Person/Person';
import {useParams} from "react-router-dom";

const PersonPage = () => {
    let { id } = useParams();
    return <Person id = {id}/>;
};

export default PersonPage;