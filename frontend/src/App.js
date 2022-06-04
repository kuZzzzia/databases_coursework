import { Switch, Route, Redirect } from 'react-router-dom';
import { useContext } from 'react';
import AuthContext from './db/auth-context';

function App() {
  const authContext = useContext();

  return (
      <Layout>
          <Switch>
              <Route path="/" exact>
                  <HomePage/>
              </Route>
              {!authContext.loggedIn && (
                  <Route path="/auth">
                      <AuthPage />
                  </Route>
              )}
              <Route path="/profile">
                  {authContext.loggedIn && <UserPage />}
                  {!authContext.loggedIn && <Redirect to="/auth" />}
              </Route>
              <Route path="*">
                  <Redirect to="/"/>
              </Route>
          </Switch>
      </Layout>
  );
}

export default App;
