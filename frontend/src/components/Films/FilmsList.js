import {Link} from "react-router-dom";

const FilmsList = (props) => {
    return (
        <div>
            {props.films.map((film) => (
                <div className="row">
                    <div className="col-sm">
                        {film.Name}
                    </div>
                    <div className="col-sm">
                        Год производства: {film.Year}
                    </div>
                    <div className="col-sm">
                        <Link className="card-link-link" to={'/person/'+film.ID}>View more</Link>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default FilmsList;