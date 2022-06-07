import { useState, useContext, useEffect, useCallback } from 'react';

import AuthContext from '../../db/auth-context';
import Errors from '../Errors/Errors';

const PostForm = (props) => {
    const authContext = useContext(AuthContext);

    const [contentValue, setContentValue] = useState('');

    const [userName, setUserName] = useState({});

    const populateFields = useCallback(() => {
        if (props.post) {
            setUserName(props.post.Username);
            setContentValue(props.post.Review);
        }
    }, [props.post]);

    useEffect(() => {
        populateFields();
    }, [populateFields]);

    async function submitHandler(event) {
        event.preventDefault();
        setUserName({});

        try {
            const method = 'POST';
            let body = {
                Review: contentValue,
            }
            const response = await fetch('/auth/film/' + props.filmID,
                {
                    method: method,
                    body: JSON.stringify(body),
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'Failed to add new post.';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setUserName({ 'unknown': data['error'] })
                } else {
                    setUserName(data['error']);
                }
            } else {
                setUserName({});
                setContentValue('');
                if (props.onAddPost) {
                    props.onAddPost(data.post);
                }
            }
        } catch (error) {
            setUserName({ "error": error.message });
        }
    }

    const contentChangeHandler = (event) => { setContentValue(event.target.value) }

    const errorContent = Object.keys(userName).length === 0 ? null : Errors(userName);
    const submitButtonText = 'Add Post';

    return (
        <section>
            <div className="container w-75 pb-4">
                <form onSubmit={submitHandler}>
                    <div className="form-group pb-3">
                        <label htmlFor="content">Message</label>
                        <textarea id="content" className="form-control" rows="5" required value={contentValue} onChange={contentChangeHandler}></textarea>
                    </div>
                    <button type="submit" className="btn btn-success">{submitButtonText}</button>
                </form>
                {errorContent}
            </div>
        </section>
    );
}

export default PostForm;
