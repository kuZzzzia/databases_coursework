import FilmCard from "./FilmCard";

const FilmSearchContainer = (props) => {
    return (
        <div className="card-columns">
            {props.films.map((film) => (
                <FilmCard
                    key={film.ID}
                    film={film}
                    onAddFilm={props.onAddFilm}
                />
            ))}
        </div>
    );
};

export default FilmSearchContainer;