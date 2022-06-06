import {useState, useCallback, useEffect} from "react";

import RolesList from "./RolesList";
import FilmsList from "../Films/FilmsList";
import Errors from "../Errors/Errors";

const Person = (props) => {
    const [errors, setErrors] = useState({});
    const [films, setFilms] = useState([]);
    const [roles, setRoles] = useState([]);
    const [person, setPerson] = useState({
        AltName: {String: "", Valid: false},
        Date: {String: "", Valid: false}
    });
    const [status, setStatus] = useState(false);

    const fetchPersonHandler = useCallback(async () => {
        setErrors({});
        setStatus(false);
        try {
            const response = await fetch('/person/' + props.id);
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'No person found';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                setPerson(data.person);
                setFilms(data.films);
                setRoles(data.roles);
                setStatus(true);
            }
        } catch (error) {
            setErrors({ "error": error.message });
        }
    }, [props.id]);

    useEffect(() => {
        fetchPersonHandler();
    }, [fetchPersonHandler]);


    const filmsContent =
        films.length === 0 ?
            <p>Еще не было режиссерских работ</p>
            :
            <FilmsList
                films={films}
            />;

    const rolesContent =
        roles.length === 0 ?
            <p>Еще не было актёрских работ</p>
            :
            <RolesList
                roles={roles}
            />;

    const photo = '/' + person.Photo;
    const name = person.Name;
    const altName = person.AltName.Valid ? person.AltName.String : '';
    const date = person.Date.Valid ? 'Дата рождения: ' + person.Date.String : '';

    const Content = Object.keys(errors).length === 0 ?
       status ?
           <div>
                <div className="row " >
                    <div className="col-lg-4">
                        <img src={photo}  alt={name} style={{width : '100%' }}/>
                    </div>
                    <div className="card-body">
                        <h5 className="card-title">{name}</h5>
                        <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                        <p className="card-text">{date}</p>
                    </div>
                </div>
               <div>
                   <h5>Режиссерские работы</h5>
                   {filmsContent}
               </div>
               <div>
                   <h5>Актёрские работы</h5>
                   {rolesContent}
               </div>
           </div>
           : <div>Processing...</div>
        : Errors(errors);

    return (
        <section>
            {Content}
        </section>
    );
};

export default Person;