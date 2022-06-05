import PersonRow from "./PersonRow";

const PeopleList = (props) => {
    return (
        <ul>
            {props.people.map((person) => (
                <PersonRow
                    key={person.ID}
                    person={person} />
            ))}
        </ul>
    );
};

export default PeopleList;