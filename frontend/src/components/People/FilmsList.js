const FilmList = (props) => {
    return (
        <div>
            {props.films.map((film) => (
                <div className="row">
                    <div className="col-lg-5">
                        {film.Name}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default RolesList;