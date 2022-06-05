import ActorRow from "./ActorRow";

const ActorsList = (props) => {
    return (
        <ul>
            {props.actors.map((actor) => (
                <ActorRow
                    key={actor.ID}
                    actor={actor} />
            ))}
        </ul>
    );
};

export default ActorsList;