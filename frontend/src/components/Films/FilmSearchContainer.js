import FilmCard from "./FilmCard";

const FilmSearchContainer = (props) => {
    return (
        <div className="card-columns">
            {props.films.map((film) => (
                <FilmCard
                    key={film.ID}
                    film={film} />
            ))}
        </div>
    );
};

export default FilmSearchContainer;