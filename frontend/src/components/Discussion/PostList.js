import Post from "./Post";

const PostsList = (props) => {
    return (
        <div style={{overflowY: 'scroll', maxHeight: '40rem'}}>
            {props.posts.map((post) => (
                <Post
                    key={post.ID}
                    post={post} />
            ))}
        </div>
    );
};

export default PostsList;