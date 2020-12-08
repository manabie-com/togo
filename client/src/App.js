import React, { useEffect, Suspense } from 'react';
import { Route, Switch, withRouter, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';

import Task from './container/Task';
import Logout from './container/Auth/Logout';
import Auth from './container/Auth';

import * as actions from './redux/actions/index';

const _App = props => {
  const { tryAuth } = props;

  useEffect(() => {
    tryAuth();
  }, [tryAuth]);

  let routes = ( 
    <Switch>
      <Route path="/auth" render={props => <Auth {...props} />} />
      <Redirect to="/auth" />
    </Switch>
  );

  if (props.isAuthenticated) {
    routes = (
      <Switch>
        <Route path="/todo" render={props => <Task {...props} />} />  
        <Route path="/logout" component={Logout} />
        <Redirect to="/todo" />
      </Switch>
    )
  }

  return (
    <div>  
      <Suspense fallback={<p>Loading...</p>}>{routes}</Suspense>
    </div>
  );
}

const mapStateToProps = state => {
  return {
    isAuthenticated: state.auth.token !== null
  };
};

const mapDispatchToProps = dispatch => {
  return {
    tryAuth: () => dispatch(actions.authCheckState())
  };
};

export default withRouter(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(_App)
);
