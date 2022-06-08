import {Link} from "react-router-dom";

const FilmRow = (props) => {
    return (
        <div className="row">
            <div className="col-sm">
                Фильм: {props.film.FilmName}
            </div>
            {props.film.Year.Valid ?
                <div className="col-sm">
                    Год производства: {props.film.Year.Int16}
                </div> : <div></div>}
            {props.film.FilmRating !== -1 ?
                <div className="col-2">
                    Рейтинг: {props.film.FilmRating}%
                </div> : <div>Нет оценки</div>}
            <div className="col-sm">
                <Link className="card-link-link" to={'/film/'+props.film.FilmID}>View more</Link>
            </div>
        </div>
    )
};

export default FilmRow;