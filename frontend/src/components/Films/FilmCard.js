import { Link } from 'react-router-dom';

const FilmCard = (props) => {
    const poster = props.film.Poster;
    const name = props.film.Name;
    const altName = props.film.AltName;
    const year = props.film.Year;
    const duration = props.film.Duration;
    const id = ":" + props.film.ID;

    return (
        <div className="card mb-5 pb-2" style={{maxWidth: '18rem'}}>
            <img className="card-img-top" src={poster}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">Год выпуска: {year}</p>
                <p className="card-text">Продолжительность: {duration} минут</p>
                <Link className="card-link-link" to={id}>View more</Link>
            </div>
        </div>
    );
};

export default FilmCard;