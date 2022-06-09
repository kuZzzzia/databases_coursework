import PersonRow from "./PersonRow";

const PeopleList = (props) => {
    return (
        <div>
            <h5>В ролях</h5>
            <div className="col">

                {props.people.map((person) => (
                    <PersonRow
                        key={person.ID}
                        person={person}
                    />
                ))}
            </div>
        </div>
    );
};

export default PeopleList;