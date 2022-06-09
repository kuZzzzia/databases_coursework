import {Link} from "react-router-dom";

const FilmRow = (props) => {
    return (
        <div className="row pb-2">
            <div className="col-sm">
                Фильм: {props.film.Name}
            </div>
            {props.film.Year.Valid ?
                <div className="col-sm">
                    Год производства: {props.film.Year.Int16}
                </div> : <div></div>}
            <div className="col-2">
                {props.film.Rating !== -1 ? 'Рейтинг: ' + props.film.Rating + '%' : 'Нет оценки'}
            </div>
            <div className="col-sm">
                <Link className="card-link-link" to={'/film/'+props.film.ID}>Подробнее</Link>
            </div>
        </div>
    )
};

export default FilmRow;