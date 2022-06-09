import { Link } from 'react-router-dom';

const FilmCard = (props) => {
    const poster = '/' + props.film.Poster;
    const name = props.film.Name;
    const altName = props.film.AltName.Valid ? props.film.AltName.String : '';
    const year = props.film.Year.Valid ? 'Год выпуска: ' + props.film.Year.Int16 : '';
    const duration = props.film.Duration.Valid ? 'Продолжительность: ' + props.film.Duration.Int16 + ' минут': '';
    const id = "/film/" + props.film.ID;


    function submitHandler() {
        props.onAddFilm(props.film)
    }

    return (
        <div className="card mb-5 pb-2" style={{maxWidth: '18rem'}}>
            <img className="card-img-top" src={poster}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">{year}</p>
                <p className="card-text">{duration}</p>
                {props.film.Rating !== -1 ?
                    <p className="card-text">
                        Рейтинг: {props.film.Rating}%
                    </p> : <p className="card-text">Нет оценки</p>}
                { props.onAddFilm
                    ? <button className="btn btn-primary" onClick={submitHandler}>Добавить</button>
                    : <Link className="card-link-link" to={id}>View more</Link>
                }

            </div>
        </div>
    );
};

export default FilmCard;