import PersonCard from "./PersonCard";

const PeopleSearchContainer = (props) => {
    return (
        <div className="card-columns">
            {props.people.map((person) => (
                <PersonCard
                    key={person.ID}
                    person={person} />
            ))}
        </div>
    );
};

export default PeopleSearchContainer;