import {useState} from "react";
import PostForm from "./PostForm";
import PostsList from "./PostList";

const Discussion = (props) => {
    const [posts, setPosts] = useState(props.discussion);

    const addPostHandler = (postData) => {
        setPosts((prevState) => { return [postData, ...prevState] });
    }


    return (
        <section>
            <h4 className="pb-4">Discussion</h4>
            <PostsList posts={posts}/>
            <PostForm onAddPost={addPostHandler} filmID={props.filmID}/>
        </section>
    );
};

export default Discussion;