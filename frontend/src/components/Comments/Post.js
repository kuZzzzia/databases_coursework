const Post = (props) => {
    const user = props.post.UserName.Valid
        ? props.post.UserName.String
        : 'User deleted'

    return (
        <div className="card mb-2">
            <div className="card-header">{user}</div>
            <div className="card-body">{props.post.Review}</div>
        </div>
    )
};

export default Post;