import { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import AuthContext from '../../db/auth-context';

const NavigationBar = () => {
    const navigate = useNavigate();

    const authContext = useContext(AuthContext);

    const loggedIn = authContext.loggedIn;

    const logoutHandler = () => {
        authContext.logout();
        navigate('/');
    };

    return (
        <nav className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
            <button type="button" className="navbar-toggler" data-toggle="collapse" data-target="#navbarCollapse"
                    aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse justify-content-between" id="navbarCollapse">
                <div className="d-flex flex-row">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item">
                            <Link className="nav-link" to="/">КИНОСЕТЬ</Link>
                        </li>
                        {loggedIn && (
                            <li className="nav-item">
                                <Link className="nav-link" to="/profile">Мой профиль</Link>
                            </li>
                        )}
                        {loggedIn && (
                            <li className="nav-item">
                                <Link className="nav-link" to="/profile/playlist">Создать подборку</Link>
                            </li>
                        )}
                    </ul>
                </div>
                <div className="d-flex flex-row justify-content-end">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item">
                            <Link className="nav-link" to="/people">Персоны</Link>
                        </li>
                        <li className="nav-item">
                            <Link className="nav-link" to="/films">Фильмы</Link>
                        </li>
                        {!loggedIn && (
                            <li className="nav-item">
                                <Link className="nav-link" to="/auth">Войти</Link>
                            </li>
                        )}
                        {loggedIn && (
                            <li className="nav-item">
                                <button className="btn btn-dark" onClick={logoutHandler}>Выйти</button>
                            </li>
                        )}
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default NavigationBar;
