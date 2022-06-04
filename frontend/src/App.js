import { Routes, Route, Navigate } from 'react-router-dom';
import { useContext } from 'react';
import AuthContext from './db/auth-context';
import HomePage from "./pages/HomePage";
import AuthPage from "./pages/AuthPage";
import Layout from './components/Layout/Layout';
import UserPage from "./pages/UserPage";

function App() {
  const authContext = useContext(AuthContext);

  return (
      <Layout>
          <Routes>
              <Route path='/'>
                  <HomePage/>
              </Route>
              {!authContext.loggedIn && (
                  <Route path='/auth'>
                      <AuthPage />
                  </Route>
              )}
              <Route path='/profile'>
                  {authContext.loggedIn && <UserPage />}
                  {!authContext.loggedIn && <Navigate to="/auth" />}
              </Route>
              <Route path='*'>
                  <Navigate to="/"/>
              </Route>
          </Routes>
      </Layout>
  );
}

export default App;
