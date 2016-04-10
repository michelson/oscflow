
#ReactDOM.render(<App/>, document.getElementById('content'));



ReactDOM.render((
  <ReactRouter.Router history={ReactRouter.browserHistory}>
    <ReactRouter.Route path="/" component={App}>
      <ReactRouter.Route path="patches/:patch" component={Patch}/>
    </ReactRouter.Route>
  </ReactRouter.Router>
), document.getElementById('content'))

