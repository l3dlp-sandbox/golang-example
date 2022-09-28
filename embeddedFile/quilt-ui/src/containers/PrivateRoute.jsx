import React from 'react';
import { Route, Redirect } from 'react-router-dom';

import { getCookie } from '../helper';

// PrivateRoute Checks that user is logged in, if not logged in, user will be redirected to landing page to login
const PrivateRoute = ({ children, ...rest }) => { 
    const isLoggedIn = getCookie("isLoggedIn")
    return (
        <Route
            /* eslint-disable react/jsx-props-no-spreading */
            {...rest}
            render={({ location }) =>
            isLoggedIn ? (
                    children
                ) : (
                    <Redirect
                        to={{
                            pathname: "/",
                            state: { from: location }
                        }}
                    />
                )
            }
      />
    );
};

export default PrivateRoute