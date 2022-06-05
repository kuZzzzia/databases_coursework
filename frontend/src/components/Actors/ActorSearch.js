import { useRef, useState } from "react";

import Errors from "../Errors/Errors";
import ActorsList from "./ActorsList";

const ActorSearch = () => {
    const [searchStatus, setSearchStatus] = useState(false);
    const [actors, setActors] = useState([]);
    const [errors, setErrors] = useState({});

    const actorRef = useRef();

    async function submitHandler(event) {
        event.preventDefault();
        setErrors({});

        const actorValue = actorRef.current.value;

        try {
            const response = await fetch("/api/actors",
                {
                    method: 'POST',
                    body: JSON.stringify({
                        Actor: actorValue,
                    }),
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'No actors found';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({'unknown': data['error']})
                } else {
                    setErrors(data['error']);
                }
            } else {
                setSearchStatus(true);
                setActors(data.data);
            }
        } catch (error) {
            setErrors({"error": error.message});
        }
    }

    const actorsContent = searchStatus ?
        actors.length === 0 ?
            <p>No actors found</p>
            :
            <ActorsList
                actors={actors}
            />
        : <p></p>;

    const header = 'Actors';
    const mainButtonText = 'Search';
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <section>
            <h1 className="text-center">{header}</h1>
            <div className="container w-75">
                <form onSubmit={submitHandler}>
                    <div className="form-row">
                        <div className="col-9">
                            <input id="username" type="text" className="form-control" placeholder={"Найти персону..."} required ref={actorRef} ></input>
                        </div>
                        <div className="col">
                            <button type="submit" className="btn btn-success mb-2">{mainButtonText}</button>
                        </div>
                    </div>
                </form>
            </div>
            {errorContent}
            {actorsContent}
        </section>
    );
};

export default ActorSearch;