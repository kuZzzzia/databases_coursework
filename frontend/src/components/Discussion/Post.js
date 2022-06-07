import {Link} from "react-router-dom";

const Post = (props) => {
    const user = props.post.UserName.Valid
        ? <Link to={'/user/' + props.post.UserID.Int16}>{props.post.UserName.String}</Link>
        : 'User deleted'

    return (
        <div className="card mb-2">
            <div className="card-header">{user}</div>
            <div className="card-body">{props.post.Review}</div>
        </div>
    )
};

export default Post;