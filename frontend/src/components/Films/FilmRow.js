import { Link } from 'react-router-dom';

const FilmRow = (props) => {
    const poster = props.film.poster;
    const name = props.film.name;
    const altName = props.film.altName;
    const year = props.film.year;
    const duration = props.film.duration;

    return (
        <div className="card mb-5 pb-2">
            <img className="card-img-left" src={poster}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">Год выпуска: {year}</p>
                <p className="card-text">Продолжительность: {duration} минут</p>
                <Link className="card-link-link" to="/auth">View more</Link>
            </div>
        </div>
    );
};

export default FilmRow;