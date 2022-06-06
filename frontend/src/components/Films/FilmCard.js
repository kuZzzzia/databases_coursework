import { Link } from 'react-router-dom';

const FilmCard = (props) => {
    const poster = props.film.Poster;
    const name = props.film.Name;
    const altName = props.person.AltName.Valid ? props.person.AltName.String : '';
    const year = props.person.Year.Valid ? 'Год выпуска: ' + props.person.Year.Int16 : '';
    const duration = props.person.Duration.Valid ? 'Продолжительность: ' + props.person.Duration.Int16 + ' минут': '';
    const id = "/film/" + props.film.ID;

    return (
        <div className="card mb-5 pb-2" style={{maxWidth: '18rem'}}>
            <img className="card-img-top" src={poster}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">{year}</p>
                <p className="card-text">{duration}</p>
                <Link className="card-link-link" to={id}>View more</Link>
            </div>
        </div>
    );
};

export default FilmCard;