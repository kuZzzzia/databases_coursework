import FilmRow from "./FilmRow";

const FilmsList = (props) => {
    return (
        <div>
            {props.films.map((film) => (
                <FilmRow
                    key={film.FilmID}
                    film={film}
                />
            ))}
        </div>
    );
};

export default FilmsList;