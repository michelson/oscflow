
#ReactDOM.render(<Welcome/>, document.getElementById('content'));


ReactDOM.render((
  <ReactRouter.Router history={ReactRouter.browserHistory}>
    <ReactRouter.Route path="/" component={Welcome}>
    </ReactRouter.Route>
  </ReactRouter.Router>
), document.getElementById('content'))

