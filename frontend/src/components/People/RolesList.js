import {Link} from "react-router-dom";

const RolesList = (props) => {
    return (
        <div>
            {props.roles.map((role) => (
                <div className="row">
                    <div className="col-sm">
                        Имя персонажа: {role.Name}
                    </div>
                    <div className="col-sm">
                        Фильм: {role.FilmName}
                    </div>
                    <div className="col-sm">
                        Год производства: {role.Year}
                    </div>
                    <div className="col-sm">
                        <Link className="card-link-link" to={'/film/'+role.FilmID}>View more</Link>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default RolesList;