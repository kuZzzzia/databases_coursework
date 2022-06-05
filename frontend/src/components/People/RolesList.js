const RolesList = (props) => {
    return (
        <div>
            {props.roles.map((role) => (
                <div className="row">
                    <div className="col-lg-5">
                        {role.Name}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default RolesList;