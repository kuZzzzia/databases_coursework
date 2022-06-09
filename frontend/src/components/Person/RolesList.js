import RoleRow from "./RoleRow";

const RolesList = (props) => {
    return (
        <div>
            {props.roles.map((role) => (
                <RoleRow
                    key={role.FilmID}
                    role={role}
                />
            ))}
        </div>
    );
};

export default RolesList;