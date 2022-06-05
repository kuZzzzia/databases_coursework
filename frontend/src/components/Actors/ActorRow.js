import { Link } from 'react-router-dom';

const ActorRow = (props) => {
    const photo = props.actor.image;
    const name = props.actor.name;
    const altName = props.actor.altName;
    const date = props.actor.date;

    return (
        <div className="card mb-5 pb-2">
            <img className="card-img-left" src={photo}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">Дата рождения: {date}</p>
                <Link className="card-link-link" to="/auth">View more</Link>
            </div>
        </div>
    );
};

export default ActorRow;