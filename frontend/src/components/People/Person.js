import RolesList from "./RolesList";
import FilmsList from "./FilmsList";

const Person = (props) => {
    const filmsContent =
        props.films.length === 0 ?
            <p>Еще не было режиссерских работ</p>
            :
            <FilmsList
                films={props.films}
            />;

    const rolesContent =
        props.films.length === 0 ?
            <p>Еще не было актёрских работ</p>
            :
            <RolesList
                roles={props.roles}
            />;

    const photo = props.person.Photo;
    const name = props.person.Name;
    const altName = props.person.AltName;
    const date = props.person.Date;

    return (
        <section>
            <div className="row " >
                <div className="col-lg-5">
                    <img src={photo}  alt={name} style={{width : '100%' }}/>
                </div>
                <div className="card-body">
                    <h5 className="card-title">{name}</h5>
                    <h6 className="card-subtitle mb-2 text-muted">{altName}</h6>
                    <p className="card-text">Дата рождения: {date}</p>
                </div>
            </div>
            {filmsContent}
            {rolesContent}
        </section>
    );
};

export default Person;