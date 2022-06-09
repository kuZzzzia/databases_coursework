import { Link } from 'react-router-dom';

const PersonCard = (props) => {
    const photo = props.person.Photo;
    const name = props.person.Name;
    const altName = props.person.AltName.Valid ? props.person.AltName.String : '';
    const date = props.person.Date.Valid ? 'Дата рождения: ' + props.person.Date.String : '';
    const id = "/person/" + props.person.ID;

    return (
        <div className="card mb-5 pb-2" style={{maxWidth: '18rem'}}>
            <img className="card-img-top" src={photo}  alt={name}/>
            <div className="card-body">
                <h5 className="card-title">{name}</h5>
                <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                <p className="card-text">{date}</p>
                <Link className="card-link-link" to={id}>Подробнее</Link>
            </div>
        </div>
    );
};

export default PersonCard;