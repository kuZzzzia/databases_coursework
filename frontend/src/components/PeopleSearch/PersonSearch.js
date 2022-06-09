import { useRef, useState } from "react";

import Errors from "../Errors/Errors";
import PeopleSearchContainer from "./PeopleSearchContainer";

const PersonSearch = () => {
    const [searchStatus, setSearchStatus] = useState(false);
    const [people, setPeople] = useState([]);
    const [errors, setErrors] = useState({});

    const personRef = useRef();

    async function submitHandler(event) {
        event.preventDefault();
        setErrors({});

        const personValue = personRef.current.value;

        try {
            const response = await fetch("/api/people",
                {
                    method: 'POST',
                    body: JSON.stringify({
                        Pattern: personValue,
                    }),
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'ошибочный запрос';
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
                setPeople(data.data);
            }
        } catch (error) {
            setErrors({"error": error.message});
        }
    }

    const peopleContent = searchStatus ?
        people.length === 0 ?
            <p>Людей по данному запросу не найдено</p>
            :
            <PeopleSearchContainer
                people={people}
            />
        : <p></p>;

    const header = 'Персоны';
    const mainButtonText = 'Поиск';
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <section>
            <h1 className="text-center">{header}</h1>
            <div className="container w-75 pb-3">
                <form onSubmit={submitHandler}>
                    <div className="form-row pb-2">
                        <div className="ml-3 pb-2 col-10 d-flex justify-content-center">
                            <input id="username" type="text" className="form-control" placeholder={"Найти персону..."} required ref={personRef} ></input>
                        </div>
                        <div className="col d-flex justify-content-center">
                            <button type="submit" className="btn btn-success mb-2">{mainButtonText}</button>
                        </div>
                    </div>
                </form>
            </div>
            {errorContent}
            {peopleContent}
        </section>
    );
};

export default PersonSearch;