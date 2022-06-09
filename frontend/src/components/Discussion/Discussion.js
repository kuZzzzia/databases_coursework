import {useState} from "react";
import PostForm from "./PostForm";
import PostsList from "./PostsList";

const Discussion = (props) => {
    const [posts, setPosts] = useState(props.discussion);

    const addPostHandler = (postData) => {
        setPosts((prevState) => { return [postData, ...prevState] });
    }

    return (
        <section>
            <h5 className="pb-4 pt-4">Комментарии</h5>
            <PostsList posts={posts}/>
            <PostForm onAddPost={addPostHandler} filmID={props.filmID}/>
        </section>
    );
};

export default Discussion;