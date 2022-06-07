import PersonRow from "./PersonRow";

const PeopleList = (props) => {
    return (
        <div className="col">
            <h5>Персонажи</h5>
            {props.people.map((person) => (
                <PersonRow
                    key={person.ID}
                    person={person}
                />
            ))}
        </div>
    );
};

export default PeopleList;