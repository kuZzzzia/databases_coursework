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
              <Route path='/' element={<HomePage />} />
              {!authContext.loggedIn && (
                  <Route path='/auth' element={
                      <AuthPage />
                  }/>
              )}
              <Route path='/profile' element={
                  authContext.loggedIn ?
                      <UserPage /> : <Navigate to="/auth" />
              }/>
              <Route path='*' element={
                  <Navigate to="/"/>
              }/>
          </Routes>
      </Layout>
  );
}

export default App;
