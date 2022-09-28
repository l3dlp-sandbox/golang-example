import React from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom';

import Login from './Login';
import PrivateRoute from './PrivateRoute'
import DefaultDropDown from './DefaultDropDown'; 

const App = () => (
    <Router basename="/web">
        <Switch>
            <Redirect exact from="/" to="/login" />
            <Route path="/login">
                <Login />
            </Route>
            <PrivateRoute path="/*">
                <DefaultDropDown />
            </PrivateRoute>
        </Switch>
    </Router>
);

export default App;
