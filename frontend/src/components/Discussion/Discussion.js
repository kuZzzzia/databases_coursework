import {Link} from "react-router-dom";

const DiscussionList = (props) => {
    return (
        <section>
            <h4>Discussion</h4>
        </section>
        <div className="col">
            <h5>Персонажи</h5>
            {props.people.map((person) => (
                <div className="row">
                    <div className="col-sm">
                        {person.Name}
                    </div>
                    <div className="col-md-auto">
                        <Link className="card-link-link" to={'/person/'+person.ID}>View more</Link>
                    </div>
                    {person.Character.Valid ?
                        <div className="col-sm">
                            Имя персонажа: {person.Character.String}
                        </div> : <div></div>}
                </div>
            ))}
        </div>
    );
};

export default DiscussionList;