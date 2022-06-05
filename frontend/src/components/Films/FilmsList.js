import FilmRow from "./FilmRow";

const FilmsList = (props) => {
    return (
        <ul>
            {props.films.map((film) => (
                <FilmRow
                    key={film.ID}
                    film={film} />
            ))}
        </ul>
    );
};

export default FilmsList;