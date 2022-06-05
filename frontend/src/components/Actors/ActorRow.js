import { useState, useContext } from 'react';

import Errors from '../Errors/Errors';
import PostForm from "./PostForm";

const ActorRow = (props) => {
const ActorRow = (props) => {

    const cardTitle = editing ? 'Edit post' : props.actor.Name;
    const cardBody = editing ? <PostForm post={props.post} onEditPost={editPostHandler} editing={true}/> : props.post.Content;
    const switchModeButtonText = editing ? 'Cancel' : 'Edit';
    const cardButtons = editing ?
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
        </div>
        :
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
        </div>
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <div className="card mb-5 pb-2">
            <div className="card-header">{cardTitle}</div>
            <div className="card-body">{cardBody}</div>
            {cardButtons}
            {errorContent}
        </div>
    );
};

export default Post;