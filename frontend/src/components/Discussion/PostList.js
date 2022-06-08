import Post from "./Post";

const PostsList = (props) => {
    return (
        <div className="container w-75 mb-3" style={{overflowY: 'scroll', maxHeight: '20rem'}}>
            {props.posts.map((post) => (
                <Post
                    key={post.ID}
                    post={post} />
            ))}
        </div>
    );
};

export default PostsList;