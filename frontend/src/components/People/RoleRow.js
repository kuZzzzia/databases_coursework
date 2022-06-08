import {Link} from "react-router-dom";

const RoleRow = (props) => {
    return (
        <div className="row">
            <div className="col-sm">
                Фильм: {props.role.FilmName}
            </div>
            {props.role.Year.Valid ?
                <div className="col-sm">
                    Год производства: {props.role.Year.Int16}
                </div> : <div></div>}
            {props.role.FilmRating !== -1 ?
                <div className="col-2">
                    Рейтинг: {props.role.FilmRating}%
                </div> : <div>Нет оценки</div>}
            {props.role.Name.Valid ?
                <div className="col-sm">
                    Имя персонажа: {props.role.Name.String}
                </div> : <div></div>}
            <div className="col-sm">
                <Link className="card-link-link" to={'/film/'+props.role.FilmID}>View more</Link>
            </div>
        </div>
    )
};

export default RoleRow;