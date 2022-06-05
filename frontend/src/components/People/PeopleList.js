import PersonRow from "./PersonRow";

const PeopleList = (props) => {
    return (
        <div className="card-columns">
            {props.people.map((person) => (
                <PersonRow
                    key={person.ID}
                    person={person} />
            ))}
        </div>
    );
};

export default PeopleList;