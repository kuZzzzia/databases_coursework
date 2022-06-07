import {Link} from "react-router-dom";

const FilmsList = (props) => {
    return (
        <div>
            {props.films.map((film) => (
                <div className="row">
                    <div className="col-sm">
                        Фильм: {film.FilmName}
                    </div>
                    {film.Year.Valid ?
                        <div className="col-sm">
                            Год производства: {film.Year.Int16}
                        </div> : <div></div>}
                    {film.FilmRating !== -1 ?
                        <div className="col-2">
                            Рейтинг: {film.FilmRating}%
                        </div> : <div>Нет оценки</div>}
                    <div className="col-sm">
                        <Link className="card-link-link" to={'/film/'+film.FilmID}>View more</Link>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default FilmsList;