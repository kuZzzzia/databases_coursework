import {useState, useContext} from "react";
import Errors from "../Errors/Errors";
import AuthContext from '../../db/auth-context';

const Rate = (props) => {
    const authContext = useContext(AuthContext);

    const [errors, setErrors] = useState({});
    const [state, setState] = useState({
        Like: props.Like,
        Dislike: props.Dislike,
        Status: props.Status
    });

    async function handleLike(event) {
        event.preventDefault();
        if (state.Status === 0 || state.Status === -1) {
            try {
                const response = await fetch('/auth/' + props.Addr + '/rate/' + props.ID,
                    {
                        method: 'POST',
                        body: JSON.stringify({
                            Like : true,
                            Src : props.Addr
                        }),
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + authContext.token,
                        },
                    }
                );
                await response.json();
                if (!response.ok) {
                    let errorText = 'не удалось оставить оценку';
                    throw new Error(errorText);
                } else {
                    state.Status === 0
                        ? setState({
                            Like: state.Like + 1,
                            Dislike: state.Dislike,
                            Status: 1
                        })
                        : setState({
                            Like: state.Like + 1,
                            Dislike: state.Dislike - 1,
                            Status: 1
                        });
                    setErrors({});
                }
            } catch (error) {
                setErrors({"error": error.message});
            }
        }
    }

    async function handleDislike(event) {
        event.preventDefault();
        if (state.Status === 0 || state.Status === 1) {
            try {
                const response = await fetch('/auth/' + props.Addr + '/rate/' + props.ID,
                    {
                        method: 'POST',
                        body: JSON.stringify({
                            Like : false,
                            Src : props.Addr
                        }),
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + authContext.token,
                        },
                    }
                );
                await response.json();
                if (!response.ok) {
                    let errorText = 'не удалось оставить оценку';
                    throw new Error(errorText);
                } else {
                    state.Status === 0
                        ? setState({
                            Like: state.Like,
                            Dislike: state.Dislike + 1,
                            Status: -1
                        })
                        : setState({
                            Like: state.Like - 1,
                            Dislike: state.Dislike + 1,
                            Status: -1
                        });
                    setErrors({});
                }
            } catch (error) {
                setErrors({"error": error.message});
            }
        }
    }

    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <div className="row">
            <div className="row h-50 ml-4 p-4 ">
                <div className="btn-group" role="group" aria-label="Rate">
                    <button type="button" className="btn btn-success" onClick={handleLike}>
                        Нравится | {state.Like}
                    </button>
                    <button type="button" className="btn btn-danger" onClick={handleDislike}>
                        Не нравится | {state.Dislike}
                    </button>
                </div>
            </div>
            {errorContent}
        </div>

    );
};

export default Rate;

