
#ReactDOM.render(<Welcome/>, document.getElementById('content'));


ReactDOM.render((
  <ReactRouter.Router history={ReactRouter.browserHistory}>
    <ReactRouter.Route path="/" component={Welcome}>
    </ReactRouter.Route>
    <ReactRouter.Route path="patches/:patch" component={Patch}/>
  </ReactRouter.Router>
), document.getElementById('content'))

