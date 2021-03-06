import {Link} from "react-router-dom";

const PersonRow = (props) => {
    return (
        <div className="row pb-2">
            <div className="col-sm">
                {props.person.Name}
            </div>
            <div className="col-md-auto">
                <Link className="card-link-link" to={'/person/'+props.person.ID}>Подробнее</Link>
            </div>
            {props.person.Character.Valid
                ? <div className="col-sm">Имя персонажа: {props.person.Character.String}</div>
                : <div></div>}
        </div>
    )
};

export default PersonRow;