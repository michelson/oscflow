@MainControl = React.createClass
  render: ->
    <div className="mdl-layout__drawer">
      <span className="mdl-layout-title">OscFlow</span>
      <nav className="mdl-navigation">
        <a className="mdl-navigation__link" href="">Play</a>
        <a className="mdl-navigation__link" href="">Rec</a>
        <a className="mdl-navigation__link" href="">Status</a>
      </nav>
    </div>