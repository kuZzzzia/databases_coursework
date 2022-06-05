import { Link } from 'react-router-dom';

const PersonRow = (props) => {
    const photo = props.person.Image;
    const name = props.person.Name;
    const altName = props.person.AltName;
    const date = props.person.Date;
    const id = ":" + props.person.ID;

    return (
        <div className="card mb-5 pb-2">
            <img className="card-img-left" src={photo}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">Дата рождения: {date}</p>
                <Link className="card-link-link" to={id}>View more</Link>
            </div>
        </div>
    );
};

export default PersonRow;